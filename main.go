package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ProviderParameters struct {
	ProviderName    string
	ProviderPackage string
}

func writeToFile(filename string, data []ProviderParameters) {
	d, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(filename, d, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {

	configFile, err := os.ReadFile("templates/configs.yaml")
	if err != nil {
		log.Fatalf("can't read yaml file:  %v", err)
	}

	var baseConfigs []ProviderParameters
	err = yaml.Unmarshal(configFile, &baseConfigs)
	if err != nil {
		log.Fatalf("could not unmarshal file : %v", err)
	}

	for _, c := range baseConfigs {
		CreateProviderManifest(c.ProviderName, c.ProviderPackage)
	}
}
