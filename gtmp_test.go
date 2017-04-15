package main

import (
	"errors"
	"io"
	"os"
	"testing"
)

type FakeLang struct {
	Fake Language
}

func (f *FakeLang) Create(w io.Writer) error {
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
	err := server.Do(os.Stdout)
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
	err := server.Do(os.Stdout)
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

func TestSwitchLang_SetConfig(t *testing.T) {
	lang := "java"
	l, err := switchLang(lang)
	if err != nil {
		t.Error(err)
	}
	j, ok := l.(*Java)
	if ok != true {
		t.Error(ok)
	}
	exp := Config{
		resource: "./resources/java/base.toml",
		temp:     "./template/java/test.java",
	}
	if exp != j.c {
		t.Error("config setting failed")
	}
}
