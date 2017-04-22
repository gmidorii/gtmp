package main

import (
	"io/ioutil"
	"os"
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

	r, _ := readResource(java.c.resource)
	tmp, err := template.ParseFiles(java.c.temp)
	if err != nil {
		t.Error(err)
	}

	java.create(r, tmp)

	act, err := ioutil.ReadFile("TestTest.java")
	if err != nil {
		t.Error(err)
	}
	exp, err := ioutil.ReadFile("./test/java/impl.java")
	if err != nil {
		t.Error(err)
	}

	if string(exp) != string(act) {
		t.Error("create template exp failed")
	}

	if err := os.Remove("TestTest.java"); err != nil {
		t.Error(err)
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
		Class:   "Test",
		Methods: m,
		Injects: i,
	}
}
