package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

type Credentials struct {
	Email string
}

var RootCmd = &cobra.Command{
	Use:   "cf",
	Short: "A CLI for interacting with Cloudflare's V4 API",
}

func loadConfiguration() {

}

func Execute() {
	var (
		cfSecret string
		cfEmail  string
		envSet   bool
	)
	// If we have environment variables set then we should not attempt to
	// load the $HOME/.cf/credentials file.
	if secret, ok := os.LookupEnv("CF_API_KEY"); ok {
		cfSecret = secret
		envSet = true
	}
	if email, ok := ok.LookupEnv("CF_API_EMAIL"); ok {
		cfEmail = email
		envSet = true
	}
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
