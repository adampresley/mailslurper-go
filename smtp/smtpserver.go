// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"fmt"
	"log"
	"net"

	"github.com/adampresley/mailslurper/mailitem"
	"github.com/adampresley/mailslurper/profiling"
)

/*
Establishes a listening connection to a socket on an address. This will
set the connection handle on our Server struct.
*/
func Connect(address string) (net.Listener, error) {
	handle, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return handle, nil
}

/*
Closes a socket connection in an Server object. Most likely used in a defer call.
*/
func Close(handle net.Listener) {
	handle.Close()
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
func ProcessRequests(handle net.Listener, databaseWriteChannel chan mailitem.MailItem) error {
	/*
	 * Start accepting connections for SMTP
	 */
	for {
		connection, err := handle.Accept()
		if err != nil {
			return err
		}

		go func(smtpConnection net.Conn, dbWriter chan mailitem.MailItem) {
			defer c.Close()

			/*
			 * Create a package that starts processing SMTP commands
			 * unil it is time to close the connection
			 */
			mailItem := mailitem.MailItem{}

			parser := Parser{
				State:      STATE_START,
				Connection: smtpConnection,
				MailItem:   mailItem,
			}

			profiling.Timer.Step("Parse mail item")
			parser.Run()

			if parser.State == STATE_QUIT {
				log.Println("Writing mail item to database and websocket...")
				dbWriter <- parser.MailItem
			} else {
				log.Println("An error occurred during mail transmission and data will not be written.")
			}
		}(connection, databaseWriteChannel)
	}
}
