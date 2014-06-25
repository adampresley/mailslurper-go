// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailheader

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseDateTime(t *testing.T) {
	Convey("Date/Time Parsing", t, func() {
		dateTime1 := "Tue, 01 Jan 2014 12:01:15 -0600 CST"
		dateTime2 := "Tue, 01 Jan 2014 12:01:15 -0600 (CST)"
		dateTime3 := "Tue, 1 Jan 2014 12:01:15 -0600 (CST)"
		invalidDate := "Tuesday 1st January 2014 at 12:01pm"

		Convey("Can parse basic long format", func() {
			expected := "2014-01-01 12:01:15"
			actual := parseDateTime(dateTime1)
			So(actual, ShouldEqual, expected)
		})

		Convey("Can parse dates with parenthensis around timezone", func() {
			expected := "2014-01-01 12:01:15"
			actual := parseDateTime(dateTime2)
			So(actual, ShouldEqual, expected)
		})

		Convey("Can parse dates with single digit days", func() {
			expected := "2014-01-01 12:01:15"
			actual := parseDateTime(dateTime3)
			So(actual, ShouldEqual, expected)
		})

		Convey("Invalid dates return original string", func() {
			expected := "Tuesday 1st January 2014 at 12:01pm"
			actual := parseDateTime(invalidDate)
			So(actual, ShouldEqual, expected)
		})
	})
}

func TestUnfoldHeaders(t *testing.T) {
	Convey("Unfolding Headers", t, func() {
		Convey("On a line with folding unfolds to single line", func() {
			expected := "This is a test\r\n"
			actual := unfoldHeaders("This is\r\n\ta test\r\n")
			So(actual, ShouldEqual, expected)
		})

		Convey("On a line with no folding returns it just the way it was", func() {
			expected := "This is a test\r\n"
			actual := unfoldHeaders("This is a test\r\n")
			So(actual, ShouldEqual, expected)
		})
	})
}