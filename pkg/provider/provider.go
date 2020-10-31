package provider

import (
	"context"
	"net/http"

	tfv1 "github.com/rancher/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
)

const (
	DefaultSecretName = "gitcredential"
)

type Provider interface {
	Supports(obj *tfv1.Module) bool
	HandleHook(ctx context.Context, req *http.Request) (int, error)
}
