// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adampresley/mailslurper/data"
)

/*
This function handles a web GET request for "/mails". It queries the storage
engine for all mail items, sets the content type header to text/json, and
returns a JSON-serialized array of mail data.
*/
func GetMailCollection(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	mailItems := data.Storage.GetMails()

	json, _ := json.Marshal(mailItems)
	fmt.Fprintf(writer, string(json))
}
