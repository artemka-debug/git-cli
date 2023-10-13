package utils

import (
	"errors"
	"fmt"
	"bytes"
	"os/exec"
)

func RunWithError(cmd *exec.Cmd) error {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil && stderr.String() != "" {
		return errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return nil
}