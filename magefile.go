//go:build mage

package main

import (
	"os"
	"os/exec"

	"github.com/carolynvs/magex/pkg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() error {
	return sh.RunV("go", "build", "-o", "bin/t1", "./cmd/t1")
}

func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

func Clean() error {
	return sh.RunV("rm", "-rf", "bin")
}

func EnsureMage() error {
	return pkg.EnsureMage("")
}

func BashCompletion() error {
	Build()

	output, err := sh.Output("./bin/t1", "completion", "bash")
	if err != nil {
		return err
	}

	file, err := os.Create("/tmp/completion")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", "-c", "source /tmp/completion")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
