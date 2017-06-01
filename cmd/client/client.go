package main

import (
	"fmt"
	"github.com/gopherjs/websocket"
	"github.com/norunners/hellows/dao"
	"github.com/norunners/hellows/util"
	"honnef.co/go/js/dom"
	"net/rpc/jsonrpc"
)

var document = dom.GetWindow().Document()
var output = document.GetElementByID("app").(*dom.HTMLPreElement)

func main() {
	a := 2
	b := 3

	conn, err := websocket.Dial("ws://localhost:1234/ws-rpc")
	util.Must(err)
	client := jsonrpc.NewClient(conn)
	dao := dao.New(client)

	sum, err := dao.Add(a, b)
	util.Must(err)

	content := fmt.Sprintf("%v + %v = %v\n", a, b, sum)
	output.SetTextContent(output.TextContent() + content)

	client.Close()
	conn.Close()
}
