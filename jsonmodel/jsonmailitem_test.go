// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jsonmodel

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJsonMailItems(t *testing.T) {
	Convey("JSON Mail Items", t, func() {
		mailItem := JsonMailItem{
			Id      :    1,
			DateSent:    "2014-01-01 00:00:00",
			FromAddress: "a@b.com",
			ToAddresses: []string { "b@b.com" },
			Subject:     "Test",
			XMailer:     "Mailer",
			Body:        "Body",
			ContentType: "text/html",
		}

		Convey("Able to create a new JSON Mail Item", func() {
			actual := NewJsonMailItem(1, "2014-01-01 00:00:00", "a@b.com", []string { "b@b.com" }, "Test", "Mailer", "Body", "text/html")
			So(actual, ShouldResemble, mailItem)
		})

		Convey("Convert to JSON strings", func() {
			expected, _ := json.Marshal(mailItem)
			actual := mailItem.ToJson()

			So(string(actual), ShouldEqual, string(expected))
		})
	})
}