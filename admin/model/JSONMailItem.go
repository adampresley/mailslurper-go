// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package model

type JSONAttachment struct {
	Id       int    `json:"id"`
	FileName string `json:"fileName"`
}

type JSONMailItem struct {
	Id              int              `json:"id"`
	DateSent        string           `json:"dateSent"`
	FromAddress     string           `json:"fromAddress"`
	ToAddresses     []string         `json:"toAddresses"`
	Subject         string           `json:"subject"`
	XMailer         string           `json:"xmailer"`
	Body            string           `json:"body"`
	ContentType     string           `json:"contentType"`
	AttachmentCount int              `json:"attachmentCount"`
	Attachments     []JSONAttachment `json:"attachments"`
}
