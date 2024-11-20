package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"test-dll-websocket/model"
)

var conn *websocket.Conn

func ConnectWebsocket() {
	target := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws",
	}

	if ws, _, err := websocket.DefaultDialer.Dial(target.String(), nil); err != nil {
		log.Fatal("Failed to dial:", err)
		return
	} else {
		conn = ws
	}
}

func CloseWebsocket() {
	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close:", err)
		return
	}
}

func ReadMessage() (*model.Message, error) {
	var msg model.Message

	if err := conn.ReadJSON(&msg); err != nil {
		log.Fatal("Failed to read message:", err)
		return nil, err
	}

	return &msg, nil
}

func SendMessage(msg *model.Message) error {
	if err := conn.WriteJSON(*msg); err != nil {
		log.Fatal("Failed to write message:", err)
		return err
	}

	return nil
}

func MainRun() {
	ConnectWebsocket()

	defer CloseWebsocket()

	for {
		var sendMsg *model.Message

		sendMsg = &model.Message{
			Status:      model.Request,
			Content:     "Hello, World!",
			Token:       "1234567890",
			JsonContent: "{}",
		}

		if err := SendMessage(sendMsg); err != nil {
			log.Fatal("Failed to send message:", err)
			return
		}

		if recvMsg, err := ReadMessage(); err != nil {
			log.Fatal("Failed to read message:", err)
			return
		} else {
			log.Printf("Received: %+v", *recvMsg)
		}

		sendMsg = &model.Message{
			Status:      model.Close,
			Content:     "Goodbye, World!",
			Token:       "1234567890",
			JsonContent: "{}",
		}

		if err := SendMessage(sendMsg); err != nil {
			log.Fatal("Failed to send message:", err)
			return
		}

		if recvMsg, err := ReadMessage(); err != nil {
			log.Fatal("Failed to read message:", err)
			return
		} else {
			log.Printf("Received: %+v", *recvMsg)
		}

		break
	}
}
