// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailheader

import (
	"log"
	"regexp"
	"strings"
	"time"
)

/*
Takes a date/time string and attempts to parse it and return a newly formatted
date/time that looks like YYYY-MM-DD HH:MM:SS
*/
func parseDateTime(dateString string) string {
	outputForm := "2006-01-02 15:04:05"
	firstForm := "Mon, 02 Jan 2006 15:04:05 -0700 MST"
	secondForm := "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"
	thirdForm := "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"

	dateString = strings.TrimSpace(dateString)
	result := ""

	t, err := time.Parse(firstForm, dateString)
	if err != nil {
		t, err = time.Parse(secondForm, dateString)
		if err != nil {
			t, err = time.Parse(thirdForm, dateString)
			if err != nil {
				log.Printf("Error parsing date: %s\n", err)
				result = dateString
			} else {
				result = t.Format(outputForm)
			}
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
