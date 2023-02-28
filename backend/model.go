package main

import "github.com/gorilla/websocket"

type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte // Channel for sending messages to the client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type ClientManager struct {
	clients    map[*Client]bool // Map of connected clients
	broadcast  chan []byte      // Channel for broadcasting messages to all clients
	register   chan *Client     // Channel for registering new clients
	unregister chan *Client     // Channel for unregistering clients
}
