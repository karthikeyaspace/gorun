package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Service struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Dir     string `yaml:"dir"`
}

type Config struct {
	Services []Service `yaml:"services"`
}

func main() {
	configFile := "gorun.yaml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Configuration file not found!")
		return
	}

	config := Config{}
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error opening configuration file: %v\n", err)
		return
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Printf("Error parsing configuration file: %v\n", err)
		return
	}

	for _, service := range config.Services {
		fmt.Println("Starting up", service.Name)
		Runcmd(&service)
	}

}
