package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
	"text/template"
)

func TestCreate(t *testing.T) {
	java := Java{
		c: Config{
			resource: "./test/java/base.toml",
			temp:     "./test/java/template.java",
		},
		part: create(),
	}

	r, _ := readResource(c.resource)
	tmp, err := template.ParseFiles(c.temp)
	if err != nil {
		log.Fatal(err)
	}

	stdOut := new(bytes.Buffer)
	java.create(stdOut, r, tmp)

	file, err := ioutil.ReadFile("./test/java/impl.java")
	if err != nil {
		t.Error(err)
	}

	if string(file) != stdOut.String() {
		t.Error("create template file failed")
	}
}

func TestCreateParts(t *testing.T) {
	r, _ := readResource("./test/java/base.toml")
	actual, err := createParts(r)
	if err != nil {
		t.Error(err)
	}

	expected := create()
	if expected.Package != actual.Package {
		t.Error("create parts Package failed")
		t.Log(expected)
		t.Log(actual)
	}
	if expected.Class != actual.Class {
		t.Error("create parts Class failed")
		t.Log(expected)
		t.Log(actual)
	}
	if !equalSlice(expected.Methods, actual.Methods) {
		t.Error("create parts Methods failed")
		t.Log(expected)
		t.Log(actual)
	}
	if !equalSlice(expected.Injects, actual.Injects) {
		t.Error("create parts Injects failed")
		t.Log(expected)
		t.Log(actual)
	}
}

func TestEqualSliceEqual(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"1", "2"}
	if !equalSlice(a, b) {
		t.Error("equals judge failed")
	}
}

func TestEqualSliceNotEqual(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"10", "10"}
	if equalSlice(a, b) {
		t.Error("equals judge failed")
	}
}

func TestEqualSliceNotEqualLen(t *testing.T) {
	a := []string{"1", "2"}
	b := []string{"10", "10", "30"}
	if equalSlice(a, b) {
		t.Error("equals judge failed")
	}
}

func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func create() Parts {
	m := []string{
		"search",
		"select",
	}

	i := []string{
		"TestImpl",
		"HogeClientImpl",
	}

	return Parts{
		Package: "house",
		Class:   "Watch",
		Methods: m,
		Injects: i,
	}
}
