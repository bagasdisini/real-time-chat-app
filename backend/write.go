package main

import "github.com/gorilla/websocket"

// write listens for outgoing messages to be sent to the socket.
// It sends the message to the socket and closes the socket if there's an error.
func (c *Client) write() {
	defer func() {
		// Close the socket before exiting.
		err := c.Socket.Close()
		if err != nil {
			// If there's an error, ignore it and return.
			return
		}
	}()

	for {
		// Listen for outgoing messages to be sent to the socket.
		select {
		case message, ok := <-c.Send:
			if !ok {
				// If the channel is closed, send a close message to the socket and return.
				err := c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					// If there's an error, ignore it and return.
					return
				}
				return
			}

			// Send the message to the socket as a text message.
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				// If there's an error, ignore it and return.
				return
			}
		}
	}
}
