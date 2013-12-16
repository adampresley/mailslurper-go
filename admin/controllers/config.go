// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adampresley/mailslurper/settings"
)

/*
Controller for the configuration page.
*/
func Config(writer http.ResponseWriter, request *http.Request) {
	settings.Config.RenderView(writer, "config")
}

func GetConfig(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})

	response["www"] = settings.Config.WWW
	response["wwwPort"] = settings.Config.WWWPort
	response["smtpAddress"] = settings.Config.SmtpAddress
	response["smtpPort"] = settings.Config.SmtpPort

	json, _ := json.Marshal(response)
	settings.Config.WriteJson(writer, json)
}

/*
Saves configuration values passed in from the administrator form.
This will write new settings to the config.json file. Note
that changes don't take effect until server restart.
*/
func SaveConfig(writer http.ResponseWriter, request *http.Request) {
	wwwPort, err := strconv.ParseFloat(request.FormValue("wwwPort"), 64)
	if err != nil {
		http.Error(writer, fmt.Sprintf("There was an error converting the WWW port to an integer: %s", err), 500)
		return
	}

	smtpPort, err := strconv.ParseFloat(request.FormValue("smtpPort"), 64)
	if err != nil {
		http.Error(writer, fmt.Sprintf("There was an error converting the SMTP port to an integer: %s", err), 500)
		return
	}

	settings.Config.WWW = request.FormValue("www")
	settings.Config.WWWPort = wwwPort
	settings.Config.SmtpAddress = request.FormValue("smtpAddress")
	settings.Config.SmtpPort = smtpPort

	err = settings.Config.SaveSettings("config.json")
	if err != nil {
		http.Error(writer, fmt.Sprintf("There was an error writing your config file: %s", err), 500)
		return
	}

	settings.Config.WriteJson(writer, []byte("{\"success\": true}"))
}
