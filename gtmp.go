package main

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Language interface {
	create(w io.Writer, r string, t *template.Template) error
}

type Config struct {
	resource string
	temp     string
}

type Parser struct {
	Language Language
}

func (p *Parser) Do(w io.Writer, r string, t *template.Template) error {
	return p.Language.create(w, r, t)
}

var c Config

func main() {
	lang := *flag.String("l", "java", "select creating template language")
	flag.Parse()

	if err := run(lang); err != nil {
		log.Fatal(err)
	}
}

func run(lang string) error {
	l, c, err := switchLang(lang)
	if err != nil {
		return err
	}

	parser := Parser{
		Language: l,
	}

	r, err := readResource(c.resource)
	if err != nil {
		return err
	}

	t, err := template.ParseFiles(c.temp)
	if err != nil {
		return err
	}

	return parser.Do(os.Stdout, string(r), t)
}

func switchLang(lang string) (Language, Config, error) {
	config := Config{
		resource: "./resources/" + lang + "/base.toml",
	}
	switch lang {
	case "java":
		config.temp = "./template/java/test.java"
		return &Java{}, config, nil
	default:
		return nil, Config{}, errors.New("Not Compatible. Language: " + lang)
	}
}

func readResource(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	r, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
