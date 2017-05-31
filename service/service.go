package main

import (
	"fmt"
	"github.com/norunners/hellows/api"
	"golang.org/x/net/websocket"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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

func main() {
	service := New()
	server := rpc.NewServer()
	server.Register(service)
	handle := handler(server)

	http.Handle("/ws-rpc", websocket.Handler(handle))
	err := http.ListenAndServe("localhost:1234", nil)
	api.Must(err)
}

func handler(server *rpc.Server) websocket.Handler {
	return func(conn *websocket.Conn) {
		fmt.Printf("Serve conn begin.\n")
		server.ServeCodec(jsonrpc.NewServerCodec(conn))
		fmt.Printf("Serve conn end.\n")
	}
}
