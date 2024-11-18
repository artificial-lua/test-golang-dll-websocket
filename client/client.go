package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"test-dll-websocket/model"
)

func ConnectWebsocket() {
	target := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws",
	}

	if conn, _, err := websocket.DefaultDialer.Dial(target.String(), nil); err != nil {
		log.Fatal("Failed to dial:", err)
		return
	} else {
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatal("Failed to close:", err)
			}
		}(conn)

		for {
			var msg *model.Message

			msg = &model.Message{
				Status:      model.Request,
				Content:     "Hello, World!",
				Token:       "1234567890",
				JsonContent: "{}",
			}

			if err := conn.WriteJSON(*msg); err != nil {
				log.Fatal("Failed to write message:", err)
				return
			}

			if err := conn.ReadJSON(msg); err != nil {
				log.Fatal("Failed to read message:", err)
				return
			}

			log.Printf("Received: %+v", *msg)

			msg = &model.Message{
				Status:      model.Close,
				Content:     "Goodbye, World!",
				Token:       "1234567890",
				JsonContent: "{}",
			}

			if err := conn.WriteJSON(*msg); err != nil {
				log.Fatal("Failed to write message:", err)
				return
			}

			if err := conn.ReadJSON(msg); err != nil {
				log.Fatal("Failed to read message:", err)
				return
			}

			break
		}
	}
}
