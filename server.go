package main

import (
	"fmt"

	"golang.org/x/net/websocket"
	"math/rand"
	"time"
)

var words = []string{"programmer", "ambient", "diffuser", "aesthetic", "alternative"}

var clients []Client

var usernames []UserScore

var wsHandler = websocket.Handler(onWsConnect)

func onWsConnect(ws *websocket.Conn) {
	defer ws.Close()
	client := NewClient(ws)
	clients = addClientAndGreet(clients, client)
	client.listen()
}

func broadcast(msg *Message) {
	fmt.Printf("Broadcasting %+v\n", msg)
	for _, c := range clients {
		c.ch <- msg
	}
}

func addClientAndGreet(list []Client, client Client) []Client {
	clients = append(list, client)
	websocket.JSON.Send(client.connection, Message{"Server", "Welcome!"})
	websocket.JSON.Send(client.connection, Message{"Word", myWord.encryptedWord})



	for i := 0; i < len(myWord.lettersGuessed); i++ {
		websocket.JSON.Send(client.connection, Message{"Letter", myWord.lettersGuessed[i]})
	}
	return clients
}

func randWord() *Word {
	rand.Seed(time.Now().UTC().UnixNano())
	return makeWord(words[rand.Int() % 5])
}