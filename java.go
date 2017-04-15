package main

import (
	"html/template"
	"io"

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

func (j *Java) create(w io.Writer) error {
	parts, err := createParts(j.c.resource)
	if err != nil {
		return err
	}

	t := template.New("java test")
	t, err = template.ParseFiles(j.c.temp)
	if err != nil {
		return err
	}
	return t.Execute(w, parts)
}

func java(config Config) (Language, error) {
	return &Java{
		c: config,
	}, nil
}

func createParts(filename string) (Parts, error) {
	var parts Parts
	_, err := toml.DecodeFile(filename, &parts)
	if err != nil {
		return Parts{}, err
	}

	return parts, nil
}
