package api

type Service interface {
	Add(in *AddIn, out *AddOut) error
}

type AddIn struct {
	A, B int
}

type AddOut struct {
	Sum int
}
