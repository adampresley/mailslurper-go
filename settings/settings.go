package settings

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var flagWWW = flag.String("www", "", "Path to the web administrator directory.")
var flagWWWPort = flag.Int("wwwport", 0, "Port number to bind to for WWW administrator.")
var flagSmtpAddress = flag.String("smtpaddress", "", "Address to bind the SMTP server to.")
var flagSmtpPort = flag.Int("smtpport", 0, "Port number to bind to for SMTP server.")

type Configuration struct {
	WWW         string  `json:"www"`
	WWWPort     float64 `json:"wwwPort"`
	SmtpAddress string  `json:"smtpAddress"`
	SmtpPort    float64 `json:"smtpPort"`
}

var Config Configuration

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
