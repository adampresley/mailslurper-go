// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jsonmodel

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonAttachments(t *testing.T) {
	Convey("JSON Attachments", t, func() {
		attachment := JsonAttachment{
			Id:       1,
			FileName: "bob.jpg",
		}

		Convey("Able to create a new JSON Attachment", func() {
			actual := NewJsonAttachment(1, "bob.jpg")
			So(actual, ShouldResemble, attachment)
		})

		Convey("Convert to JSON strings", func() {
			expected, _ := json.Marshal(attachment)
			actual := attachment.ToJson()

			So(string(actual), ShouldEqual, string(expected))
		})
	})
}