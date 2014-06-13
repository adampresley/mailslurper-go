// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"fmt"
	"log"
	"net"
)

// Represents an SMTP server with an address and connection handle.
type Server struct {
	Address          string
	ConnectionHandle net.Listener
}

/*
Establishes a listening connection to a socket on an address. This will
set the connection handle on our Server struct.
*/
func (s *Server) Connect() {
	handle, err := net.Listen("tcp", s.Address)
	if err != nil {
		panic(fmt.Sprintf("Error while setting up SMTP listener"))
	}

	s.ConnectionHandle = handle
	log.Println("SMTP listener setup at ", s.Address)
}

/*
Closes a socket connection in an Server object. Most likely used in a defer call.

Example:
	smtp := Server.Server { Address: "127.0.0.1:8000" }
	defer smtp.Close()
*/
func (s *Server) Close() {
	s.ConnectionHandle.Close()
}

/*
This function starts the process of handling SMTP client connections.
The first order of business is to setup a channel for writing
parsed mails, in the form of MailItemStruct variables, to our
SQLite database. A goroutine is setup to listen on that
channel and handles storage.

Meanwhile this method will loop forever and wait for client connections (blocking).
When a connection is recieved a goroutine is started to create a new MailItemStruct
and parser and the parser process is started. If the parsing is successful
the MailItemStruct is added to the database writing channel.
*/
func (s *Server) ProcessRequests() {
	/*
	 * Setup a channel for communicating writing mail items to our
	 * data storage. Start listening for write requests.
	 */
	dbWriteChannel := make(chan MailItemStruct, 100)
	go Storage.StartWriteListener(dbWriteChannel)

	/*
	 * Now start accepting connections for SMTP
	 */
	for {
		connection, err := s.ConnectionHandle.Accept()
		if err != nil {
			log.Panicf("Error while accepting SMTP requests: %s", err)
		}

		go func(c net.Conn, dbWriter chan MailItemStruct) {
			defer c.Close()

			/*
			 * Create a package that starts processing SMTP commands
			 * unil it is time to close the connection
			 */
			mailItem := MailItemStruct{}

			parser := Parser{
				State:      STATE_START,
				Connection: c,
				MailItem:   mailItem,
			}

			parser.Run()

			if parser.State == STATE_QUIT {
				log.Println("Writing mail item to database and websocket...")
				dbWriter <- parser.MailItem
			} else {
				log.Println("An error occurred during mail transmission and data will not be written.")
			}
		}(connection, dbWriteChannel)
	}
}
