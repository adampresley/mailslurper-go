package configuration

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Configuration struct {
	WWWAddress  string  `json:"wwwAddress"`
	WWWPort     int     `json:"wwwPort"`
	SmtpAddress string  `json:"smtpAddress"`
	SmtpPort    int     `json:"smtpPort"`
	DBEngine    string  `json:"dbEngine"`
	DBHost      string  `json:"dbHost"`
	DBPort      string  `json:"dbPort"`
	DBDatabase  string  `json:"dbDatabase"`
	DBUserName  string  `json:"dbUserName"`
	DBPassword  string  `json:"dbPassword"`
}

func (this *Configuration) GetFullSmtpBindingAddress() string {
	return fmt.Sprintf("%s:%d", this.SmtpAddress, this.SmtpPort)
}

func (this *Configuration) GetFullWwwBindingAddress() string {
	return fmt.Sprintf("%s:%d", this.WWWAddress, this.WWWPort)
}

func LoadConfiguration(reader io.Reader) (*Configuration, error) {
	var err error
	var contents bytes.Buffer
	var buffer = make([]byte, 4096)
	var bytesRead int

	result := &Configuration{}
	bufferedReader := bufio.NewReader(reader)

	for {
		bytesRead, err = bufferedReader.Read(buffer)
		if err != nil && err != io.EOF {
			return result, err
		}

		if bytesRead == 0 {
			break
		}

		if _, err := contents.Write(buffer[:bytesRead]); err != nil {
			return result, err
		}
	}

	err = json.Unmarshal(contents.Bytes(), result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (this *Configuration) SaveConfiguration(configFile string) error {
	json, err := json.Marshal(this)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFile, json, 0644)
	if err != nil {
		return err
	}

	return nil
}
