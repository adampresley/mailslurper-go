// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/adampresley/mailslurper/settings"
	"github.com/adampresley/mailslurper/smtp"
)

/*
Controller used to download an attachment.
*/
func DownloadAttachment(writer http.ResponseWriter, request *http.Request) {
	attachmentId, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "ID provided is invalid", 500)
		return
	}

	attachment := smtp.Storage.GetAttachment(attachmentId)

	data, err := base64.StdEncoding.DecodeString(attachment["content"])
	if err != nil {
		http.Error(writer, "Cannot decode attachment", 500)
		return
	}

	reader := bytes.NewReader(data)
	http.ServeContent(writer, request, attachment["fileName"], time.Now(), reader)
}

/*
Controller for the home page
*/
func Home(writer http.ResponseWriter, request *http.Request) {
	settings.Config.RenderView(writer, "index")
}

/*
This function handles a web GET request for "/mails". It queries the storage
engine for all mail items, sets the content type header to text/json, and
returns a JSON-serialized array of mail data.
*/
func GetMailCollection(writer http.ResponseWriter, request *http.Request) {
	mailItems := smtp.Storage.GetMails()
	json, _ := json.Marshal(mailItems)
	settings.Config.WriteJson(writer, json)
}

func GetMailItem(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "ID provided is invalid", 500)
		return
	}

	mailItem := smtp.Storage.GetMail(id)
	json, _ := json.Marshal(mailItem)
	settings.Config.WriteJson(writer, json)
}
