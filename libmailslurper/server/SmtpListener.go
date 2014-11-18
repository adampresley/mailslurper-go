// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package server

import (
	"log"
	"net"

	"github.com/adampresley/mailslurper/libmailslurper/model/mailitem"
)

/*
Establishes a listening connection to a socket on an address. This will
return a net.Listener handle.
*/
func SetupSmtpServerListener(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}

/*
Closes a socket connection in an Server object. Most likely used in a defer call.
*/
func CloseSmtpServerListener(handle net.Listener) {
	handle.Close()
}

/*
This function starts the process of handling SMTP client connections.
The first order of business is to setup a channel for writing
parsed mails, in the form of MailItemStruct variables, to our
database. A goroutine is setup to listen on that
channel and handles storage.

Meanwhile this method will loop forever and wait for client connections (blocking).
When a connection is recieved a goroutine is started to create a new MailItemStruct
and parser and the parser process is started. If the parsing is successful
the MailItemStruct is added to the database writing channel.
*/
func Dispatcher(serverPool *ServerPool, handle net.Listener, receiver chan mailitem.MailItem) {
	/*
	 * Now start accepting connections for SMTP
	 */
	for {
		connection, err := handle.Accept()
		if err != nil {
			log.Panicf("ERROR - Error while accepting SMTP requests: %s", err)
		}

		smtpWorker, err := serverPool.GetAvailableWorker(connection, receiver)
		if err != nil {
			log.Println("ERROR -", err)
			continue
		}

		smtpWorker.Work()
	}
}
