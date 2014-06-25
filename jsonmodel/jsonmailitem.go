// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jsonmodel

import (
	"encoding/json"
)

type JsonMailItem struct {
	Id              int              `json:"id"`
	DateSent        string           `json:"dateSent"`
	FromAddress     string           `json:"fromAddress"`
	ToAddresses     []string         `json:"toAddresses"`
	Subject         string           `json:"subject"`
	XMailer         string           `json:"xmailer"`
	Body            string           `json:"body"`
	ContentType     string           `json:"contentType"`
	AttachmentCount int              `json:"attachmentCount"`
	Attachments     []JsonAttachment `json:"attachments"`
}

func NewJsonMailItem(id int, dateSent, fromAddress string, toAddresses []string, subject, xMailer, body, contentType string) JsonMailItem {
	return JsonMailItem{
		Id      :    id,
		DateSent:    dateSent,
		FromAddress: fromAddress,
		ToAddresses: toAddresses,
		Subject:     subject,
		XMailer:     xMailer,
		Body:        body,
		ContentType: contentType,
	}
}

func (this *JsonMailItem) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}