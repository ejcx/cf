package lib

import (
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func getCredDir() (string, error) {
	var (
		gopath string
	)
	if p, ok := os.LookupEnv("GOPATH"); ok {
		gopath = p
	} else {
		return "", errors.New("Could not find GOPATH. GOPATH must be set to test.")
	}
	return gopath + "/src/github.com/ejcx/cf/lib", nil
}

func TestGetHomeDir(t *testing.T) {
	if _, err := getHomeDir(); err != nil {
		t.Fatalf("Could not get home dir: %s", err)
	}
}

func TestLoadCredsFile(t *testing.T) {
	testPath, err := getCredDir()
	if err != nil {
		t.Fatalf("Could not load cred dir: %s", err)
	}
	credentials, err := readConfigFile(testPath)
	if err != nil {
		t.Fatalf("Could not read file %s: %s", testPath, err)
	}
	if credentials.Email == "" {
		t.Fatal("Unexpected test credential file result")
	}
}

func TestLoadToEnv(t *testing.T) {
	var (
		foundKey bool
	)
	testPath, err := getCredDir()
	if err != nil {
		log.Fatalf("Could not load cred dir: %s", err)
	}
	c := &CredProvider{
		HomeDir: testPath,
	}
	err = c.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("Error while configuring env: %s", err)
	}
	env := os.Environ()
	for _, kv := range env {
		envParts := strings.Split(kv, "=")
		if len(envParts) != 2 {
			t.Fatalf("Unexpected result from env %s", kv)
		}
		if envParts[0] == "CF_API_KEY" {
			foundKey = true
			if envParts[1] != "AKIAXXX" {
				t.Fatalf("Unexpected value loaded from config")
			}
		}
	}
	if !foundKey {
		t.Fatal("Never found loaded key")
	}
}
