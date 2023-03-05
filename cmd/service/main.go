package main

import (
	"bufio"
	"net"

	"github.com/rs/zerolog/log"
)

const (
	appVersion = "0.1"
)

func main() {
	// TODO: tcp messenger server
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Err(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Err(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Info().Str("input", message).Send()
		newmessage := "OK"
		conn.Write([]byte(newmessage + "\n"))
	}
}

//func main() {
//	http.Handle("/receive", websocket.Handler(ReceiveHandler))
//	http.HandleFunc("/alarm", AlarmHandler)
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		panic("ListenAndServe: " + err.Error())
//	}
//}
//
//func ReceiveHandler(ws *websocket.Conn) {
//	for {
//		var bytes []byte
//		if err := websocket.Message.Receive(ws, &bytes); err != nil {
//			log.Err(err)
//			return
//		}
//
//		log.Info().Bytes("input", bytes).Send()
//	}
//}
//
//func AlarmHandler(w http.ResponseWriter, r *http.Request) {
//	bytes, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Err(err)
//		return
//	}
//
//	log.Info().Bytes("input", bytes).Send()
//}
