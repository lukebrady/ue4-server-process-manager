package main

import (
	"os"
	"os/exec"
)

func createServerProcess(serverBinaryPath string) (*os.Process, error) {
	command := exec.Command(serverBinaryPath)
	err := command.Start()
	if err != nil {
		return nil, err
	}
	return command.Process, nil
}
