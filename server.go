package main

import (
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"github.com/dustywilson/gopherjs-rpc/common"
	"golang.org/x/net/websocket"
)

// Chat is RPC
type Chat struct{}

// Message is Chat.Message RPC
func (c *Chat) Message(input *common.ChatMessage, output *common.ChatMessage) error {
	log.Print(input)
	output.Name = "Server"
	*output = common.ChatMessage{
		Name:    "Server",
		Time:    time.Now(),
		Message: "I got your message.",
	}
	log.Print(output)
	return nil
}

func main() {
	rpc.Register(&Chat{})
	http.Handle("/ws-rpc", websocket.Handler(func(conn *websocket.Conn) {
		jsonrpc.ServeConn(conn)
	}))
	http.Handle("/", http.FileServer(http.Dir("www")))
	panic(http.ListenAndServe(":5454", nil))
}
