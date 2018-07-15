package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/ejcx/cf/lib"
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

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func Main(cmd *cobra.Command, args []string) {
	err := lib.DefaultCredentialProvider.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("No set of credentials to use: %s", err)
	}

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal("Could not initialize api object: %s", err)
	}

	r, err := root(cmd, args, api)
	if err != nil {
		log.Fatal("Could not make cloudflare request: %s", err)
	}
	buf, err := json.MarshalIndent(r, " ", "    ")
	if err != nil {
		log.Fatal("Could not make print resp: %s", err)
	}
	fmt.Println(string(buf))
}
