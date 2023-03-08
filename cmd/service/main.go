package main

import (
	"net"
	"os"

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
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Err(err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				conn.Close()
				return
			}

			log.Err(err)
			return
		}
		if buffer == nil {
			log.Info().Msg("empty buf")
			return
		}
		content := make([]byte, 0, len(buffer))
		for _, v := range buffer {
			if v != 0 {
				content = append(content, v)
			}
		}

		log.Info().Bytes("input", content).Send()
		newmessage := "OK"
		conn.Write([]byte(newmessage))
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
