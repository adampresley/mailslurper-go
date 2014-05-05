// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/adampresley/mailslurper/data"
	"github.com/adampresley/mailslurper/settings"
)

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
	mailItems := data.Storage.GetMails()
	json, _ := json.Marshal(mailItems)
	settings.Config.WriteJson(writer, json)
}
