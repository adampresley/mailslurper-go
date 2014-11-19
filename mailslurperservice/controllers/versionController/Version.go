// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package versionController

import (
	"fmt"
	"net/http"

	"github.com/adampresley/GoHttpService"
)

func GetVersion_v1(writer http.ResponseWriter, request *http.Request) {
	GoHttpService.Success(writer, fmt.Sprintf("MailSlurperService server version 1.0"))
}
