package main

type Java struct {
	c    Config
	part Parts
}

type Parts struct {
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

func (j *Java) Create() error {
	return nil
}

func NewJava(config Config) (Language, error) {
	return &Java{
		c: config,
	}, nil
}
