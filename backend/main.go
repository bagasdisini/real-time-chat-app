package main

import (
	"fmt"
	"net/http"
)

// Initialize a client manager
var clientsManager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func main() {
	fmt.Println("Apps started!")

	// Start the client manager
	go clientsManager.start()

	// Create a WebSocket endpoint
	http.HandleFunc("/ws", wsPage)

	// Listen for incoming HTTP requests
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		return
	}
}
