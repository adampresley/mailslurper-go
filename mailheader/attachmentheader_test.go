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

func TestParseAttachmentHeader(t *testing.T) {
	Convey("Parsing attachment headers", t, func() {
		Convey("Content and headers not separated by two CRLFs returns error", func() {
			attachmentHeader := &AttachmentHeader{}
			contents := "Content-Disposition: attachment\r\n"+
				"Content-Transfer-Encoding: UTF-8\r\n"+
				"Content-Type: image/png\r\n"+
				"MIME-Version: 1.0\r\n"+
				"Contents!!ABCDEFG"

			expected := errors.New(messages.ERROR_INVALID_ATTACHMENT_BLOCK)
			actual := attachmentHeader.ParseAttachmentHeader(contents)

			So(actual, ShouldResemble, expected)
		})

		Convey("Can get basic headers", func() {
			contents := "Content-Disposition: attachment\r\n"+
				"Content-Transfer-Encoding: UTF-8\r\n"+
				"Content-Type: image/png\r\n"+
				"MIME-Version: 1.0\r\n"+
				"\r\n"+
				"Contents!!ABCDEFG"

			expected := &AttachmentHeader{
				ContentType:             "image/png",
				MIMEVersion:             "1.0",
				ContentTransferEncoding: "UTF-8",
				ContentDisposition:      "attachment",
				FileName:                "",
				Body:                    "Contents!!ABCDEFG",
			}

			actual := &AttachmentHeader{}
			actual.ParseAttachmentHeader(contents)

			So(actual, ShouldResemble, expected)
		})

		Convey("Can get file name from Content-Disposition", func() {
			contents := "Content-Disposition: attachment; filename=\"bob.png\"\r\n"+
				"Content-Transfer-Encoding: UTF-8\r\n"+
				"Content-Type: image/png\r\n"+
				"MIME-Version: 1.0\r\n"+
				"\r\n"+
				"Contents!!ABCDEFG"

			expected := &AttachmentHeader{
				ContentType:             "image/png",
				MIMEVersion:             "1.0",
				ContentTransferEncoding: "UTF-8",
				ContentDisposition:      "attachment",
				FileName:                "bob.png",
				Body:                    "Contents!!ABCDEFG",
			}

			actual := &AttachmentHeader{}
			actual.ParseAttachmentHeader(contents)

			So(actual, ShouldResemble, expected)
		})

		Convey("Can get file name from Content-Type", func() {
			contents := "Content-Disposition: attachment\r\n"+
				"Content-Transfer-Encoding: UTF-8\r\n"+
				"Content-Type: image/png; name=\"bob.png\"\r\n"+
				"MIME-Version: 1.0\r\n"+
				"\r\n"+
				"Contents!!ABCDEFG"

			expected := &AttachmentHeader{
				ContentType:             "image/png",
				MIMEVersion:             "1.0",
				ContentTransferEncoding: "UTF-8",
				ContentDisposition:      "attachment",
				FileName:                "bob.png",
				Body:                    "Contents!!ABCDEFG",
			}

			actual := &AttachmentHeader{}
			actual.ParseAttachmentHeader(contents)

			So(actual, ShouldResemble, expected)
		})

	})
}