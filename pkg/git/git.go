package git

import (
	"context"
	"fmt"

	shell "github.com/rancher/terraform-controller/pkg/cmd"
	"github.com/sirupsen/logrus"
)

func GetCommit(ctx context.Context, url, branch, tag string, auth *Auth) (string, error) {
	url, env, close := auth.Populate(url)
	defer close()

	var cmd = shell.Command{
		Command: "git",
		Env:     env,
		Args:    []string{"ls-remote", url, formatRef(branch, tag)},
	}
	lines, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return firstField(lines, fmt.Sprintf("no commit for branch: %s or tag: %s", branch, tag))
}

func CloneRepo(ctx context.Context, url string, commit string, auth *Auth) error {
	url, env, close := auth.Populate(url)
	defer close()

	var cloneCmd = shell.Command{
		Command: "git",
		Env:     env,
		Args:    []string{"clone", "-n", url, "."},
	}
	var checkoutCmd = shell.Command{
		Command: "git",
		Env:     env,
		Args:    []string{"checkout", commit},
	}

	_, err := shell.Execute(cloneCmd)
	if err != nil {
		return err
	}

	logrus.Infof("git clone: Done")

	_, err = shell.Execute(checkoutCmd)
	if err != nil {
		return err
	}

	logrus.Infof("git checkout: done")

	return nil
}
