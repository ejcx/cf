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

// Main is where the high level program execution takes place.
// The first thing that happens is the credentials are attempted
// to be loaded from a file or your environment. If the creds are
// loaded from a file then they are set as env vars to be used
// by cloudflare-go.
// Next, `root` is called, which is where API calls are made to
// the cloudflare v4 API by using the `cloudflare-go` library.
// Finally, the results are outputted to stdout (or stderr if
// something disasterous happens).
func Main(cmd *cobra.Command, args []string, name string) {
	err := lib.DefaultCredentialProvider.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("No set of credentials to use: %s", err)
	}

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatalf("Could not initialize api object: %s", err)
	}
	if serviceKey, ok := os.LookupEnv("CF_USER_SERVICE_KEY"); ok {
		api.APIUserServiceKey = serviceKey
	}

	r, err := Run(cmd, args, name, api)
	if err != nil {
		log.Fatalf("Could not make cloudflare request: %s", err)
	}
	buf, err := json.MarshalIndent(r, " ", "    ")
	if err != nil {
		log.Fatalf("Could not make print resp: %s", err)
	}
	if string(buf) != "null" {
		fmt.Println(string(buf))
	}
}
