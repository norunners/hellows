package dao

import (
	"github.com/norunners/hellows/api"
	"net/rpc"
)

type Dao struct {
	client *rpc.Client
}

func New(client *rpc.Client) *Dao {
	return &Dao{client: client}
}

func (dao *Dao) Add(a, b int) (int, error) {
	in := &api.AddIn{A: a, B: b}
	out := &api.AddOut{}
	err := dao.client.Call("Service.Add", in, out)
	if err != nil {
		return 0, err
	}
	return out.Sum, nil
}
