package main

import (
	"log"
	"net/http"
	"net/rpc"
	"time"

	"github.com/dustywilson/gopherjs-rpc/shared"
	"github.com/nytimes/gziphandler"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/ws-rpc", websocket.Handler(func(conn *websocket.Conn) {
		conn.PayloadType = websocket.BinaryFrame
		r := rpc.NewClient(conn)
		chatter(r)
	}))
	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("www"))))
	panic(http.ListenAndServe(":5454", nil))
}

func chatter(r *rpc.Client) {
	for {
		output := shared.ChatMessage{
			Name:    "Server",
			Time:    time.Now(),
			Message: "Something important.",
		}
		log.Println("Output: ", output)
		var input shared.ChatMessage
		err := r.Call("RPC.Message", &output, &input)
		if err != nil {
			log.Println("ERROR: ", err)
			if err == rpc.ErrShutdown {
				log.Println("... exiting!")
				return
			}
		} else {
			log.Println("Input: ", input)
		}
		time.Sleep(time.Second)
	}
}
