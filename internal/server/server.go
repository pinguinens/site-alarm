package server

import (
	"net"
	"strings"

	log "github.com/rs/zerolog"

	"github.com/pinguinens/site-alarm/internal/service"
)

type API struct {
	service *service.Service
}

func Start(svc *service.Service) error {
	ln, err := net.Listen("tcp", svc.GetAddr())
	if err != nil {
		svc.Logger.Error().Msg(err.Error())
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn, svc.Logger)
	}
}

func handleConnection(conn net.Conn, logger *log.Logger) {
	for {
		buffer := make([]byte, 128)
		_, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				conn.Close()
				return
			}

			logger.Error().Msg(err.Error())
			return
		}
		if buffer == nil {
			logger.Info().Msg("empty buf")
			return
		}
		content := make([]byte, 0, len(buffer))
		for _, v := range buffer {
			if v != 0 {
				content = append(content, v)
			}
		}

		parts := strings.Split(string(content), "\n")

		logger.Info().Str("code", parts[0]).Str("method", parts[1]).Str("url", parts[2]).Str("addr", parts[3]).Send()
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
