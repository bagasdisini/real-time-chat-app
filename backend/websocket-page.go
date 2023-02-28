package main

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func wsPage(res http.ResponseWriter, req *http.Request) {
	// Upgrade the HTTP connection to a WebSocket
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	// Create a new client and register it with the manager
	client := &Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}
	clientsManager.register <- client

	// Start a new goroutine to handle incoming messages from the client
	go client.read()
	// Start a new goroutine to handle outgoing messages to the client
	go client.write()
}
