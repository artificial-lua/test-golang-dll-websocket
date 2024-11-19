package main

import (
	"C"
	"test-dll-websocket/client"
	"test-dll-websocket/model"
)

func main() {
}

func ConnectWebsocket() {
	client.ConnectWebsocket()
}

func CloseWebsocket() {
	client.CloseWebsocket()
}

func ReadMessage() (*model.Message, error) {
	return client.ReadMessage()
}

func SendMessage(msg *model.Message) error {
	return client.SendMessage(msg)
}
