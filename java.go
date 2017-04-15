package main

import (
	"io"
	"io/ioutil"

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

func (j *Java) Create(w io.Writer) error {
	_, err := ioutil.ReadFile(j.c.temp)
	if err != nil {
		return err
	}

	var parts Parts
	_, err = toml.DecodeFile(j.c.resource, &parts)
	if err != nil {
		return err
	}
	return nil
}

func NewJava(config Config) (Language, error) {
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
