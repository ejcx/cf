package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

var (
	Default = &CredProvider{}
)

type CredProvider struct {
	HomeDir string
}

type Credentials struct {
	Email string `json:"Email"`
	Key   string `json:"Key"`
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func readConfigFile(homedir string) (*Credentials, error) {
	filename := homedir + "/.cf/credentials"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Credentials{}
	err = json.Unmarshal(buf, c)
	return c, err
}

func (c *CredProvider) ConfigureEnvironment() error {
	// If we already have the cloudflare environment variables set that
	// are used by the cloudflare-go library then we should just return.
	_, keyOk := os.LookupEnv("CF_API_KEY")
	_, emailOk := os.LookupEnv("CF_API_EMAIL")
	if keyOk || emailOk {
		return nil
	}

	homedir := c.HomeDir

	// Otherwise, we need to read the ~/.cf/credentials file in the users
	// home directory. It would also be nice to store this in the keychain
	if c.HomeDir == "" {
		h, err := getHomeDir()
		if err != nil {
			return err
		}
		homedir = h
	}
	creds, err := readConfigFile(homedir)
	if err != nil {
		return err
	}
	os.Setenv("CF_API_KEY", creds.Key)
	os.Setenv("CF_EMAIL_KEY", creds.Email)
	return nil
}
