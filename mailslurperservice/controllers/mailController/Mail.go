// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailController

import (
	"net/http"
	"strconv"

	"github.com/adampresley/GoHttpService"
	"github.com/adampresley/mailslurper/libmailslurper/storage"

	"github.com/gorilla/mux"
)

func GetMail_v1(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	/*
	 * Validate incoming arguments
	 */
	id, ok := vars["mailId"];
	if !ok {
		GoHttpService.BadRequest(writer, "A valid mail ID is required")
		return
	}

	/*
	 * Retrieve the mail item
	 */
	mailItem, err := storage.GetMail(id);
	if err != nil {
		GoHttpService.Error(writer, err.Error())
		return
	}

	result := make(map[string]interface{})
	result["mailItem"] = mailItem

	GoHttpService.WriteJson(writer, result, 200)
}

func GetMailCollection_v1(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	/*
	 * Validate incoming arguments. A page is currently 100 items, hard coded
	 */
	pageNumberString, ok := vars["pageNumber"]
	if !ok {
		GoHttpService.BadRequest(writer, "A valid page number is required")
		return
	}

	pageNumber, _ := strconv.Atoi(pageNumberString)
	length := 100
	offset := (pageNumber - 1) * length

	/*
	 * Retrieve mail items
	 */
	mailItems, err := storage.GetMailCollection(offset, length)
	if err != nil {
		GoHttpService.Error(writer, err.Error())
		return
	}

	result := make(map[string]interface{})
	result["mailItems"] = mailItems

	GoHttpService.WriteJson(writer, result, 200)
}
