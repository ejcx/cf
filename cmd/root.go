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

func Main(cmd *cobra.Command, args []string, name string) {
	err := lib.DefaultCredentialProvider.ConfigureEnvironment()
	if err != nil {
		log.Fatalf("No set of credentials to use: %s", err)
	}

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal("Could not initialize api object: %s", err)
	}

	r, err := root(cmd, args, name, api)
	if err != nil {
		log.Fatal("Could not make cloudflare request: %s", err)
	}
	buf, err := json.MarshalIndent(r, " ", "    ")
	if err != nil {
		log.Fatal("Could not make print resp: %s", err)
	}
	fmt.Println(string(buf))
}

func root(cmd *cobra.Command, args []string, name string, api *cloudflare.API) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	switch name {
	case "ListZones":
		if ZoneNameFilter != "" {
			resp, err = api.ListZones(ZoneNameFilter)
		} else {
			resp, err = api.ListZones()
		}
	case "DNSRecords":
		rec := cloudflare.DNSRecord{}
		if Type != "" {
			rec.Type = Type
		}
		if Name != "" {
			rec.Name = Name
		}
		if Content != "" {
			rec.Content = Content
		}
		resp, err = api.DNSRecords(ZoneID, rec)
	case "CreateDNSRecord":
		rec := cloudflare.DNSRecord{}
		if Type != "" {
			rec.Type = Type
		}
		if Name != "" {
			rec.Name = Name
		}
		if Content != "" {
			rec.Content = Content
		}
		if Ttl != 0 {
			rec.TTL = Ttl
		}
		resp, err = api.CreateDNSRecord(ZoneID, rec)
	case "DeleteDNSRecord":
		err = api.DeleteDNSRecord(ZoneID, RecordID)
		if err == nil {
			resp = map[string]interface{}{
				"Success": true,
			}
		}
	case "DeleteZone":
		resp, err = api.DeleteZone(ZoneID)
	default:
		break
	}
	return resp, err
}
