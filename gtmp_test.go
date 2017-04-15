package main

import (
	"errors"
	"testing"
)

type FakeLang struct {
	Fake Language
}

func (f *FakeLang) Create() error {
	switch f.Fake.(type) {
	case *Java:
		return nil
	default:
		return errors.New("language error")
	}
}

func TestJavaServer(t *testing.T) {
	javaLang := &FakeLang{
		Fake: &Java{},
	}
	server := &Parser{
		Language: javaLang,
	}
	err := server.Do()
	if err != nil {
		t.Error("Failed server")
		t.Log()
	}
}

func TestDefaultServer(t *testing.T) {
	nilLang := &FakeLang{
		Fake: nil,
	}
	server := &Parser{
		Language: nilLang,
	}
	err := server.Do()
	if err.Error() != "language error" {
		t.Error("Failed server")
		t.Log()
	}
}

func TestSwitchLang_Java(t *testing.T) {
	lang := "java"
	l, err := switchLang(lang)
	if err != nil {
		t.Error(err)
	}
	switch v := l.(type) {
	case *Java:
	default:
		t.Error("switch failed")
		t.Log(v)
	}
}

func TestSwitchLang_NotCompativle(t *testing.T) {
	lang := "swift"
	l, err := switchLang(lang)
	if err == nil {
		t.Error(err)
		t.Log(l)
	}
}
