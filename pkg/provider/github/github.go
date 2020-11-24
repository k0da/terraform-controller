package github

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/go-github/v28/github"
	tfv1 "github.com/rancher/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	tfv1controller "github.com/rancher/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io/v1"
	"github.com/rancher/terraform-controller/pkg/types"
	corev1controller "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/kv"
)

const (
	GitWebHookParam = "tf-module"
	TokenKeyName    = "token"
)

type GitHub struct {
	modules     tfv1controller.ModuleController
	secretCache corev1controller.SecretCache
	namespace   string
}

func NewGitHub(rContext *types.Context) *GitHub {
	return &GitHub{
		namespace:   rContext.Namespace,
		modules:     rContext.Tfv1.Terraformcontroller().V1().Module(),
		secretCache: rContext.Core.Core().V1().Secret().Cache(),
	}
}

func (w *GitHub) Supports(obj *tfv1.Module) bool {
	return obj.Spec.WebhookSecretName != ""
}

func (w *GitHub) HandleHook(ctx context.Context, req *http.Request) (int, error) {
	receiverID := req.URL.Query().Get(GitWebHookParam)
	if receiverID == "" {
		return 0, nil
	}

	ns, name := kv.Split(receiverID, ":")
	module, err := w.modules.Cache().Get(ns, name)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	var token []byte
	if module.Spec.WebhookSecretName != "" {
		secret, err := w.secretCache.Get(ns, module.Spec.WebhookSecretName)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		token = secret.Data[TokenKeyName]
	}
	payload, err := github.ValidatePayload(req, []byte(token))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return w.handleEvent(ctx, event, module)
}

func matchEvent(ref string, module *tfv1.Module) bool {
	return strings.HasSuffix(ref, module.Status.Content.Git.Branch) ||
		strings.HasSuffix(ref, module.Status.Content.Git.Tag)
}

func (w *GitHub) handleEvent(ctx context.Context, event interface{}, module *tfv1.Module) (int, error) {
	switch event := event.(type) {
	case *github.PushEvent:
		parsed := event

		if !matchEvent(parsed.GetRef(), module) {
			return http.StatusOK, nil
		}
		tfv1.ModuleConditionGitUpdated.False(module)
	}
	if _, err := w.modules.Update(module); err != nil {
		return http.StatusConflict, err
	}
	return http.StatusOK, nil
}
