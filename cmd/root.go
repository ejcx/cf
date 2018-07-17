package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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
		log.Fatalf("Could not make cloudflare request: %s", err)
	}
	buf, err := json.MarshalIndent(r, " ", "    ")
	if err != nil {
		log.Fatal("Could not make print resp: %s", err)
	}
	if string(buf) != "null" {
		fmt.Println(string(buf))
	}
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
	case "EditDNSRecord":
		rec := cloudflare.DNSRecord{
			Type:    Type,
			Name:    Name,
			Content: Content,
			Proxied: Proxied,
			TTL:     Ttl,
		}
		err = api.UpdateDNSRecord(ZoneID, RecordID, rec)
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
		if Priority > 0 {
			rec.Priority = Priority
		}
		if Ttl != 0 {
			rec.TTL = Ttl
		}
		rec.Proxied = true
		if NotProxied {
			rec.Proxied = false
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
	case "DNSRecord":
		resp, err = api.DNSRecord(ZoneID, RecordID)
	case "ListAllRateLimits":
		resp, err = api.ListAllRateLimits(ZoneID)
	case "ListLoadBalancers":
		resp, err = api.ListLoadBalancers(ZoneID)
	case "ListLoadBalancerMonitors":
		resp, err = api.ListLoadBalancerMonitors()
	case "ListLoadBalancerPools":
		resp, err = api.ListLoadBalancerPools()
	case "ListOrganizations":
		resp, _, err = api.ListOrganizations()
	case "ListOrganizationAccessRules":
		ar := cloudflare.AccessRule{}
		if Notes != "" {
			ar.Notes = Notes
		}
		if Mode != "" {
			ar.Mode = Mode
		}
		resp, err = api.ListOrganizationAccessRules(OrganizationID, ar, Page)
	case "ListUserAccessRules":
		ar := cloudflare.AccessRule{}
		if Notes != "" {
			ar.Notes = Notes
		}
		if Mode != "" {
			ar.Mode = Mode
		}
		resp, err = api.ListUserAccessRules(ar, Page)
	case "ListPageRules":
		resp, err = api.ListPageRules(ZoneID)
	case "ListCustomCerts":
		resp, err = api.ListSSL(ZoneID)
	case "ListWAFPackages":
		resp, err = api.ListWAFPackages(ZoneID)
	case "ListVirtualDns":
		resp, err = api.ListVirtualDNS()
	case "CreateLoadBalancer":
		d := strings.Split(DefaultPools, ",")
		l := cloudflare.LoadBalancer{
			Name:         Name,
			FallbackPool: FallbackPool,
			DefaultPools: d,
			Proxied:      Proxied,
		}
		if Ttl > 0 {
			l.TTL = Ttl
		}
		api.CreateLoadBalancer(ZoneID, l)
	case "CreateLoadBalancerMonitor":
		l := cloudflare.LoadBalancerMonitor{ExpectedCodes: ExpectedCodes}
		if Method != "" {
			l.Method = Method
		}
		if Timeout > 0 {
			l.Timeout = Timeout
		}
		if Path != "" {
			l.Path = Path
		}
		if Interval > 0 {
			l.Interval = Interval
		}
		if Retries > 0 {
			l.Retries = Retries
		}
		if Type != "" {
			l.Type = Type
		}
		if Description != "" {
			l.Description = Description
		}
		resp, err = api.CreateLoadBalancerMonitor(l)
	case "ListWAFRules":
		resp, err = api.ListWAFRules(ZoneID, PackageID)
	case "ListRailguns":
		resp, err = api.ListRailguns(cloudflare.RailgunListOptions{})
	case "AvailableZoneRatePlans":
		resp, err = api.AvailableZoneRatePlans(ZoneID)
	case "ConnectZoneRailgun":
		resp, err = api.ConnectZoneRailgun(ZoneID, RailgunID)
	case "CreateCustomHostname":
		resp, err = api.CreateCustomHostname(ZoneID, cloudflare.CustomHostname{
			Hostname: Hostname,
			SSL: cloudflare.CustomHostnameSSL{
				Method: Method,
				Type:   Type,
			},
		})
	case "EditZonePaused":
		z := cloudflare.ZoneOptions{
			Paused: &Paused,
		}
		resp, err = api.EditZone(ZoneID, z)
	case "EditZoneVanityNS":
		vns := strings.Split(VanityNS, ",")
		z := cloudflare.ZoneOptions{
			VanityNS: vns,
		}
		resp, err = api.EditZone(ZoneID, z)
	case "ListZoneLockdowns":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListZoneLockdowns(ZoneID, page)
	case "ListUserAgentRules":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListUserAgentRules(ZoneID, page)
	case "CreateZone":
		org := cloudflare.Organization{}
		if OrganizationID != "" {
			org.ID = OrganizationID
		}
		resp, err = api.CreateZone(Name, false, org)
	default:
		break
	}
	return resp, err
}
