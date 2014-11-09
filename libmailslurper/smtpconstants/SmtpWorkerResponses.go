package smtpconstants

const (
	SMTP_CRLF                   string = "\r\n"
	SMTP_DATA_TERMINATOR        string = "\r\n.\r\n"
	SMTP_WELCOME_MESSAGE        string = "220 Welcome to MailSlurper!"
	SMTP_CLOSING_MESSAGE        string = "221 Bye"
	SMTP_OK_MESSAGE             string = "250 Ok"
	SMTP_DATA_RESPONSE_MESSAGE  string = "354 End data with <CR><LF>.<CR><LF>"
	SMTP_HELLO_RESPONSE_MESSAGE string = "250 Hello. How very nice to meet you!"
)