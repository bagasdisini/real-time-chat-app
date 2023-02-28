package main

// send is a helper method of the ClientManager that sends a message to all connected clients except one (the "ignore" client)
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for client := range manager.clients {
		if client != ignore {
			client.Send <- message
		}
	}
}
