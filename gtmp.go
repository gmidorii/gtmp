package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

type Language interface {
	create(w io.Writer) error
}

type Config struct {
	resource string
	temp     string
}

type Parser struct {
	Language Language
}

func (p *Parser) Do(w io.Writer) error {
	return p.Language.create(w)
}

func main() {
	lang := *flag.String("l", "java", "select creating template language")
	flag.Parse()

	l, err := switchLang(lang)
	if err != nil {
		log.Fatalln(err)
	}

	parser := Parser{
		Language: l,
	}

	if err := parser.Do(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

func switchLang(lang string) (Language, error) {
	c := Config{
		resource: "./resources/" + lang + "/base.toml",
	}
	switch lang {
	case "java":
		c.temp = "./template/java/test.java"
		return java(c)
	default:
		return nil, errors.New("Not Compatible. Language: " + lang)
	}
}
