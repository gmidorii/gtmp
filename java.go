package main

type Java struct {
	Package string
	Class   string
	Methods []Method
	Injects []Inject
}

type Method struct {
	method string
}

type Inject struct {
	inject string
}

func (b *Java) Create() error {
	return nil
}

func NewJava() (Language, error) {
	return &Java{}, nil
}
