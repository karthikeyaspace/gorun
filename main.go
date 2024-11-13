package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Service struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Dir     string `yaml:"dir"`
	Type    string `yaml:"type"`
}

type Config struct {
	Services []Service `yaml:"services"`
}

func main() {
	fmt.Print("Gorun v1.3\n\n")

	configFile := "gorun.yaml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Configuration file not found!")
		return
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening configuration file: %v\n", err)
		return
	}

	config := Config{}
	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Printf("Error parsing configuration file: %v\n", err)
		return
	}

	for _, service := range config.Services {
		err := Runcmd(&service)
		if err != nil {
			fmt.Printf("Error running service %v\n", service.Name)
		}
	}

}
