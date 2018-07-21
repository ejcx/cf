package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	case "ChangePageRule":
		var (
			prt []cloudflare.PageRuleTarget
			pra []cloudflare.PageRuleAction
			r   = cloudflare.PageRule{}
		)

		err = json.Unmarshal([]byte(Targets), &prt)
		if err != nil {
			break
		}
		err = json.Unmarshal([]byte(Actions), &pra)
		if err != nil {
			break
		}
		if Priority > 0 {
			r.Priority = Priority
		}
		r.Status = Status
		r.Actions = pra
		r.Targets = prt
		err = api.ChangePageRule(ZoneId, PageruleId, r)
	case "UpdateRateLimit":
		var (
			rlkv []cloudflare.RateLimitKeyValue
			m    cloudflare.RateLimitTrafficMatcher
			a    cloudflare.RateLimitAction
		)
		err = json.Unmarshal([]byte(Match), &m)
		if err != nil {
			break
		}
		err = json.Unmarshal([]byte(Action), &a)
		if err != nil {
			break
		}
		rl := cloudflare.RateLimit{
			ID:        LimitId,
			Disabled:  Enabled,
			Period:    Period,
			Threshold: Threshold,
			Match:     m,
			Action:    a,
		}
		if Bypass != "" {
			err = json.Unmarshal([]byte(Bypass), &rlkv)
			if err != nil {
				break
			}
			rl.Bypass = rlkv
		}
		if Description != "" {
			rl.Description = Description
		}
		resp, err = api.CreateRateLimit(ZoneId, rl)
	case "CreateSSL":
		z := cloudflare.ZoneCustomSSLOptions{
			Certificate: Certificate,
			PrivateKey:  PrivateKey,
		}
		if BundleMethod != "" {
			z.BundleMethod = BundleMethod
		}
		resp, err = api.CreateSSL(ZoneId, z)
	case "UpdateSSL":
		z := cloudflare.ZoneCustomSSLOptions{
			Certificate: Certificate,
			PrivateKey:  PrivateKey,
		}
		if BundleMethod != "" {
			z.BundleMethod = BundleMethod
		}
		resp, err = api.UpdateSSL(ZoneId, CertificateId, z)
	case "CreateRateLimit":
		var (
			rlkv []cloudflare.RateLimitKeyValue
			m    cloudflare.RateLimitTrafficMatcher
			a    cloudflare.RateLimitAction
		)
		err = json.Unmarshal([]byte(Match), &m)
		if err != nil {
			break
		}
		err = json.Unmarshal([]byte(Action), &a)
		if err != nil {
			break
		}
		rl := cloudflare.RateLimit{
			Disabled:  Enabled,
			Period:    Period,
			Threshold: Threshold,
			Match:     m,
			Action:    a,
		}
		if Bypass != "" {
			err = json.Unmarshal([]byte(Bypass), &rlkv)
			if err != nil {
				break
			}
			rl.Bypass = rlkv
		}
		if Description != "" {
			rl.Description = Description
		}
		resp, err = api.CreateRateLimit(ZoneId, rl)
	case "CreateOriginCertificate":
		h := strings.Split(Hostnames, ",")
		c := cloudflare.OriginCACertificate{
			Hostnames:       h,
			RequestValidity: RequestValidity,
			RequestType:     RequestType,
			CSR:             Csr,
		}
		resp, err = api.CreateOriginCertificate(c)
	case "CreateUserAccessRule":
		var arc cloudflare.AccessRuleConfiguration
		err = json.Unmarshal([]byte(Configuration), &arc)
		if err != nil {
			break
		}
		ar := cloudflare.AccessRule{
			Configuration: arc,
			Mode:          Mode,
		}
		if Notes != "" {
			ar.Notes = Notes
		}
		resp, err = api.CreateUserAccessRule(ar)
	case "UpdateUserAccessRule":
		var arc cloudflare.AccessRuleConfiguration
		err = json.Unmarshal([]byte(Configuration), &arc)
		if err != nil {
			break
		}
		ar := cloudflare.AccessRule{
			Configuration: arc,
			Mode:          Mode,
		}
		if Notes != "" {
			ar.Notes = Notes
		}
		resp, err = api.UpdateUserAccessRule(AccessRuleId, ar)
	case "UpdateOrganizationAccessRule":
		var arc cloudflare.AccessRuleConfiguration
		err = json.Unmarshal([]byte(Configuration), &arc)
		if err != nil {
			break
		}
		ar := cloudflare.AccessRule{
			Configuration: arc,
			Mode:          Mode,
		}
		if Notes != "" {
			ar.Notes = Notes
		}
		resp, err = api.UpdateOrganizationAccessRule(OrganizationId, AccessRuleId, ar)
	case "UpdateZoneAccessRule":
		var arc cloudflare.AccessRuleConfiguration
		err = json.Unmarshal([]byte(Configuration), &arc)
		if err != nil {
			break
		}
		ar := cloudflare.AccessRule{
			Configuration: arc,
			Mode:          Mode,
		}
		if Notes != "" {
			ar.Notes = Notes
		}
		resp, err = api.UpdateZoneAccessRule(OrganizationId, AccessRuleId, ar)
	case "CreateOrganizationAccessRule":
		var arc cloudflare.AccessRuleConfiguration
		err = json.Unmarshal([]byte(Configuration), &arc)
		if err != nil {
			break
		}
		ar := cloudflare.AccessRule{
			Configuration: arc,
			Mode:          Mode,
		}
		if Notes != "" {
			ar.Notes = Notes
		}
		resp, err = api.CreateOrganizationAccessRule(OrganizationId, ar)
	case "CreatePageRule":
		var (
			prt []cloudflare.PageRuleTarget
			pra []cloudflare.PageRuleAction
			r   = cloudflare.PageRule{}
		)

		err = json.Unmarshal([]byte(Targets), &prt)
		if err != nil {
			break
		}
		err = json.Unmarshal([]byte(Actions), &pra)
		if err != nil {
			break
		}
		if Priority > 0 {
			r.Priority = Priority
		}
		r.Status = Status
		r.Actions = pra
		r.Targets = prt
		resp, err = api.CreatePageRule(ZoneId, r)
	case "CreateVirtualDNS":
		v := &cloudflare.VirtualDNS{}
		v.Name = Name
		v.OriginIPs = strings.Split(OriginIps, ",")
		if MinimumCacheTtl > 0 {
			v.MinimumCacheTTL = uint(MinimumCacheTtl)
		}
		if MaximumCacheTtl > 0 {
			v.MaximumCacheTTL = uint(MaximumCacheTtl)
		}
		v.DeprecateAnyRequests = DeprecateAnyRequest
		resp, err = api.CreateVirtualDNS(v)
	case "UpdateVirtualDNS":
		v := &cloudflare.VirtualDNS{}
		if Name != "" {
			v.Name = Name
		}
		if OriginIps != "" {
			v.OriginIPs = strings.Split(OriginIps, ",")
		}
		if MinimumCacheTtl > 0 {
			v.MinimumCacheTTL = uint(MinimumCacheTtl)
		}
		if MaximumCacheTtl > 0 {
			v.MaximumCacheTTL = uint(MaximumCacheTtl)
		}
		v.DeprecateAnyRequests = DeprecateAnyRequest
		err = api.UpdateVirtualDNS(VirtualDnsId, *v)
	case "CreateZoneLockdown":
		var c []cloudflare.ZoneLockdownConfig
		err = json.Unmarshal([]byte(Configuration), &c)
		if err != nil {
			break
		}
		urlList := strings.Split(Urls, ",")
		zl := cloudflare.ZoneLockdown{
			URLs:           urlList,
			Paused:         Paused,
			Configurations: c,
		}
		if Description != "" {
			zl.Description = Description
		}
		resp, err = api.CreateZoneLockdown(ZoneId, zl)
	case "UpdateZoneLockdown":
		var c []cloudflare.ZoneLockdownConfig
		err = json.Unmarshal([]byte(Configuration), &c)
		if err != nil {
			break
		}
		urlList := strings.Split(Urls, ",")
		zl := cloudflare.ZoneLockdown{
			URLs:           urlList,
			Paused:         Paused,
			Configurations: c,
		}
		if Description != "" {
			zl.Description = Description
		}
		resp, err = api.UpdateZoneLockdown(ZoneId, LockdownId, zl)
	case "UpdateZoneSettings":
		var zs []cloudflare.ZoneSetting
		err = json.Unmarshal([]byte(ZoneSettings), &zs)
		if err != nil {
			break
		}
		resp, err = api.UpdateZoneSettings(ZoneId, zs)
	case "ModifyLoadBalancerPool":
		var lbo []cloudflare.LoadBalancerOrigin
		err = json.Unmarshal([]byte(Origins), &lbo)
		if err != nil {
			break
		}
		l := cloudflare.LoadBalancerPool{
			ID:      PoolId,
			Name:    Name,
			Origins: lbo,
		}
		if Description != "" {
			l.Description = Description
		}
		if Disabled {
			l.Enabled = false
		}
		if MinimumOrigins > 0 {
			l.MinimumOrigins = MinimumOrigins
		}
		if Monitor != "" {
			l.Monitor = Monitor
		}
		if NotificationEmail != "" {
			l.NotificationEmail = NotificationEmail
		}
		resp, err = api.ModifyLoadBalancerPool(l)
	case "CreateLoadBalancerPool":
		var lbo []cloudflare.LoadBalancerOrigin
		err = json.Unmarshal([]byte(Origins), &lbo)
		if err != nil {
			break
		}
		l := cloudflare.LoadBalancerPool{
			Name:    Name,
			Origins: lbo,
		}
		if Description != "" {
			l.Description = Description
		}
		if Disabled {
			l.Enabled = false
		}
		if MinimumOrigins > 0 {
			l.MinimumOrigins = MinimumOrigins
		}
		if Monitor != "" {
			l.Monitor = Monitor
		}
		if NotificationEmail != "" {
			l.NotificationEmail = NotificationEmail
		}
		resp, err = api.CreateLoadBalancerPool(l)
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
	case "ReprioritizeSSL":
		var p []cloudflare.ZoneCustomSSLPriority
		err = json.Unmarshal([]byte(PriorityList), &p)
		if err != nil {
			break
		}
		resp, err = api.ReprioritizeSSL(ZoneId, p)
	case "CustomHostname":
		resp, err = api.CustomHostname(ZoneId, CustomHostnameId)
	case "CustomHostnames":
		resp, _, err = api.CustomHostnames(ZoneId, Page, cloudflare.CustomHostname{})
	case "CustomHostnameIDByName":
		resp, err = api.CustomHostnameIDByName(ZoneId, Name)
	case "DeleteDNSRecord":
		err = api.DeleteDNSRecord(ZoneId, RecordId)
		if err == nil {
			resp = map[string]interface{}{
				"Success": true,
			}
		}
	case "DeleteZone":
		resp, err = api.DeleteZone(ZoneId)
	case "DeleteLoadBalancer":
		err = api.DeleteLoadBalancer(ZoneId, LoadbalancerId)
	case "DeleteUserAgentRule":
		resp, err = api.DeleteUserAgentRule(ZoneId, UserAgentId)
	case "CreateUserAgentRule":
		var c cloudflare.UserAgentRuleConfig
		u := cloudflare.UserAgentRule{
			Mode:   Mode,
			Paused: Paused,
		}
		err = json.Unmarshal([]byte(Configuration), &c)
		if err != nil {
			break
		}
		u.Configuration = c
		if Description != "" {
			u.Description = Description
		}
		resp, err = api.CreateUserAgentRule(ZoneId, u)
	case "UpdateUserAgentRule":
		var c cloudflare.UserAgentRuleConfig
		u := cloudflare.UserAgentRule{
			Mode:   Mode,
			Paused: Paused,
		}
		err = json.Unmarshal([]byte(Configuration), &c)
		if err != nil {
			break
		}
		u.Configuration = c
		if Description != "" {
			u.Description = Description
		}
		resp, err = api.UpdateUserAgentRule(ZoneId, UserAgentId, u)
	case "DeleteUserAccessRule":
		resp, err = api.DeleteUserAccessRule(AccessRuleId)
	case "CreateRailgun":
		resp, err = api.CreateRailgun(Name)
	case "DeleteLoadBalancerMonitor":
		err = api.DeleteLoadBalancerMonitor(MonitorId)
	case "DeleteLoadBalancerPool":
		err = api.DeleteLoadBalancerPool(PoolId)
	case "DeleteCustomHostname":
		err = api.DeleteCustomHostname(ZoneId, CustomHostnameId)
	case "DeleteZoneLockdown":
		resp, err = api.DeleteZoneLockdown(ZoneId, LockdownId)
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
	case "ListZoneAccessRules":
		ar := cloudflare.AccessRule{}
		if Notes != "" {
			ar.Notes = Notes
		}
		if Mode != "" {
			ar.Mode = Mode
		}
		resp, err = api.ListOrganizationAccessRules(ZoneId, ar, Page)
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
	case "DeleteZoneAccessRule":
		resp, err = api.DeleteZoneAccessRule(ZoneId, AccessRuleId)
	case "DisableRailgun":
		resp, err = api.DisableRailgun(RailgunId)
	case "EnableRailgun":
		resp, err = api.EnableRailgun(RailgunId)
	case "DeleteRateLimit":
		err = api.DeleteRateLimit(ZoneId, RatelimitId)
	case "DeleteSSL":
		err = api.DeleteSSL(ZoneId, CertificateId)
	case "DeleteOrganizationAccessRule":
		resp, err = api.DeleteOrganizationAccessRule(OrganizationId, AccessRuleId)
	case "UpdateUser":
		u := &cloudflare.User{}
		var setVar bool
		if FirstName != "" {
			u.FirstName = FirstName
			setVar = true
		}
		if LastName != "" {
			u.LastName = LastName
			setVar = true
		}
		if Telephone != "" {
			u.Telephone = Telephone
			setVar = true
		}
		if Country != "" {
			u.Country = Country
			setVar = true
		}
		if Zipcode != "" {
			u.Zipcode = Zipcode
			setVar = true
		}
		if setVar {
			resp, err = api.UpdateUser(u)
		} else {
			resp, err = api.UserDetails()
		}
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
	case "ModifyLoadBalancer":
		d := strings.Split(DefaultPools, ",")
		l := cloudflare.LoadBalancer{
			ID:           LoadbalancerId,
			Name:         Name,
			FallbackPool: FallbackPool,
			DefaultPools: d,
			Proxied:      Proxied,
		}
		if Ttl > 0 {
			l.TTL = Ttl
		}
		resp, err = api.ModifyLoadBalancer(ZoneId, l)
	case "ModifyLoadBalancerMonitor":
		l := cloudflare.LoadBalancerMonitor{
			ID:            MonitorId,
			ExpectedCodes: ExpectedCodes,
		}
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
		resp, err = api.ModifyLoadBalancerMonitor(l)
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
	case "Purge":
		p := cloudflare.PurgeCacheRequest{}
		if Files != "" {
			p.Files = strings.Split(Files, ",")
		}
		if Hosts != "" {
			p.Hosts = strings.Split(Hosts, ",")
		}
		if Tags != "" {
			p.Tags = strings.Split(Tags, ",")
		}
		resp, err = api.PurgeCache(ZoneId, p)
	case "ConnectZoneRailgun":
		resp, err = api.ConnectZoneRailgun(ZoneId, RailgunId)
	case "DisconnectZoneRailgun":
		resp, err = api.DisconnectZoneRailgun(ZoneId, RailgunId)
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
	case "UpdateCustomHostnameSSL":
		h := cloudflare.CustomHostnameSSL{
			Method: Method,
			Type:   Type,
		}
		resp, err = api.UpdateCustomHostnameSSL(ZoneId, CustomHostnameId, h)
	case "EditZonePaused":
		z := cloudflare.ZoneOptions{
			Paused: &Paused,
		}
		resp, err = api.EditZone(ZoneId, z)
	case "EditZoneVanityNS":
		vns := strings.Split(VanityNs, ",")
		z := cloudflare.ZoneOptions{
			VanityNS: vns,
		}
		resp, err = api.EditZone(ZoneId, z)
	case "ZoneSetVanityNS":
		vns := strings.Split(VanityNs, ",")
		resp, err = api.ZoneSetVanityNS(ZoneId, vns)
	case "ListZoneLockdowns":
		page := 1
		if Page != 0 {
			page = Page
		}
		resp, err = api.ListZoneLockdowns(ZoneId, page)
	case "ZoneLockdown":
		resp, err = api.ZoneLockdown(ZoneId, LockdownId)
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
		resp, err = api.CreateZone(Name, Jumpstart, org)
	case "ZoneAnalyticsDashboard":
		z := cloudflare.ZoneAnalyticsOptions{}
		if Since != "" {
			t, err := time.Parse(time.RFC3339, Since)
			if err != nil {
				log.Fatalf("Invalid timestamp passed to Since: %s", err)
			}
			z.Since = &t
		}
		if Until != "" {
			t, err := time.Parse(time.RFC3339, Until)
			if err != nil {
				log.Fatalf("Invalid timestamp passed to Until: %s", err)
			}
			z.Until = &t
		}
		z.Continuous = &Continuous
		resp, err = api.ZoneAnalyticsDashboard(ZoneId, z)
	case "ZoneAnalyticsByColocation":
		z := cloudflare.ZoneAnalyticsOptions{}
		if Since != "" {
			t, err := time.Parse(time.RFC3339, Since)
			if err != nil {
				log.Fatalf("Invalid timestamp passed to Since: %s", err)
			}
			z.Since = &t
		}
		if Until != "" {
			t, err := time.Parse(time.RFC3339, Until)
			if err != nil {
				log.Fatalf("Invalid timestamp passed to Until: %s", err)
			}
			z.Until = &t
		}
		z.Continuous = &Continuous
		resp, err = api.ZoneAnalyticsByColocation(ZoneId, z)
	default:
		break
	}
	return resp, err
}
