package entity

type Amiibo struct {
	Name string
	Path string
}

func NewAmiibo(name string) *Amiibo {
	return &Amiibo{Name: name}
}
