package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
	"strconv"
)

type Client struct {
	connection *websocket.Conn
	ch         chan *Message
	close      chan bool
}

type UserScore struct {
	username string
	score int
}

func NewClient(ws *websocket.Conn) Client {
	ch := make(chan *Message, 100)
	close := make(chan bool)

	return Client{ws, ch, close}
}

func (c *Client) listen() {
	go c.listenToWrite()
	c.listenToRead()
}

func (c *Client) listenToWrite() {
	for {
		select {
		case msg := <-c.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.connection, msg)

		case <-c.close:
			c.close <- true
			return
		}
	}
}

func (c *Client) listenToRead() {
	log.Println("Listening read from client")
	for {
		select {
		case <-c.close:
			c.close <- true
			return

		default:
			var msg Message
			err := websocket.JSON.Receive(c.connection, &msg)
			fmt.Printf("Received: %+v\n", msg)
			if err == io.EOF {
				c.close <- true
			} else if err != nil {
				// c.server.Err(err)
			} else if msg.Author == "ClientName" {
				usernames = append(usernames, UserScore{msg.Body, 0})
				broadcast(&Message{"User", msg.Body}) // Alert all users there is a new player

			} else if guessLetter(myWord, msg.Body) {

				broadcast(&Message{"Letter", msg.Body}) // if the user gets a correct letter, notify clients

				for index, user := range usernames {
					if msg.Author == user.username {
						usernames[index].score++
						broadcast(&Message{msg.Author, strconv.Itoa(usernames[index].score)})
						break
					}
				}

				if len(myWord.lettersLeft) == 0 { // if any letters are left
					broadcast(&Message{"Winner", msg.Author})
				}
			}
		}
	}
}
