package main

import (
	"C"
	"encoding/json"
	"log"
	"test-dll-websocket/client"
	"test-dll-websocket/model"
)

func main() {
	//client.MainRun()

	ConnectWebsocket()

	defer CloseWebsocket()

	var sendMsg string

	//snedMsg = &model.Message{
	//	Status:      model.Request,
	//	Content:     "Hello, World!",
	//	Token:       "1234567890",
	//	JsonContent: "{}",
	//}

	sendMsg = "{\"content\":\"Hello, World!\",\"json_content\":\"\",\"error\":\"\",\"token\":\"\",\"status\":1}"

	if err := SendMessage(sendMsg); err != nil {
		log.Fatal("Failed to send message:", err)
		return
	}

	if recvMsg, err := ReadMessage(); err != nil {
		log.Fatal("Failed to read message:", err)
		return
	} else {
		log.Printf("Received: %+v", recvMsg)
	}

	//sendMsg = &model.Message{
	//	Status:      model.Close,
	//	Content:     "Goodbye, World!",
	//	Token:       "1234567890",
	//	JsonContent: "{}",
	//}

	sendMsg = "{\"content\":\"Goodbye, World!\",\"json_content\":\"\",\"error\":\"\",\"token\":\"\",\"status\":8}"

	if err := SendMessage(sendMsg); err != nil {
		log.Fatal("Failed to send message:", err)
		return
	}

	if recvMsg, err := ReadMessage(); err != nil {
		log.Fatal("Failed to read message:", err)
		return
	} else {
		log.Printf("Received: %+v", recvMsg)
	}

}

//export ConnectWebsocket
func ConnectWebsocket() {
	client.ConnectWebsocket()
}

//export CloseWebsocket
func CloseWebsocket() {
	client.CloseWebsocket()
}

//export ReadMessage
func ReadMessage() (string, error) {
	if msg, err := client.ReadMessage(); err != nil {
		return "", err
	} else {
		jsonMsg, _ := json.Marshal(*msg)
		return string(jsonMsg), nil
	}
}

//export SendMessage
func SendMessage(jsonMsg string) error {
	/*
		json messsage like this:
		type Message struct {
			Content     string `json:"content"`
			JsonContent string `json:"json_content"`
			Error       string `json:"error"`
			Token       string `json:"token"`
			Status      Status `json:"status"`
		}
		{
			"content": "Hello, World!",
			"json_content": "",
			"error": "",
			"token": "",
			"status": 0
		}
	*/
	msg := &model.Message{}

	if err := json.Unmarshal([]byte(jsonMsg), msg); err != nil {
		return err
	}

	return client.SendMessage(msg)
}
