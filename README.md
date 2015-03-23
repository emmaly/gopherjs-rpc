# gopherjs-rpc
Example of use of gopherjs + websocket + jsonrpc

## Preparation
* Have Go 1.4 or better ([`gopherjs`](https://github.com/gopherjs/gopherjs) requirement)
* `go get -u -v github.com/gopherjs/gopherjs` [[`godoc`](https://godoc.org/github.com/gopherjs/gopherjs)]
* `go get -u -v github.com/gopherjs/websocket` [[`godoc`](https://godoc.org/github.com/gopherjs/websocket)]
* `go get -u -v github.com/gopherjs/jquery` [[`godoc`](https://godoc.org/github.com/gopherjs/jquery)]
* `go get -u -v honnef.co/go/js/dom` [[`godoc`](https://godoc.org/honnef.co/go/js/dom)]

## Get the code and run!
* `cd $GOROOT` (or to any of the directories featured in your `$GOROOT` if you have more than one, like me)
* `git clone https://github.com/dustywilson/gopherjs-rpc.git`
* `cd gopherjs-rpc`
* `cd www` to enter the `www` directory where the `client.go` file is waiting for you to build
* `gopherjs build -o client.js` to build the `client.go` file into `client.js`
* `cd ..` to return to parent
* `go run server.go` to run the server-side
* Go to [`http://localhost:5454/`](http://localhost:5454/) in your browser

If you got log messages in the bottom part of the resulting test page, it's working.  The `Output` line is
what the JavaScript client is sending to the Go server.  The `Input` line is what the Go server sends to the
client in response.  The only thing that differs in each message is the timestamp.  The server spits out the
same info to the console on its end as well.

## Notes

Notice that both the backend and frontend share the file in the `common` directory.  In this case, they use
that to share the `ChatMessage` struct.  For the most part, you can share code between frontend and backend.

Notice that I'm using a channel on the frontend to handle log messages.  I just wanted to prove that channels
work as expected.

You *CAN'T* use the core websocket lib with `gopherjs`.  If you start getting errors saying something like
`gopherjs doesn't support network` then you've used the wrong package.

This is `jsonrpc` on top of `websocket`.  I didn't test it, but I fully expect that the frontend could host
RPC services for the server to call.
