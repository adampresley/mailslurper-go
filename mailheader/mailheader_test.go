// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailheader

import (
	"errors"
	"testing"

	"github.com/adampresley/mailslurper/messages"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMailHeader(t *testing.T) {
	Convey("NewMailHeader returns a MailHeader struct", t, func() {
		expected := &MailHeader{
			ContentType: "text/html",
			Boundary:    "--abcdefg",
			MIMEVersion: "1.0",
			Subject:     "Subject",
			Date:        "2014-01-01",
			XMailer:     "MailSlurper!",
		}

		actual := NewMailHeader("text/html", "--abcdefg", "1.0", "Subject", "2014-01-01", "MailSlurper!")
		So(actual, ShouldResemble, expected)
	})
}

func TestParseMailHeader(t *testing.T) {
	Convey("Parsing mail headers", t, func() {
		Convey("Content and headers not separated by two CRLFs returns error", func() {
			mailHeader := &MailHeader{}
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: text/html; boundary=\"--abcd\"\r\n"+
				"Start of contents!"

			expected := errors.New(messages.ERROR_INVALID_DATA_BLOCK)
			actual := mailHeader.ParseMailHeader(contents)

			So(actual, ShouldResemble, expected)
		})

		Convey("Can get basic headers", func() {
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: text/html\r\n"+
				"\r\n"+
				"Start of contents!"

			expected := &MailHeader{
				ContentType: "text/html",
				Boundary:    "",
				MIMEVersion: "1.0",
				Subject:     "Testing",
				Date:        "2014-06-22 22:29:05",
				XMailer:     "MailSlurper!",
			}

			actual := &MailHeader{}
			actual.ParseMailHeader(contents)

			So(actual, ShouldResemble, expected)
		})

		Convey("Can get boundary", func() {
			contents := "From: a@b.com\r\n"+
				"To: b@c.com\r\n"+
				"Subject: Testing\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Date: Sun, 22 Jun 2014 22:29:05 -0600 CST\r\n"+
				"Content-Type: text/html; boundary=\"--abcd\"\r\n"+
				"\r\n"+
				"Start of contents!"

			expected := &MailHeader{
				ContentType: "text/html",
				Boundary:    "--abcd",
				MIMEVersion: "1.0",
				Subject:     "Testing",
				Date:        "2014-06-22 22:29:05",
				XMailer:     "MailSlurper!",
			}

			actual := &MailHeader{}
			actual.ParseMailHeader(contents)

			So(actual, ShouldResemble, expected)
		})
	})
}