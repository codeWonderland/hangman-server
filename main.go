package main

import (
	"net/http"
	"golang.org/x/net/websocket"
)

var myWord *Word

type Message struct {
	Author string
	Body   string
}

//~/go $ go get -u golang.org/x/net/websocket

func main() {
	myWord = randWord()
	http.HandleFunc("/ws", broadcastHandler)
	http.ListenAndServe(":3000", nil)
}

func broadcastHandler(w http.ResponseWriter, req *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(wsHandler)}
	s.ServeHTTP(w, req)
}