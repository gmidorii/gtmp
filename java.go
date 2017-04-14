package main

type Base struct {
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

func (b Base) Create() error {
	return nil
}

func NewJava(lang string) (Language, error) {
	return &Base{}, nil
}
