package git

import (
	"context"
	"fmt"

	getter "github.com/hashicorp/go-getter"
)

func BranchCommit(ctx context.Context, url string, branch string, auth *Auth) (string, error) {
	url, env, close := auth.Populate(url)
	defer close()

	lines, err := git(ctx, env, "ls-remote", url, formatRefForBranch(branch))
	if err != nil {
		return "", err
	}

	return firstField(lines, fmt.Sprintf("no commit for branch: %s", branch))
}

func CloneRepo(ctx context.Context, url string, auth *Auth) error {
	url, _, close := auth.Populate(url)
	defer close()

	g := new(getter.GitGetter)

	client := &getter.Client{
		Ctx:     ctx,
		Src:     url,
		Dst:     ".",
		Pwd:     ".",
		Detectors: []getter.Detector{
			new(getter.GitDetector),
		},
		Getters: map[string]getter.Getter{
			"git": g,
		},
		Mode:    getter.ClientModeDir,
	}

	err := client.Get()
	if err != nil {
		return err
	}

	return nil
}
