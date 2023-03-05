package main

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/rs/zerolog/log"
)

const (
	appVersion = "0.1"
)

func main() {
	http.Handle("/alarm", websocket.Handler(Server))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func Server(ws *websocket.Conn) {
	for {
		var bytes []byte
		if err := websocket.Message.Receive(ws, &bytes); err != nil {
			log.Err(err)
			return
		}

		log.Info().Bytes("input", bytes).Send()
	}
}
