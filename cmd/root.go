package cmd

import (
	"fmt"
	"log"
	"os"

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
	err := lib.DefaultCredentialProvider.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("No set of credentials to use: %s", err)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
