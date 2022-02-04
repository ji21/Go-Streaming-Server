package main

import (
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func readMessage(conn *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil || err == io.EOF {
			log.Fatal("Error reading: ", err)
			break
		}
		fmt.Printf("recv: %s", message)
	}
}

var addr string = "localhost:7000"

func main() {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go readMessage(conn, done)
	<-done
}
