package smtp

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type AttachmentHeader struct {
	Contents                string
	Boundary                string
	ContentType             string
	MIMEVersion             string
	ContentTransferEncoding string
	ContentDisposition      string
	FileName                string
}

type MailHeader struct {
	Contents    string
	ContentType string
	Boundary    string
	MIMEVersion string
	Subject     string
	Date        string
	XMailer     string
}

/*
Parses a set of mail headers. Splits lines up and figures out what
header data goes into what structure key. Most headers follow this format:

Header-Name: Some value here\r\n

However some headers, such as Content-Type, may have additiona information,
especially when the content type is a multipart and there are attachments.
Then it can look like this:

Content-Type: multipart/mixed; boundary="==abcsdfdfd=="\r\n
*/
func (this *MailHeader) Parse() {
	var key string

	this.XMailer = "MailSlurper!"
	this.Boundary = ""

	/*
	 * Unfold and split the header into lines. Loop over each line
	 * and figure out what headers are present. Store them.
	 * Sadly some headers require special processing.
	 */
	this.Contents = unfoldHeaders(this.Contents)
	splitHeader := strings.Split(this.Contents, "\r\n")
	numLines := len(splitHeader)

	for index := 0; index < numLines; index++ {
		splitItem := strings.Split(splitHeader[index], ":")
		key = splitItem[0]

		switch strings.ToLower(key) {
		case "content-type":
			contentType := strings.Join(splitItem[1:], "")
			contentTypeSplit := strings.Split(contentType, ";")

			this.ContentType = strings.TrimSpace(contentTypeSplit[0])
			fmt.Println("Content-Type: ", this.ContentType)

			/*
			 * Check to see if we have a boundary marker
			 */
			if len(contentTypeSplit) > 1 {
				contentTypeRightSide := strings.Join(contentTypeSplit[1:], ";")

				if strings.Contains(strings.ToLower(contentTypeRightSide), "boundary") {
					boundarySplit := strings.Split(contentTypeRightSide, "=")
					this.Boundary = strings.Replace(strings.Join(boundarySplit[1:], "="), "\"", "", -1)
					fmt.Println("Boundary: ", this.Boundary)
				}
			}

		case "date":
			this.Date = parseDateTime(strings.Join(splitItem[1:], ":"))
			fmt.Println("Date: ", this.Date)

		case "mime-version":
			this.MIMEVersion = strings.TrimSpace(strings.Join(splitItem[1:], ""))
			fmt.Println("MIME-Version: ", this.MIMEVersion)

		case "subject":
			this.Subject = strings.TrimSpace(strings.Join(splitItem[1:], ""))
			fmt.Println("Subject: ", this.Subject)
		}
	}
}

/*
Parses a set of attachment headers. Splits lines up and figures out what
header data goes into what structure key. Most headers follow this format:

Header-Name: Some value here\r\n
*/
func (this *AttachmentHeader) Parse() {
	var key string

	/*
	 * Unfodl and split the header into lines. Loop over each line
	 * and figure out what headers are present. Store them.
	 * Sadly some headers require special processing.
	 */
	this.Contents = unfoldHeaders(this.Contents)
	splitHeader := strings.Split(this.Contents, "\r\n")
	numLines := len(splitHeader)

	for index := 0; index < numLines; index++ {
		splitItem := strings.Split(splitHeader[index], ":")
		key = splitItem[0]

		switch strings.ToLower(key) {
		case "content-disposition":
			contentDisposition := strings.TrimSpace(strings.Join(splitItem[1:], ""))
			contentDispositionSplit := strings.Split(contentDisposition, ";")

			if len(contentDispositionSplit) < 2 {
				this.ContentDisposition = contentDisposition
			} else {
				this.ContentDisposition = strings.TrimSpace(contentDispositionSplit[0])
				contentDispositionRightSide := strings.TrimSpace(strings.Join(contentDispositionSplit[1:], ";"))

				/*
				 * See if we have an attachment and filename
				 */
				if strings.ToLower(this.ContentDisposition) == "attachment" {
					filenameSplit := strings.Split(contentDispositionRightSide, "=")
					this.FileName = strings.Replace(strings.Join(filenameSplit[1:], "="), "\"", "", -1)
					fmt.Println("Attachment File Name: ", this.FileName)
				}
			}

			fmt.Println("Content-Disposition: ", this.ContentDisposition)

		case "content-transfer-encoding":
			this.ContentTransferEncoding = strings.TrimSpace(strings.Join(splitItem[1:], ""))
			fmt.Println("Content-Transfer-Encoding: ", this.ContentTransferEncoding)

		case "content-type":
			this.ContentType = strings.TrimSpace(strings.Join(splitItem[1:], ""))
			fmt.Println("Content-Type: ", this.ContentType)

		case "mime-version":
			this.MIMEVersion = strings.TrimSpace(strings.Join(splitItem[1:], ""))
			fmt.Println("MIME-Version: ", this.MIMEVersion)
		}
	}
}

/*
Takes a date/time string and attempts to parse it and return a newly formatted
date/time that looks like YYYY-MM-DD HH:MM:SS
*/
func parseDateTime(dateString string) string {
	outputForm := "2006-01-02 15:04:05"
	firstForm := "Mon, 02 Jan 2006 15:04:05 -0700 MST"
	secondForm := "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"

	dateString = strings.TrimSpace(dateString)
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

/*
The RFC-2822 defines "folding" as the process of breaking up large
header lines into multiple lines. Long Subject lines or Content-Type
lines (with boundaries) sometimes do this. This function will "unfold"
them into a single line.
*/
func unfoldHeaders(contents string) string {
	headerUnfolderRegex := regexp.MustCompile("(.*?)\r\n\\s{1}(.*?)\r\n")
	return headerUnfolderRegex.ReplaceAllString(contents, "$1 $2\r\n")
}
