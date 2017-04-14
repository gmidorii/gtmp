package main

import "testing"
import "errors"

type FakeLang struct {
	Fake Language
}

func (f *FakeLang) Create() error {
	switch f.Fake.(type) {
	case *Base:
		return nil
	default:
		return errors.New("language error")
	}
}

func TestJavaServer(t *testing.T) {
	javaLang := &FakeLang{
		Fake: &Base{},
	}
	server := &Server{
		Language: javaLang,
	}
	err := server.Start()
	if err != nil {
		t.Error("Failed server")
		t.Log()
	}
}

func TestDefaultServer(t *testing.T) {
	nilLang := &FakeLang{
		Fake: nil,
	}
	server := &Server{
		Language: nilLang,
	}
	err := server.Start()
	if err.Error() != "language error" {
		t.Error("Failed server")
		t.Log()
	}
}
