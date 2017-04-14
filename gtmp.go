package main

import (
	"flag"
	"log"
)

type Language interface {
	Create() error
}

type Server struct {
	Language Language
}

func (s *Server) Start() error {
	return s.Language.Create()
}

func main() {
	lang := *flag.String("l", "java", "select creating template language")
	flag.Parse()

	l, err := NewJava(lang)
	if err != nil {
		log.Fatalln(err)
	}

	server := Server{
		Language: l,
	}

	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
