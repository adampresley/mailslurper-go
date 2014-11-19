// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailController

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/adampresley/GoHttpService"
	"github.com/adampresley/mailslurper/libmailslurper/storage"

	"github.com/gorilla/mux"
)

func DownloadAttachment_v1(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	/*
	 * Validate incoming arguments
	 */
	mailId, ok := vars["mailId"]
	if !ok {
		GoHttpService.BadRequest(writer, "A valid mail ID is required")
	}

	attachmentId, ok := vars["attachmentId"]
	if !ok {
		GoHttpService.BadRequest(writer, "A valid attachment ID is required")
		return
	}

	/*
	 * Retrieve the attachment
	 */
	attachment, err := storage.GetAttachment(mailId, attachmentId);
	if err != nil {
		GoHttpService.Error(writer, err.Error())
		return
	}

	/*
	 * Decode the base64 data and stream it back
	 */
	data, err := base64.StdEncoding.DecodeString(attachment.Contents)
	if err != nil {
		GoHttpService.Error(writer, "Cannot decode attachment")
		return
	}

	reader := bytes.NewReader(data)
	http.ServeContent(writer, request, attachment.Headers.FileName, time.Now(), reader)
}
