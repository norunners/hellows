package service

import (
	"fmt"
	"github.com/norunners/hellows/api"
)

type Service struct {
}

func New() api.Service {
	return &Service{}
}

func (service *Service) Add(in *api.AddIn, out *api.AddOut) error {
	out.Sum = in.A + in.B
	fmt.Printf("%v + %v = %v\n", in.A, in.B, out.Sum)
	return nil
}
