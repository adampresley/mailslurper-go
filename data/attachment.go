// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package data

/*
Attachment is a struct describing an attached file parsed from
a mail transmission.
*/
type Attachment struct {
	Id          int    `json:"id"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}
