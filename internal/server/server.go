package server

import (
	"net"

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
		go handleConnection(conn, svc)
	}
}

func handleConnection(conn net.Conn, svc *service.Service) {
	for {
		buffer, err := read(conn)
		if err != nil {
			if err.Error() == "EOF" {
				conn.Close()
				svc.Logger.Info().Msg("close connection")
				return
			}

			svc.Logger.Error().Msg(err.Error())
			return
		}
		if buffer == nil {
			svc.Logger.Info().Msg("empty buf")
			return
		}

		err = svc.Log(buffer)
		if err != nil {
			svc.Logger.Error().Msg(err.Error())
			return
		}

		if err = write(conn, []byte("OK")); err != nil {
			svc.Logger.Error().Msg("empty buf")
			return
		}
	}
}

func read(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 128)
	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func write(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		return err
	}

	return nil
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
