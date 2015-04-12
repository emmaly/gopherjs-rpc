# gopherjs-rpc
Example of use of gopherjs + websocket + jsonrpc

## Preparation
* Have Go 1.4 or better ([`gopherjs`](https://github.com/gopherjs/gopherjs) requirement)
* `go get -u -v github.com/gopherjs/gopherjs` [[`godoc`](https://godoc.org/github.com/gopherjs/gopherjs/js)]
* `go get -u -v github.com/gopherjs/websocket` [[`godoc`](https://godoc.org/github.com/gopherjs/websocket)]
* `go get -u -v honnef.co/go/js/dom` [[`godoc`](https://godoc.org/honnef.co/go/js/dom)]

## Get the code and run!
* `cd $GOROOT` (or to any of the directories featured in your `$GOROOT` if you have more than one, like me)
* `git clone https://github.com/dustywilson/gopherjs-rpc.git`
* `cd gopherjs-rpc`
* `cd www` to enter the `www` directory where the `client.go` file is waiting for you to build
* `gopherjs build -m -o client.js` to build the `client.go` file into `client.js`
* `cd ..` to return to parent
* `go run server.go` to run the server-side
* Go to [`http://localhost:5454/`](http://localhost:5454/) in your browser

If you got log messages in the browser dev console, it's working.  The `Output` line is what the JavaScript
client is sending to the Go server.  The `Input` line is what the Go server sends to the client in response.
The only thing that differs in each message is the timestamp.  The server spits out the same info to the
console on its end as well.

## Notes

Notice that both the backend and frontend share the file in the `shared` directory.  In this case, they use
that to share the `ChatMessage` struct.  For the most part, you can share code between frontend and backend.

You *CAN'T* use the core websocket lib with `gopherjs`.  If you start getting errors saying something like
`gopherjs doesn't support network` then you've used the wrong package.

This is `rpc` on top of `websocket`.  In this case, the web browser is hosting RPC for the server to call.
For as long as the client is connected to the server, the server is able to call on any of its RPC
functions.  Totally backwards, but really cool that it's possible.
