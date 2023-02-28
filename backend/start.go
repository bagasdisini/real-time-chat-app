package main

import "encoding/json"

// start is a method of the ClientManager that runs in a separate goroutine and listens for incoming messages on its channels
func (manager *ClientManager) start() {
	for {
		select {
		case client := <-manager.register:
			// Add new client to the map of connected clients
			manager.clients[client] = true
			// Send a message to all clients that a new client has connected
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, client)
		case client := <-manager.unregister:
			if _, ok := manager.clients[client]; ok {
				// Close the channel used to send messages to this client
				close(client.Send)
				// Remove the client from the map of connected clients
				delete(manager.clients, client)
				// Send a message to all clients that a client has disconnected
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, client)
			}
		case message := <-manager.broadcast:
			// Send the incoming message to all connected clients
			for client := range manager.clients {
				if client != nil {
					select {
					case client.Send <- message:
					default:
						// If we can't send the message to this client, it means it is no longer connected, so remove it
						close(client.Send)
						delete(manager.clients, client)
					}
				}
			}
		}
	}
}
