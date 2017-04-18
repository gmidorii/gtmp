package main

import (
	"errors"
	"io"
	"os"
	"testing"
	"text/template"
)

type FakeLang struct {
	Fake Language
}

func (f *FakeLang) create(w io.Writer, r string, t *template.Template) error {
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
	err := server.Do(os.Stdout, "", template.New("tmp"))
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
	err := server.Do(os.Stdout, "", template.New("tmp"))
	if err.Error() != "language error" {
		t.Error("Failed server")
		t.Log()
	}
}

func TestSwitchLang_Java(t *testing.T) {
	lang := "java"
	l, _, err := switchLang(lang)
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
	l, _, err := switchLang(lang)
	if err == nil {
		t.Error(err)
		t.Log(l)
	}
}

func TestSwitchLang_GetConfig(t *testing.T) {
	lang := "java"
	_, c, err := switchLang(lang)
	if err != nil {
		t.Error(err)
	}
	exp := Config{
		resource: "./resources/java/base.toml",
		temp:     "./template/java/test.java",
	}
	if exp != c {
		t.Error("config setting failed")
	}
}

func TestReadResouce(t *testing.T) {
	r, err := readResource("./test/read.txt")
	if err != nil {
		t.Error(err)
	}

	exp := `test
test test`

	if r != exp {
		t.Error("Read File Failed")
		t.Log(exp)
		t.Log(r)
	}
}
