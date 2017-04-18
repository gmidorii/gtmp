package main

import (
	"io"
	"text/template"

	"github.com/BurntSushi/toml"
)

type Java struct {
	c    Config
	part Parts
}

type Parts struct {
	Package string
	Class   string
	Methods []string
	Injects []string
}

func (j *Java) create(w io.Writer, r string, t *template.Template) error {
	parts, err := createParts(r)
	if err != nil {
		return err
	}
	return t.Execute(w, parts)
}

func createParts(r string) (Parts, error) {
	var parts Parts
	_, err := toml.Decode(r, &parts)
	if err != nil {
		return Parts{}, err
	}

	return parts, nil
}
