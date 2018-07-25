package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/99designs/keyring"
	"github.com/segmentio/aws-okta/lib"
)

var (
	DefaultCredentialProvider = &CredProvider{}
)

type CredProvider struct {
	HomeDir string
}

type Credentials struct {
	Email          string `json:"Email"`
	Key            string `json:"Key"`
	UserServiceKey string `json:"UserServiceKey"`
	Keychain       bool   `json:"Keychain"`
}

func GetKeyring() (keyring.Keyring, error) {
	var allowedBackends []keyring.BackendType
	return keyring.Open(keyring.Config{
		AllowedBackends:          allowedBackends,
		KeychainTrustApplication: true,
		ServiceName:              "cloudflare-credentials",
		LibSecretCollectionName:  "cloudflare",
		FileDir:                  "~/.cf/",
		FilePasswordFunc: func(prompt string) (string, error) {
			return lib.Prompt("\n"+prompt, true)
		},
	})
}

func GetHomeDir() (string, error) {
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

func (c *Credentials) SetEnv() {
	os.Setenv("CF_API_KEY", c.Key)
	os.Setenv("CF_API_EMAIL", c.Email)
	os.Setenv("CF_USER_SERVICE_KEY", c.UserServiceKey)
}

func isEnvSet() bool {
	// If we already have the cloudflare environment variables set that
	// are used by the cloudflare-go library then we should just return.
	_, keyOk := os.LookupEnv("CF_API_KEY")
	_, emailOk := os.LookupEnv("CF_API_EMAIL")
	_, serviceOk := os.LookupEnv("CF_USER_SERVICE_KEY")
	return keyOk || emailOk || serviceOk
}

func (c *CredProvider) ConfigureEnvironment() error {
	// Nothing to do
	if isEnvSet() {
		return nil
	}

	homedir := c.HomeDir

	// Otherwise, we need to read the ~/.cf/credentials file in the users
	// home directory. It would also be nice to store this in the keychain
	if c.HomeDir == "" {
		h, err := GetHomeDir()
		if err != nil {
			return err
		}
		homedir = h
	}

	creds, err := readConfigFile(homedir)
	if err != nil {
		return err
	}

	if creds.Keychain {
		kr, err := GetKeyring()
		if err != nil {
			return err
		}
		keychainCreds, err := kr.Get("cloudflare-creds")
		if err != nil {
			return err
		}
		if err = json.Unmarshal(keychainCreds.Data, &creds); err != nil {
			return err
		}
	}

	creds.SetEnv()
	return nil
}
