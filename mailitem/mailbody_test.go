// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailitem

import (
	"errors"
	"testing"

	"github.com/adampresley/mailslurper/messages"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParseMailBody(t *testing.T) {
	Convey("ParseMailBody handles different types of bodies and attachments", t, func() {
		Convey("Content and headers not separated by two CRLFs returns error", func() {
			mailBody := &MailBody{}
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: text/html; boundary=\"--abcd\"\r\n"+
				"Start of contents!"

			expected := errors.New(messages.ERROR_INVALID_DATA_BLOCK)
			actual := mailBody.ParseMailBody(contents, "--abcd")

			So(actual, ShouldResemble, expected)
		})

		Convey("Simple text mail (no boundary) puts body into Text body", func() {
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: text/html; boundary=\"--abcd\"\r\n"+
				"Start of contents!"

			expected := &MailBody{
				TextBody: "Start of contents!",
			}
			actual := &MailBody{}
			actual.ParseMailBody(contents, "")

			So(actual, ShouldResemble, expected)
		})

		Convey("Multipart mail (text and HTML) returns body bodies", func() {
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: multipart/mixed; boundary=\"--abcd\"\r\n"+
				"\r\n"+
				"--abcd\r\n"+
				"Content-Type: text/plain\r\n"+
				"\r\n"+
				"Start of contents!"+
				"--abcd\r\n"+
				"Content-Type: text/html\r\n"+
				"\r\n"+
				"<p>Start of contents!</p>\r\n"+
				"--abcd--"

			expected := &MailBody{
				TextBody: "Start of contents!",
				HTMLBody: "<p>Start of contents!</p>",
			}

			actual := &MailBody{}
			actual.ParseMailBody(contents, "")

			So(actual, ShouldResemble, expected)
		})
	})
}