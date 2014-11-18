package versionController

import (
	"fmt"
	"net/http"

	"github.com/adampresley/GoHttpService"
)

func GetVersion(writer http.ResponseWriter, request *http.Request) {
	GoHttpService.Success(writer, fmt.Sprintf("MailSlurperService server version 1.0"))
}
