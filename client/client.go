package main

import (
	"fmt"
	"github.com/gopherjs/websocket"
	"github.com/norunners/hellows/api"
	"honnef.co/go/js/dom"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var document = dom.GetWindow().Document()
var output = document.GetElementByID("app").(*dom.HTMLPreElement)

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

func main() {
	a := 2
	b := 3

	conn, err := websocket.Dial("ws://localhost:1234/ws-rpc")
	api.Must(err)
	client := jsonrpc.NewClient(conn)
	dao := New(client)

	sum, err := dao.Add(a, b)
	api.Must(err)

	content := fmt.Sprintf("%v + %v = %v\n", a, b, sum)
	output.SetTextContent(output.TextContent() + content)

	client.Close()
	conn.Close()
}
