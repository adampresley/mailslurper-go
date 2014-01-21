// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package smtp

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/adampresley/mailslurper/data"
)

// Constants representing the commands that an SMTP client will
// send during the course of communicating with our server.
const (
	DATA int = iota
	RCPT int = iota
	MAIL int = iota
	HELO int = iota
	RSET int = iota
	QUIT int = iota
)

// Constants for the various states the parser can be in. The parser
// always starts with STATE_START and will end in either STATE_QUIT
// if the transmission was successful, or STATE_ERROR if something
// bad happened along the way.
const (
	STATE_START       int = iota
	STATE_HEADER      int = iota
	STATE_DATA_HEADER int = iota
	STATE_BODY        int = iota
	STATE_QUIT        int = iota
	STATE_ERROR       int = iota
)

// Constants for parser buffer sizes and timeouts.
// CONN_TIMEOUT_MILLISECONDS is how many milliseconds to wait before
// attempting to read from the socket again. COMMAND_TIMEOUT_SETTINGS
// is how long to hold the socket open without recieving commands
// before closing with an error.
const (
	RECEIVE_BUFFER_LEN        = 1024
	CONN_TIMEOUT_MILLISECONDS = 5
	COMMAND_TIMEOUT_SECONDS   = 5
)

// This is a command map of SMTP command strings to their int
// representation. This is primarily used because there can
// be more than one command to do the same things. For example,
// a client can send "helo" or "ehlo" to initiate the handshake.
var Commands = map[string]int{
	"helo":      HELO,
	"ehlo":      HELO,
	"rcpt to":   RCPT,
	"mail from": MAIL,
	"send":      MAIL,
	"rset":      RSET,
	"quit":      QUIT,
	"data":      DATA,
}

// SMTP parser. The parser type keeps the current state of a parsing session,
// the socket connection handle, and finally collects all information into
// a MailItemStruct.
type Parser struct {
	State      int
	Connection net.Conn
	MailItem   data.MailItemStruct
}

/*
This function takes a command and the raw data read from the socket
connection and executes the correct handler function to process
the data and potentially respond to the client to continue SMTP negotiations.
*/
func (parser *Parser) CommandRouter(command int, input string) bool {
	var result bool
	var response string

	var date string
	var subject string
	var body string

	switch command {
	case HELO:
		result, response = parser.Process_HELO(strings.TrimSpace(input))
		fmt.Println(response)
		return result

	case MAIL:
		result, response = parser.Process_MAIL(strings.TrimSpace(input))
		if result == false {
			fmt.Println("An error occurred processing the MAIL FROM command: ", response)
		} else {
			parser.MailItem.FromAddress = response
			fmt.Println("Mail from: ", parser.MailItem.FromAddress)
		}

		return result

	case RCPT:
		result, response = parser.Process_RCPT(strings.TrimSpace(input))
		if result == false {
			fmt.Println("An error occurred process the RCPT TO command: ", response)
		} else {
			parser.MailItem.ToAddresses = append(parser.MailItem.ToAddresses, response)
			fmt.Printf("Mail to: %s\n", response)
		}

		return result

	case DATA:
		result, response, date, subject, body = parser.Process_DATA(strings.TrimSpace(input))
		if result == false {
			fmt.Println("An error occurred while reading the DATA chunk: ", response)
		} else {
			parser.MailItem.Body = body
			parser.MailItem.Subject = subject
			parser.MailItem.DateSent = date
			parser.MailItem.XMailer = "MailSlurper!"
		}

		return result

	default:
		return true
	}
}

/*
Takes a string and returns the integer command representation. For example
if the string contains "DATA" then the value 1 (the constant DATA) will be returned.
*/
func (parser *Parser) ParseCommand(line string) int {
	result := -1

	for key, value := range Commands {
		if strings.Index(strings.ToLower(line), key) > -1 {
			result = value
			break
		}
	}

	return result
}

/*
Function to process the HELO and EHLO SMTP commands. This command
responds to clients with a 250 greeting code and returns success
or false and an error message (if any).
*/
func (parser *Parser) Process_HELO(line string) (bool, string) {
	lowercaseLine := strings.ToLower(line)

	commandCheck := (strings.Index(lowercaseLine, "helo") + 1) + (strings.Index(lowercaseLine, "ehlo") + 1)
	if commandCheck < 0 {
		return false, "Invalid command"
	}

	split := strings.Split(line, " ")
	if len(split) < 2 {
		return false, "HELO command format is invalid"
	}

	result, _ := parser.SendResponse("250 Hello. How very nice to meet you!")
	if result != true {
		return false, "Error writing to connection stream in response to HELO"
	}

	return true, ""
}

/*
Function to process the MAIL FROM command (constant MAIL). This command
will respond to clients with 250 Ok response and returns true/false for success
and a string containing the sender's address.
*/
func (parser *Parser) Process_MAIL(line string) (bool, string) {
	commandCheck := strings.Index(strings.ToLower(line), "mail from")
	if commandCheck < 0 {
		return false, "Invalid command"
	}

	split := strings.Split(line, ":")
	if len(split) < 2 {
		return false, "MAIL FROM command format is invalid"
	}

	from := strings.Join(split[1:], "")

	result, _ := parser.SendOkResponse()
	if result != true {
		return false, "Error writing to connection stream in response to MAIL FROM"
	}

	return true, strings.TrimSpace(from)
}

/*
Function to process the RCPT TO command (constant RCPT). This command
will respond to clients with a 250 Ok response and returns true/false for
success and a string containing the recipients address. Note that a client
can send one or more RCPT TO commands.
*/
func (parser *Parser) Process_RCPT(line string) (bool, string) {
	commandCheck := strings.Index(strings.ToLower(line), "rcpt to")
	if commandCheck < 0 {
		return false, "Invalid command"
	}

	split := strings.Split(line, ":")
	if len(split) < 2 {
		return false, "RCPT TO command format is invalid"
	}

	to := strings.Join(split[1:], "")

	result, _ := parser.SendOkResponse()
	if result != true {
		return false, "Error writing to connection stream in response to RCPT TO"
	}

	return true, strings.TrimSpace(to)
}

/*
Function to process the DATA command (constant DATA). When a client sends the DATA
command there are three parts to the transmission content. Before this data
can be processed this function will tell the client how to terminate the DATA block.
We are asking clients to terminate with "\r\n.\r\n".

The first part is a set of header lines. Each header line is a header key (name), followed
by a colon, followed by the value for that header key. For example a header key might
be "Subject" with a value of "Testing Mail!".

After the header section there should be two sets of carriage return/line feed characters.
This signals the end of the header block and the start of the message body.

Finally when the client sends the "\r\n.\r\n" the DATA transmission portion is complete.
This function will return the following items.

	1. True/false for success
	2. Error or success message
	3. Mail date sent header
	4. Mail subject header
	5. Mail message body
*/
func (parser *Parser) Process_DATA(line string) (bool, string, string, string, string) {
	var dataBuffer bytes.Buffer

	commandCheck := strings.Index(strings.ToLower(line), "data")
	if commandCheck < 0 {
		return false, "Invalid command", "", "", ""
	}

	parser.SendResponse("354 End data with <CR><LF>.<CR><LF>")
	parser.State = STATE_HEADER

	for {
		dataResponse := parser.ReadChunk()

		terminatorPos := strings.Index(dataResponse, "\r\n.\r\n")
		if terminatorPos <= -1 {
			dataBuffer.WriteString(dataResponse)
		} else {
			dataBuffer.WriteString(dataResponse[0:terminatorPos])
			break
		}
	}

	/*
	 * Split the DATA content by CRLF CRLF. The first item will be the data
	 * headers. Everything past that is body/message.
	 */
	headerBodySplit := strings.Split(dataBuffer.String(), "\r\n\r\n")
	if len(headerBodySplit) < 2 {
		panic("Expected DATA block to contain a header section and a body section")
	}

	/*
	 * Parse the header content
	 */
	parser.State = STATE_DATA_HEADER
	headerData := parseDataHeader(headerBodySplit[0])

	parser.SendOkResponse()
	return true, "Success", headerData["date"], headerData["subject"], strings.Join(headerBodySplit[1:], "\r\n\r\n")
}

/*
This function reads the raw data from the socket connection to our client. This will
read on the socket until there is nothing left to read and an error is generated.
This method blocks the socket for the number of milliseconds defined in CONN_TIMEOUT_MILLISECONDS.
It then records what has been read in that time, then blocks again until there is nothing left on
the socket to read. The final value is stored and returned as a string.
*/
func (parser *Parser) ReadChunk() string {
	var raw bytes.Buffer
	var bytesRead int

	bytesRead = 1

	for bytesRead > 0 {
		parser.Connection.SetReadDeadline(time.Now().Add(time.Millisecond * CONN_TIMEOUT_MILLISECONDS))

		buffer := make([]byte, RECEIVE_BUFFER_LEN)
		bytesRead, err := parser.Connection.Read(buffer)

		if err != nil {
			break
		}

		if bytesRead > 0 {
			raw.WriteString(string(buffer[:bytesRead]))
		}
	}

	return raw.String()
}

/*
This is the main entry function called when a new client connection is established.
It begins by sending a 220 welcome message to the client to indicate we are ready
to communicate. From here we initialize a parser and blank MailItemStruct to hold
the data we recieve. Once we recieve the quit command we close out.

A parsing session can end with either a STATE_QUIT if all was successful, or
a STATE_ERROR if there was a problem.
*/
func (parser *Parser) Run() {
	var raw string
	var command int
	var commandRouterResult bool

	parser.SendResponse("220 Welcome to MailSlurper!")
	fmt.Println("Reading data from client connection...")

	/*
	 * Initialize the recipient list to handle up to 20 items to start.
	 */
	parser.MailItem.ToAddresses = make([]string, 0, 20)

	/*
	 * Read from the connection until we receive a QUIT command
	 * or some critical error occurs and we force quit.
	 */
	parser.State = STATE_START
	startTime := time.Now()

	for parser.State != STATE_QUIT && parser.State != STATE_ERROR {
		raw = parser.ReadChunk()
		command = parser.ParseCommand(raw)

		if command == QUIT {
			parser.State = STATE_QUIT
			fmt.Println("Closing connection.")
		} else {
			commandRouterResult = parser.CommandRouter(command, raw)

			if commandRouterResult != true {
				parser.State = STATE_ERROR
				fmt.Println("Error occured executing command ", command)
			}
		}

		if int(time.Since(startTime).Seconds()) > COMMAND_TIMEOUT_SECONDS {
			parser.State = STATE_ERROR
		}
	}

	parser.SendClosingResponse()
}

/*
Function to tell a client that we are done communicating. This sends
a 221 response. It returns true/false for success and a string
with any response.
*/
func (parser *Parser) SendClosingResponse() (bool, string) {
	result, response := parser.SendResponse("221 Bye")
	return result, response
}

/*
Function to tell a client that we recieved the last communication successfully
and are ready to get our next command. This sends a 250 response. It returns true/false for success and a string
with any response.
*/
func (parser *Parser) SendOkResponse() (bool, string) {
	result, response := parser.SendResponse("250 Ok")
	return result, response
}

/*
Function to send a response to a client connection. It returns true/false for success and a string
with any response.
*/
func (parser *Parser) SendResponse(resp string) (bool, string) {
	result := true
	response := ""

	_, err := parser.Connection.Write([]byte(string(resp + "\r\n")))
	if err != nil {
		result = false
		response = err.Error()
	}

	return result, response
}

/*
Takes a string block and parses header items. Returns a map of
those parsed headers, where the key is the header name and the
value is the header value.
*/
func parseDataHeader(headerLines string) map[string]string {
	splitHeader := strings.Split(headerLines, "\r\n")
	numLines := len(splitHeader)

	result := make(map[string]string, numLines)

	result["xmailer"] = "MailSlurper!"
	result["date"] = ""
	result["subject"] = ""

	for index := 0; index < numLines; index++ {
		splitHeaderItem := strings.Split(splitHeader[index], ":")

		if strings.ToLower(splitHeaderItem[0]) == "date" {
			result["date"] = parseDateTime(strings.TrimSpace(strings.Join(splitHeaderItem[1:], ":")))
			fmt.Println("Date: ", result["date"])
		}

		if strings.ToLower(splitHeaderItem[0]) == "subject" {
			result["subject"] = strings.TrimSpace(strings.Join(splitHeaderItem[1:], ""))
			fmt.Println("Subject: ", result["subject"])
		}
	}

	return result
}

/*
Takes a date/time string and attempts to parse it and return a newly formatted
date/time that looks like YYYY-MM-DD HH:MM:SS
*/
func parseDateTime(dateString string) string {
	outputForm := "2006-01-02 15:04:05"
	firstForm := "Mon, 02 Jan 2006 15:04:05 -0700 MST"
	secondForm := "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"

	result := ""

	t, err := time.Parse(firstForm, dateString)
	if err != nil {
		t, err = time.Parse(secondForm, dateString)
		if err != nil {
			fmt.Printf("Error parsing date: %s\n", err)
			result = dateString
		} else {
			result = t.Format(outputForm)
		}
	} else {
		result = t.Format(outputForm)
	}

	return result
}
