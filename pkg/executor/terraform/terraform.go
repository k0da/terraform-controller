// Package terraform to run terraform commands
package terraform

import (
	"os"

	shell "github.com/rancher/terraform-controller/pkg/cmd"
)

func Apply() (string, error) {
	var cmd = shell.Command{
		WorkingDir: getWorkingDir(),
		Command:    "terraform",
		Args:       []string{"apply", "-input=false", "-auto-approve", "tfplan"},
	}
	output, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}

func getWorkingDir() string {
	dir := os.Getenv("TF_DIR")
	if dir == "" {
		dir = "."
	}
	return dir
}

func Destroy() (string, error) {
	var cmd = shell.Command{
		WorkingDir: getWorkingDir(),
		Command:    "terraform",
		Args:       []string{"destroy", "-input=false", "-auto-approve"},
	}
	output, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}

func Init() (string, error) {
	var cmd = shell.Command{
		WorkingDir: getWorkingDir(),
		Command:    "terraform",
		Args:       []string{"init", "-input=false"},
	}
	output, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}

// Output runs 'terraform output -json' and returns the blob as a string
func Output() (string, error) {
	var cmd = shell.Command{
		WorkingDir: getWorkingDir(),
		Command:    "terraform",
		Args:       []string{"output", "-json"},
	}
	output, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}

// Plan runs 'terraform plan' with the destroy flag controlling the play type
func Plan(destroy bool) (string, error) {
	args := []string{"plan", "-input=false", "-out=tfplan"}
	if destroy {
		args = append(args, "-destroy")
	}

	var cmd = shell.Command{
		WorkingDir: getWorkingDir(),
		Command:    "terraform",
		Args:       args,
	}
	output, err := shell.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}
