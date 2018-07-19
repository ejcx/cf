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
	if serviceKey, ok := os.LookupEnv("CF_USER_SERVICE_KEY"); ok {
		api.APIUserServiceKey = serviceKey
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
		resp, err = api.DNSRecords(ZoneId, rec)
	case "EditDNSRecord":
		rec := cloudflare.DNSRecord{
			Type:    Type,
			Name:    Name,
			Content: Content,
			Proxied: Proxied,
			TTL:     Ttl,
		}
		err = api.UpdateDNSRecord(ZoneId, RecordId, rec)
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
		resp, err = api.CreateDNSRecord(ZoneId, rec)
	case "DeleteDNSRecord":
		err = api.DeleteDNSRecord(ZoneId, RecordId)
		if err == nil {
			resp = map[string]interface{}{
				"Success": true,
			}
		}
	case "DeleteZone":
		resp, err = api.DeleteZone(ZoneId)
	case "DeleteCustomHostname":
		err = api.DeleteCustomHostname(ZoneId, CustomHostnameId)
	case "DNSRecord":
		resp, err = api.DNSRecord(ZoneId, RecordId)
	case "ListAllRateLimits":
		resp, err = api.ListAllRateLimits(ZoneId)
	case "ListLoadBalancers":
		resp, err = api.ListLoadBalancers(ZoneId)
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
		resp, err = api.ListOrganizationAccessRules(OrganizationId, ar, Page)
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
		resp, err = api.ListPageRules(ZoneId)
	case "DeletePageRule":
		err = api.DeletePageRule(ZoneId, PageruleId)
	case "DeleteRailgun":
		err = api.DeleteRailgun(RailgunId)
	case "DisableRailgun":
		resp, err = api.DisableRailgun(RailgunId)
	case "EnableRailgun":
		resp, err = api.EnableRailgun(RailgunId)
	case "DeleteRateLimit":
		err = api.DeleteRateLimit(ZoneId, RatelimitId)
	case "DeleteSSL":
		err = api.DeleteSSL(ZoneId, CertificateId)
	case "ListCustomCerts":
		resp, err = api.ListSSL(ZoneId)
	case "SSLDetails":
		resp, err = api.SSLDetails(ZoneId, CertificateId)
	case "ListWAFPackages":
		resp, err = api.ListWAFPackages(ZoneId)
	case "ListVirtualDns":
		resp, err = api.ListVirtualDNS()
	case "ZoneSettings":
		resp, err = api.ZoneSettings(ZoneId)
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
		api.CreateLoadBalancer(ZoneId, l)
	case "CreateLoadBalancerMonitor":
		l := cloudflare.LoadBalancerMonitor{ExpectedCodes: ExpectedCodes}
		if Method != "" {
			l.Method = Method
		}
		if Timeout > 0 {
			l.Timeout = Timeout
		}
		if Header != "" {
			h := make(map[string][]string)
			err = json.Unmarshal([]byte(Header), &h)
			if err != nil {
				return nil, err
			}
			l.Header = h
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
		resp, err = api.ListWAFRules(ZoneId, PackageId)
	case "ListRailguns":
		resp, err = api.ListRailguns(cloudflare.RailgunListOptions{})
	case "RailgunDetails":
		resp, err = api.RailgunDetails(RailgunId)
	case "RailgunZones":
		resp, err = api.RailgunZones(RailgunId)
	case "AvailableZoneRatePlans":
		resp, err = api.AvailableZoneRatePlans(ZoneId)
	case "ActivationCheck":
		resp, err = api.ZoneActivationCheck(ZoneId)
	case "ZoneDetails":
		resp, err = api.ZoneDetails(ZoneId)
	case "PurgeEverything":
		resp, err = api.PurgeEverything(ZoneId)
	case "ConnectZoneRailgun":
		resp, err = api.ConnectZoneRailgun(ZoneId, RailgunId)
	case "GetIDByName":
		resp, err = api.ZoneIDByName(ZoneName)
	case "ListZoneRailguns":
		resp, err = api.ZoneRailguns(ZoneId)
	case "ZoneSSLSettings":
		resp, err = api.ZoneSSLSettings(ZoneId)
	case "UserDetails":
		resp, err = api.UserDetails()
	case "UserBillingProfile":
		resp, err = api.UserBillingProfile()
	case "VirtualDNS":
		resp, err = api.VirtualDNS(VirtualDnsId)
	case "DeleteVirtualDNS":
		err = api.DeleteVirtualDNS(VirtualDnsId)
	case "PageRule":
		resp, err = api.PageRule(ZoneId, PageruleId)
	case "LoadBalancerDetails":
		resp, err = api.LoadBalancerDetails(ZoneId, LoadbalancerId)
	case "LoadBalancerMonitorDetails":
		resp, err = api.LoadBalancerMonitorDetails(MonitorId)
	case "LoadBalancerPoolDetails":
		resp, err = api.LoadBalancerPoolDetails(PoolId)
	case "OrganizationDetails":
		resp, err = api.OrganizationDetails(OrganizationId)
	case "OrganizationInvites":
		resp, _, err = api.OrganizationInvites(OrganizationId)
	case "OrganizationMembers":
		resp, _, err = api.OrganizationMembers(OrganizationId)
	case "OrganizationRoles":
		resp, _, err = api.OrganizationRoles(OrganizationId)
	case "OriginCertificates":
		resp, err = api.OriginCertificates(cloudflare.OriginCACertificateListOptions{ZoneID: ZoneId})
	case "OriginCertificate":
		resp, err = api.OriginCertificate(CertificateId)
	case "RateLimit":
		resp, err = api.RateLimit(ZoneId, RatelimitId)
	case "ZoneRailgunDetails":
		resp, err = api.ZoneRailgunDetails(ZoneId, RailgunId)
	case "RevokeOriginCertificate":
		resp, err = api.RevokeOriginCertificate(CertificateId)
	case "TestRailgunConnection":
		resp, err = api.TestRailgunConnection(ZoneId, RailgunId)
	case "ZoneSetPaused":
		resp, err = api.ZoneSetPaused(ZoneId, Paused)
	case "CreateCustomHostname":
		resp, err = api.CreateCustomHostname(ZoneId, cloudflare.CustomHostname{
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
		resp, err = api.EditZone(ZoneId, z)
	case "EditZoneVanityNS":
		vns := strings.Split(VanityNS, ",")
		z := cloudflare.ZoneOptions{
			VanityNS: vns,
		}
		resp, err = api.EditZone(ZoneId, z)
	case "ListZoneLockdowns":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListZoneLockdowns(ZoneId, page)
	case "ListUserAgentRules":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListUserAgentRules(ZoneId, page)
	case "CreateZone":
		org := cloudflare.Organization{}
		if OrganizationId != "" {
			org.ID = OrganizationId
		}
		resp, err = api.CreateZone(Name, false, org)
	default:
		break
	}
	return resp, err
}
