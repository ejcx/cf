package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/99designs/keyring"
	cloudflare "github.com/cloudflare/cloudflare-go"
	cflib "github.com/ejcx/cf/lib"
	"github.com/segmentio/aws-okta/lib"
	"github.com/spf13/cobra"
)

var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A command for configuring your cloudflare api credentials",
	Run: func(cmd *cobra.Command, args []string) {
		err := Configure(cmd, args)
		if err != nil {
			log.Fatalf("Could not configure cf cli: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(ConfigureCmd)
}

func Configure(cmd *cobra.Command, args []string) error {
	email, err := lib.Prompt("Cloudflare Email", false)
	if err != nil {
		return err
	}

	apiKey, err := lib.Prompt("Cloudflare APIKey", true)
	if err != nil {
		return err
	}

	// Add a newline at the beginning and end because sensitive
	// prompts all end up on one line.
	serviceKey, err := lib.Prompt("\nService APIKey", true)
	if err != nil {
		return err
	}
	fmt.Println("")

	// Now that we have the credentials, validate that the apikey and the email
	// are real by calling the User Details API.
	creds := &cflib.Credentials{
		Email:          email,
		Key:            apiKey,
		UserServiceKey: serviceKey,
	}
	creds.SetEnv()

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		return fmt.Errorf("Could not initialize api object: %s", err)
	}

	_, err = api.UserDetails()
	if err != nil {
		return fmt.Errorf("Invalid user credentials: %s", err)
	}

	// Now marshal the data to store in the keychain and set a cloudflare creds
	// file that has keychain set to true and nothing else
	encoded, err := json.Marshal(creds)
	if err != nil {
		return err
	}
	kr, err := cflib.GetKeyring()
	if err != nil {
		return err
	}
	err = kr.Set(keyring.Item{
		Key:   "cloudflare-creds",
		Data:  encoded,
		Label: "cloudflare credentials",
		KeychainNotTrustApplication: false,
	})
	if err != nil {
		return err
	}

	// Write a dumby creds file that points to the keychain.
	buf, err := json.Marshal(cflib.Credentials{
		Keychain: true,
	})
	if err != nil {
		return err
	}
	home, err := cflib.GetHomeDir()
	if err != nil {
		return err
	}
	outfile := home + "/.cf/credentials"
	err = ioutil.WriteFile(outfile, buf, 0600)

	return err
}
