package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"test-dll-websocket/model"
)

var ws = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 간단한 echo 웹소켓 핸들러
func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade ws: %v", err)
		http.Error(w, "Failed to upgrade ws", http.StatusInternalServerError)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Failed to close: %v", err)
		}
	}(conn)
	for {
		var msg model.Message
		// Read message from the browser

		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Failed to read message: %v", err)
			return
		}

		// Print the message to the console
		log.Printf("Received: %+v", msg)

		// Write message back to the browser

		if err = conn.WriteJSON(msg); err != nil {
			log.Printf("Failed to write message: %v", err)
			return
		}

		if msg.Status == model.Close {
			log.Println("Close the connection")
			return
		}
	}
}

func Run() {
	http.HandleFunc("/ws", HandleWebsocket)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
