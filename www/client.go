package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"github.com/dustywilson/gopherjs-rpc/common"
	"github.com/gopherjs/jquery"
	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
)

var jQuery = jquery.NewJQuery
var consoleChan chan interface{}
var ws *websocket.Conn

func main() {
	jQuery().Ready(func() { go ready() })
}

func consoleHandler() chan interface{} {
	consoleChan := make(chan interface{})
	container := jQuery("#debug")
	go func() {
		for {
			value := <-consoleChan
			now := time.Now()
			div := dom.GetWindow().Document().CreateElement("div")
			div.SetTextContent(fmt.Sprintf("%.19s %v", now, value))
			container.Prepend(div)
			log.Print(div.TextContent())
			log.Print(value)
		}
	}()
	return consoleChan
}

func debug(out ...interface{}) {
	consoleChan <- fmt.Sprint(out)
}

func ready() {
	consoleChan = consoleHandler()
	var err error
	ws, err = websocket.Dial("ws://localhost:5454/ws-rpc")
	if err != nil {
		debug(err.Error())
	}
	r := jsonrpc.NewClient(ws)
	go chatter(r)
}

func chatter(r *rpc.Client) {
	for {
		output := common.ChatMessage{
			Name:    "Mr. VIP",
			Time:    time.Now(),
			Message: "Something important.",
		}
		debug("Output: ", output)
		var input common.ChatMessage
		err := r.Call("Chat.Message", &output, &input)
		if err != nil {
			debug("ERROR: ", err)
		} else {
			debug("Input: ", input)
		}
		time.Sleep(time.Second)
	}
}
