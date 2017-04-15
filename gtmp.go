package main

import (
	"errors"
	"flag"
	"log"
)

type Language interface {
	Create() error
}

type Parser struct {
	Language Language
}

func (p *Parser) Do() error {
	return p.Language.Create()
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

	if err := parser.Do(); err != nil {
		log.Fatalln(err)
	}
}

func switchLang(lang string) (Language, error) {
	switch lang {
	case "java":
		return NewJava()
	default:
		return nil, errors.New("Not Compatible. Language: " + lang)
	}
}
