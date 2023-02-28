package main

import (
	"encoding/json"
	"log"
)

func (c *Client) read() {
	// Defer the closing of the connection until the end of the function
	defer func() {
		c.close()
	}()

	for {
		// Read a message from the socket
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			// Close the connection and exit the loop if there's an error reading a message
			c.close()
			break
		}

		// Handle the incoming message
		c.handleMessage(message)
	}
}

// Handle an incoming message
func (c *Client) handleMessage(message []byte) {
	// Marshal the message into JSON
	jsonMessage, err := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
	if err != nil {
		// Log an error if there's a problem marshaling the message
		log.Println("Error marshaling message:", err)
		return
	}

	// Broadcast the message to all clients
	clientsManager.broadcast <- jsonMessage
}

// Close the client connection
func (c *Client) close() {
	// Unregister the client from the manager
	clientsManager.unregister <- c

	// Close the socket connection
	err := c.Socket.Close()
	if err != nil {
		// Log an error if there's a problem closing the socket
		log.Println("Error closing socket:", err)
	}
}
