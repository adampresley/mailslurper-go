// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package jsonmodel

import (
	"encoding/json"
)

type JsonAttachment struct {
	Id       int    `json:"id"`
	FileName string `json:"fileName"`
}

func NewJsonAttachment(id int, fileName string) JsonAttachment {
	return JsonAttachment{
		Id      : id,
		FileName: fileName,
	}
}

func (this *JsonAttachment) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}