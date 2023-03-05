package main

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

const (
	appVersion = "0.1"
)

func main() {
	http.Handle("/receive", websocket.Handler(ReceiveHandler))
	http.HandleFunc("/alarm", AlarmHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func ReceiveHandler(ws *websocket.Conn) {
	for {
		var bytes []byte
		if err := websocket.Message.Receive(ws, &bytes); err != nil {
			log.Err(err)
			return
		}

		log.Info().Bytes("input", bytes).Send()
	}
}

func AlarmHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err)
		return
	}

	log.Info().Bytes("input", bytes).Send()
}
