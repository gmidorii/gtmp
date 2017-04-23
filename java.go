package main

import (
	"text/template"

	"os"

	"github.com/BurntSushi/toml"
)

// Java is language struct
type Java struct {
	c    Config
	part Parts
}

// Parts is template parameter
type Parts struct {
	Package string
	Class   string
	Methods []string
	Injects []string
	Fields  []string
}

func (j *Java) create(r string, t *template.Template) error {
	parts, err := createParts(r)
	if err != nil {
		return err
	}
	file, err := os.Create(parts.Class + "Test.java")
	if err != nil {
		return err
	}
	return t.Execute(file, parts)
}

func createParts(r string) (Parts, error) {
	var parts Parts
	_, err := toml.Decode(r, &parts)
	if err != nil {
		return Parts{}, err
	}

	return parts, nil
}
