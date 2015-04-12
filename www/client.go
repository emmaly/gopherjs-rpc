package main

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/dustywilson/gopherjs-rpc/shared"
	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go ready()
	})
}

func ready() {
	rpc.Register(&RPC{})
	ws, err := websocket.Dial("ws://localhost:5454/ws-rpc")
	if err != nil {
		println(err.Error())
	}
	rpc.ServeConn(ws)
}

type RPC struct{}

func (r RPC) Message(input *shared.ChatMessage, output *shared.ChatMessage) error {
	output.Name = "Client"
	output.Time = time.Now()
	output.Message = fmt.Sprintf("Hello, %s.", input.Name)
	return nil
}
