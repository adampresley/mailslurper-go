// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package websockets

import (
	"net/http"

	"github.com/adampresley/mailslurper/data"
	"github.com/gorilla/websocket"
)

// Structure for tracking and working with websockets
type WebsocketConnection struct {
	// Websocket connection handle
	WS *websocket.Conn

	// Buffered channel for outbound messages
	SendChannel chan data.MailItemStruct
}

var WebsocketConnections map[*WebsocketConnection]bool = make(map[*WebsocketConnection]bool)

/*
This function takes a MailItemStruct and sends it to all open websockets.
*/
func BroadcastMessageToWebsockets(message data.MailItemStruct) {
	for connection := range WebsocketConnections {
		connection.SendChannel <- message
	}
}

/*
This function handles the handshake for our websocket connection.
It sets up a goroutine to handle sending MailItemStructs to the
other side.
*/
func WebsocketHandler(writer http.ResponseWriter, request *http.Request) {
	ws, err := websocket.Upgrade(writer, request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(writer, "Invalid handshake", 400)
		return
	} else if err != nil {
		return
	}

	/*
	 * Create a new websocket connection struct and add it's pointer
	 * address to our web socket tracking map.
	 */
	connection := &WebsocketConnection{WS: ws, SendChannel: make(chan data.MailItemStruct, 256)}
	WebsocketConnections[connection] = true
	defer destroyConnection(connection)

	for {
		for message := range connection.SendChannel {
			err := connection.WS.WriteJSON(message)
			if err != nil {
				break
			}
		}
	}

	connection.WS.Close()
}

func destroyConnection(connection *WebsocketConnection) {
	// Remove the connection from our map, and close its channel
	delete(WebsocketConnections, connection)
	close(connection.SendChannel)
}
