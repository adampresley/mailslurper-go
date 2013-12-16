package settings

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
)

var flagWWW = flag.String("www", "", "Path to the web administrator directory.")
var flagWWWPort = flag.Int("wwwport", 0, "Port number to bind to for WWW administrator.")
var flagSmtpAddress = flag.String("smtpaddress", "", "Address to bind the SMTP server to.")
var flagSmtpPort = flag.Int("smtpport", 0, "Port number to bind to for SMTP server.")

type Configuration struct {
	Header      string
	Footer      string
	WWW         string `json:"www"`
	WWWAbs      string
	WWWPort     float64 `json:"wwwPort"`
	SmtpAddress string  `json:"smtpAddress"`
	SmtpPort    float64 `json:"smtpPort"`
}

var Config Configuration

/*
Returns a fully qualified address and port
*/
func (c *Configuration) GetFullListenAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", int(c.WWWPort))
}

/*
Sets values in this Configuration struct with any command line
flag values.
*/
func (c *Configuration) LoadFlags() {
	flag.Parse()

	if *flagWWW != "" {
		c.WWW = *flagWWW
	}

	if *flagWWWPort != 0 {
		c.WWWPort = float64(*flagWWWPort)
	}

	if *flagSmtpAddress != "" {
		c.SmtpAddress = *flagSmtpAddress
	}

	if *flagSmtpPort != 0 {
		c.SmtpPort = float64(*flagSmtpPort)
	}
}

/*
Loads a header view for HTML output
*/
func (c *Configuration) LoadHeader(fileName string) {
	c.Header = c.loadView(fileName)
}

/*
Loads a footer view for HTML output
*/
func (c *Configuration) LoadFooter(fileName string) {
	c.Footer = c.loadView(fileName)
}

/*
Loads the specified configuration JSON file and parses the contents into a
Configuration struct. It then calls loadFlags() to determine if any command
line arguments override the config file.
*/
func (c *Configuration) LoadSettings(configFile string) error {
	/*
	 * Parse config file contents
	 */
	fileContents, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContents, c)
	if err != nil {
		return err
	}

	/*
	 * Check for command line overrides
	 */
	c.LoadFlags()

	return nil
}

/*
Renders a view to the HTTP response stream
*/
func (c *Configuration) RenderView(writer http.ResponseWriter, fileName string) {
	body := fmt.Sprintf("%s%s%s", c.Header, c.loadView(fileName), c.Footer)
	fmt.Fprintf(writer, body)
}

/*
Saves the current settings structure to the config file.
*/
func (c *Configuration) SaveSettings(configFile string) error {
	config := make(map[string]interface{})

	config["www"] = c.WWW
	config["wwwPort"] = c.WWWPort
	config["smtpAddress"] = c.SmtpAddress
	config["smtpPort"] = c.SmtpPort

	json, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFile, json, 0644)
	if err != nil {
		return err
	}

	return nil
}

/*
Takes a byte array of JSON serialized data and writes it to
the HTTP response stream.
*/
func (c *Configuration) WriteJson(writer http.ResponseWriter, jsonData []byte) {
	content := string(jsonData)

	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Add("Content-Length", strconv.Itoa(len(content)))
	fmt.Fprintf(writer, string(jsonData))
}

func (c *Configuration) loadView(viewFileName string) string {
	fullpath, _ := filepath.Abs(c.WWW)
	filename := fmt.Sprintf("%s/%s.html", fullpath, viewFileName)
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Sprintf("There was an error loading a view page: %s", err))
	}

	return string(body)
}
