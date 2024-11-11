package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// does not have concurrency - not simultaneously running client and server in the same process

func Runcmd(service *Service) {
	fmt.Printf("Starting service: %v\n", service.Name)

	execPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting executable path")
		return
	}

	if service.Dir == "root" || service.Dir == "" {
		service.Dir = "/"
	}

	service.Dir = filepath.Join(execPath, service.Dir)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", service.Command)
	} else {
		cmd = exec.Command("sh", "-c", service.Command)
	}
	cmd.Dir = service.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error starting service %v\n", service.Name)
	}

	fmt.Println()

}
