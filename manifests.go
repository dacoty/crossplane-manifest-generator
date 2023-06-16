package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func CreateProviderManifest(MetadataName string, SpecPackage string) {
	tmpl, err := template.ParseFiles("templates/provider.tmpl")
	if err != nil {
		log.Fatal("parse", err)
	}
	p := ProviderParameters{ProviderName: MetadataName, ProviderPackage: SpecPackage}

	outputFile := fmt.Sprintf("%s.yaml", MetadataName)
	output, err := os.Create("providers/" + outputFile)
	if err != nil {
		log.Fatal("output file", err)
	}
	err = tmpl.Execute(output, p)
	if err != nil {
		log.Fatal("execute", err)
		return
	}
}
