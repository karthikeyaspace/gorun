package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Runcmd(service *Service) error {

	execPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting executable path")
		return err
	}

	service.Dir = filepath.Join(execPath, service.Dir)
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		if service.Type == "server" {
			cmd = exec.Command("cmd", "/C", "start", "cmd", "/K", service.Command)
		} else {
			cmd = exec.Command("cmd", "/C", service.Command)
		}
	default: // Linux and others
		if service.Type == "server" {
			cmd = exec.Command("gnome-terminal", "--", "sh", "-c", fmt.Sprintf("cd %s && %s", service.Dir, service.Command))
		} else {
			cmd = exec.Command("sh", "-c", service.Command)
		}
	}

	cmd.Dir = service.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Start()

}
