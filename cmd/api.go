package cmd

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func ListZones(api *cloudflare.API, ZoneNameFilter string) (resp interface{}, err error) {
	if ZoneNameFilter != "" {
		resp, err = api.ListZones(ZoneNameFilter)
	} else {
		resp, err = api.ListZones()
	}
	return
}

func ListDnsRecords(api *cloudflare.API, ZoneId string, Type string, Name string, Content string) (resp interface{}, err error) {
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
	return
}

func CreateDnsRecord(api *cloudflare.API, ZoneId string, Type string, Name string, Content string, Ttl int, NotProxied bool, Priority int) (resp interface{}, err error) {
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
	return
}

func DeleteDnsRecord(api *cloudflare.API, ZoneId string, RecordId string) (resp interface{}, err error) {
	err = api.DeleteDNSRecord(ZoneId, RecordId)
	if err == nil {
		resp = map[string]interface{}{
			"Success": true,
		}
	}
	return
}

func DeleteZone(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	resp, err = api.DeleteZone(ZoneId)
	return
}

func CreateZone(api *cloudflare.API, Name string, Jumpstart bool, OrganizationId string) (resp interface{}, err error) {
	org := cloudflare.Organization{}
	if OrganizationId != "" {
		org.ID = OrganizationId
	}
	resp, err = api.CreateZone(Name, Jumpstart, org)
	return
}

func DNSRecord(api *cloudflare.API, ZoneId string, RecordId string) (resp interface{}, err error) {

	return
}

func ListAllRateLimits(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	resp, err = api.ListAllRateLimits(ZoneId)
	return
}

func ListLoadBalancers(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ListOrganizations(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func ListPageRules(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	resp, err = api.ListPageRules(ZoneId)

	return
}

func ListCustomCerts(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ListUserAgentRules(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	return
}

func ListWAFPackages(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ListWAFRules(api *cloudflare.API, ZoneId string, PackageId string) (resp interface{}, err error) {
	return
}

func ListZoneLockdowns(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	return
}

func ZoneLockdown(api *cloudflare.API, ZoneId string, LockdownId string) (resp interface{}, err error) {
	return
}

func EditZonePaused(api *cloudflare.API, ZoneId string, Paused bool) (resp interface{}, err error) {
	return
}

func EditZoneVanityNS(api *cloudflare.API, ZoneId string, VanityNs string) (resp interface{}, err error) {
	return
}

func ZoneSetVanityNS(api *cloudflare.API, ZoneId string, VanityNs string) (resp interface{}, err error) {
	return
}

func EditDNSRecord(api *cloudflare.API, Proxied bool, ZoneId string, RecordId string, Type string, Name string, Content string, Ttl int) (resp interface{}, err error) {
	rec := cloudflare.DNSRecord{
		Type:    Type,
		Name:    Name,
		Content: Content,
		Proxied: Proxied,
		TTL:     Ttl,
	}
	err = api.UpdateDNSRecord(ZoneId, RecordId, rec)
	return
}

func ListLoadBalancerMonitors(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func ListLoadBalancerPools(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func ListOrganizationAccessRules(api *cloudflare.API, OrganizationId string, Notes string, Mode string, Page int) (resp interface{}, err error) {
	return
}

func ListRailguns(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func ListZoneRailguns(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func ListUserAccessRules(api *cloudflare.API, Notes string, Mode string, Page int) (resp interface{}, err error) {
	return
}

func ListVirtualDns(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func AvailableZoneRatePlans(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ConnectZoneRailgun(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}

func CreateCustomHostname(api *cloudflare.API, ZoneId string, Hostname string, Method string, Type string) (resp interface{}, err error) {
	return
}

func CreateLoadBalancerMonitor(api *cloudflare.API, ExpectedCodes string, Method string, Header string, Timeout int, Path string, Interval int, Retries int, ExpectedBody string, Type string, Description string) (resp interface{}, err error) {
	return
}

func CreateLoadBalancer(api *cloudflare.API, ZoneId string, Name string, FallbackPool string, DefaultPools string, Proxied bool, Ttl int) (resp interface{}, err error) {
	return
}

func PurgeEverything(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ActivationCheck(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ZoneDetails(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func GetIDByName(api *cloudflare.API, ZoneName string) (resp interface{}, err error) {
	return
}

func ZoneSSLSettings(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func ZoneSettings(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func UserDetails(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func UserBillingProfile(api *cloudflare.API) (resp interface{}, err error) {
	return
}

func VirtualDNS(api *cloudflare.API, VirtualDnsId string) (resp interface{}, err error) {
	return
}

func DeleteVirtualDNS(api *cloudflare.API, VirtualDnsId string) (resp interface{}, err error) {
	return
}

func PageRule(api *cloudflare.API, ZoneId string, PageruleId string) (resp interface{}, err error) {
	return
}

func LoadBalancerDetails(api *cloudflare.API, ZoneId string, LoadbalancerId string) (resp interface{}, err error) {
	return
}

func LoadBalancerMonitorDetails(api *cloudflare.API, MonitorId string) (resp interface{}, err error) {
	return
}

func LoadBalancerPoolDetails(api *cloudflare.API, PoolId string) (resp interface{}, err error) {
	return
}

func OrganizationDetails(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}

func OrganizationInvites(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}

func OrganizationMembers(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}

func OrganizationRoles(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}

func OriginCertificates(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}

func OriginCertificate(api *cloudflare.API, CertificateId string) (resp interface{}, err error) {
	return
}

func SSLDetails(api *cloudflare.API, CertificateId string, ZoneId string) (resp interface{}, err error) {
	return
}

func RailgunDetails(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}

func RailgunZones(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}

func RateLimit(api *cloudflare.API, ZoneId string, RatelimitId string) (resp interface{}, err error) {
	return
}

func RevokeOriginCertificate(api *cloudflare.API, CertificateId string) (resp interface{}, err error) {
	return
}

func TestRailgunConnection(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}

func ZoneRailgunDetails(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}

func CustomHostname(api *cloudflare.API, ZoneId string, CustomHostnameId string) (resp interface{}, err error) {
	return
}

func CustomHostnameIDByName(api *cloudflare.API, ZoneId string, Name string) (resp interface{}, err error) {
	return
}

func ZoneSetPaused(api *cloudflare.API, ZoneId string, Paused bool) (resp interface{}, err error) {
	return
}

func DeletePageRule(api *cloudflare.API, ZoneId string, PageruleId string) (resp interface{}, err error) {
	err = api.DeletePageRule(ZoneId, PageruleId)

	return
}

func DeleteRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	err = api.DeleteRailgun(RailgunId)

	return
}

func DisableRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	resp, err = api.DisableRailgun(RailgunId)
	return
}

func DisconnectZoneRailgun(api *cloudflare.API, RailgunId string, ZoneId string) (resp interface{}, err error) {
	return
}

func EnableRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}

func DeleteRateLimit(api *cloudflare.API, ZoneId string, RatelimitId string) (resp interface{}, err error) {
	return
}

func DeleteSSL(api *cloudflare.API, ZoneId string, CertificateId string) (resp interface{}, err error) {
	err = api.DeleteSSL(ZoneId, CertificateId)
	return
}

func DeleteCustomHostname(api *cloudflare.API, ZoneId string, CustomHostnameId string) (resp interface{}, err error) {
	return
}

func DeleteLoadBalancer(api *cloudflare.API, ZoneId string, LoadbalancerId string) (resp interface{}, err error) {
	err = api.DeleteLoadBalancer(ZoneId, LoadbalancerId)
	return
}

func DeleteLoadBalancerMonitor(api *cloudflare.API, MonitorId string) (resp interface{}, err error) {
	return
}

func DeleteLoadBalancerPool(api *cloudflare.API, PoolId string) (resp interface{}, err error) {
	return
}

func DeleteOrganizationAccessRule(api *cloudflare.API, OrganizationId string, AccessRuleId string) (resp interface{}, err error) {
	resp, err = api.DeleteOrganizationAccessRule(OrganizationId, AccessRuleId)
	return
}

func CreateRailgun(api *cloudflare.API, Name string) (resp interface{}, err error) {
	resp, err = api.CreateRailgun(Name)
	return
}

func DeleteUserAccessRule(api *cloudflare.API, AccessRuleId string) (resp interface{}, err error) {
	resp, err = api.DeleteUserAccessRule(AccessRuleId)
	return
}

func DeleteUserAgentRule(api *cloudflare.API, UserAgentId string, ZoneId string) (resp interface{}, err error) {
	resp, err = api.DeleteUserAgentRule(ZoneId, UserAgentId)
	return
}

func DeleteZoneAccessRule(api *cloudflare.API, AccessRuleId string, ZoneId string) (resp interface{}, err error) {
	return
}

func DeleteZoneLockdown(api *cloudflare.API, ZoneId string, LockdownId string) (resp interface{}, err error) {
	return
}

func ZoneAnalyticsByColocation(api *cloudflare.API, ZoneId string, Since string, Until string, Continuous bool) (resp interface{}, err error) {
	return
}

func ZoneAnalyticsDashboard(api *cloudflare.API, ZoneId string, Since string, Until string, Continuous bool) (resp interface{}, err error) {
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
	return
}

func UpdateUser(api *cloudflare.API, FirstName string, LastName string, Telephone string, Country string, Zipcode string) (resp interface{}, err error) {
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
	return
}

func CreateLoadBalancerPool(api *cloudflare.API, Name string, Origins string, Description string, Disabled bool, MinimumOrigins int, Monitor string, NotificationEmail string) (resp interface{}, err error) {
	return
}

func ModifyLoadBalancerPool(api *cloudflare.API, PoolId string, Name string, Origins string, Description string, Disabled bool, MinimumOrigins int, Monitor string, NotificationEmail string) (resp interface{}, err error) {
	return
}

func UpdateZoneSettings(api *cloudflare.API, ZoneId string, ZoneSettings string) (resp interface{}, err error) {
	return
}

func UpdateZoneLockdown(api *cloudflare.API, ZoneId string, LockdownId string, Configuration string, Urls string, Paused bool, Description string) (resp interface{}, err error) {
	return
}

func CreateZoneLockdown(api *cloudflare.API, ZoneId string, Configuration string, Urls string, Paused bool, Description string) (resp interface{}, err error) {
	return
}

func CreateVirtualDNS(api *cloudflare.API, Name string, OriginIps string, MinimumCacheTtl int, MaximumCacheTtl int, DeprecateAnyRequest bool) (resp interface{}, err error) {
	return
}

func UpdateVirtualDNS(api *cloudflare.API, VirtualDnsId string, OriginIps string, MinimumCacheTtl int, MaximumCacheTtl int, DeprecateAnyRequest bool) (resp interface{}, err error) {
	return
}

func CreatePageRule(api *cloudflare.API, ZoneId string, Targets string, Actions string, Priority int, Status string) (resp interface{}, err error) {
	return
}

func ChangePageRule(api *cloudflare.API, ZoneId string, PageruleId string, Targets string, Actions string, Priority int, Status string) (resp interface{}, err error) {
	var (
		prt []cloudflare.PageRuleTarget
		pra []cloudflare.PageRuleAction
		r   = cloudflare.PageRule{}
	)

	err = json.Unmarshal([]byte(Targets), &prt)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(Actions), &pra)
	if err != nil {
		return
	}
	if Priority > 0 {
		r.Priority = Priority
	}
	r.Status = Status
	r.Actions = pra
	r.Targets = prt
	err = api.ChangePageRule(ZoneId, PageruleId, r)
	return
}

func CreateOrganizationAccessRule(api *cloudflare.API, OrganizationId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	h := strings.Split(Hostnames, ",")
	c := cloudflare.OriginCACertificate{
		Hostnames:       h,
		RequestValidity: RequestValidity,
		RequestType:     RequestType,
		CSR:             Csr,
	}
	resp, err = api.CreateOriginCertificate(c)
	return
}

func CreateOriginCertificate(api *cloudflare.API, Hostnames string, RequestValidity int, RequestType string, Csr string) (resp interface{}, err error) {
	h := strings.Split(Hostnames, ",")
	c := cloudflare.OriginCACertificate{
		Hostnames:       h,
		RequestValidity: RequestValidity,
		RequestType:     RequestType,
		CSR:             Csr,
	}
	resp, err = api.CreateOriginCertificate(c)
	return
}

func CreateRateLimit(api *cloudflare.API, ZoneId string, Match string, Threshold int, Period int, Action string, Enabled bool, Description string, Bypass string) (resp interface{}, err error) {
	var (
		rlkv []cloudflare.RateLimitKeyValue
		m    cloudflare.RateLimitTrafficMatcher
		a    cloudflare.RateLimitAction
	)
	err = json.Unmarshal([]byte(Match), &m)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(Action), &a)
	if err != nil {
		return
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
			return
		}
		rl.Bypass = rlkv
	}
	if Description != "" {
		rl.Description = Description
	}
	resp, err = api.CreateRateLimit(ZoneId, rl)
	return
}

func UpdateRateLimit(api *cloudflare.API, ZoneId string, LimitId string, Match string, Threshold int, Period int, Action string, Enabled bool, Description string, Bypass string) (resp interface{}, err error) {
	var (
		rlkv []cloudflare.RateLimitKeyValue
		m    cloudflare.RateLimitTrafficMatcher
		a    cloudflare.RateLimitAction
	)
	err = json.Unmarshal([]byte(Match), &m)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(Action), &a)
	if err != nil {
		return
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
			return
		}
		rl.Bypass = rlkv
	}
	if Description != "" {
		rl.Description = Description
	}
	resp, err = api.CreateRateLimit(ZoneId, rl)
	return
}

func CreateUserAccessRule(api *cloudflare.API, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	var arc cloudflare.AccessRuleConfiguration
	err = json.Unmarshal([]byte(Configuration), &arc)
	if err != nil {
		return
	}
	ar := cloudflare.AccessRule{
		Configuration: arc,
		Mode:          Mode,
	}
	if Notes != "" {
		ar.Notes = Notes
	}
	resp, err = api.CreateUserAccessRule(ar)
	return
}

func UpdateUserAccessRule(api *cloudflare.API, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	var arc cloudflare.AccessRuleConfiguration
	err = json.Unmarshal([]byte(Configuration), &arc)
	if err != nil {
		return
	}
	ar := cloudflare.AccessRule{
		Configuration: arc,
		Mode:          Mode,
	}
	if Notes != "" {
		ar.Notes = Notes
	}
	resp, err = api.UpdateUserAccessRule(AccessRuleId, ar)
	return
}

func UpdateZoneAccessRule(api *cloudflare.API, ZoneId string, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	var arc cloudflare.AccessRuleConfiguration
	err = json.Unmarshal([]byte(Configuration), &arc)
	if err != nil {
		return
	}
	ar := cloudflare.AccessRule{
		Configuration: arc,
		Mode:          Mode,
	}
	if Notes != "" {
		ar.Notes = Notes
	}
	resp, err = api.UpdateZoneAccessRule(OrganizationId, AccessRuleId, ar)
	return
}

func UpdateOrganizationAccessRule(api *cloudflare.API, OrganizationId string, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	var arc cloudflare.AccessRuleConfiguration
	err = json.Unmarshal([]byte(Configuration), &arc)
	if err != nil {
		return
	}
	ar := cloudflare.AccessRule{
		Configuration: arc,
		Mode:          Mode,
	}
	if Notes != "" {
		ar.Notes = Notes
	}
	resp, err = api.UpdateOrganizationAccessRule(OrganizationId, AccessRuleId, ar)
	return
}

func ListZoneAccessRules(api *cloudflare.API, ZoneId string, Notes string, Mode string, Page int) (resp interface{}, err error) {
	ar := cloudflare.AccessRule{}
	if Notes != "" {
		ar.Notes = Notes
	}
	if Mode != "" {
		ar.Mode = Mode
	}
	resp, err = api.ListOrganizationAccessRules(ZoneId, ar, Page)
	return
}

func CreateSSL(api *cloudflare.API, ZoneId string, Certificate string, PrivateKey string, BundleMethod string) (resp interface{}, err error) {
	z := cloudflare.ZoneCustomSSLOptions{
		Certificate: Certificate,
		PrivateKey:  PrivateKey,
	}
	if BundleMethod != "" {
		z.BundleMethod = BundleMethod
	}
	resp, err = api.CreateSSL(ZoneId, z)
	return
}

func UpdateSSL(api *cloudflare.API, ZoneId string, CertificateId string, Certificate string, PrivateKey string, BundleMethod string) (resp interface{}, err error) {
	z := cloudflare.ZoneCustomSSLOptions{
		Certificate: Certificate,
		PrivateKey:  PrivateKey,
	}
	if BundleMethod != "" {
		z.BundleMethod = BundleMethod
	}
	resp, err = api.UpdateSSL(ZoneId, CertificateId, z)
	return
}

func Purge(api *cloudflare.API, ZoneId string, Files string, Tags string, Hosts string) (resp interface{}, err error) {
	return
}

func CreateUserAgentRule(api *cloudflare.API, ZoneId string, Mode string, Configuration string, Description string, Paused bool) (resp interface{}, err error) {
	var c cloudflare.UserAgentRuleConfig
	u := cloudflare.UserAgentRule{
		Mode:   Mode,
		Paused: Paused,
	}
	err = json.Unmarshal([]byte(Configuration), &c)
	if err != nil {
		return
	}
	u.Configuration = c
	if Description != "" {
		u.Description = Description
	}
	resp, err = api.CreateUserAgentRule(ZoneId, u)
	return
}

func UpdateUserAgentRule(api *cloudflare.API, ZoneId string, UserAgentId string, Mode string, Configuration string, Description string, Paused bool) (resp interface{}, err error) {
	var c cloudflare.UserAgentRuleConfig
	u := cloudflare.UserAgentRule{
		Mode:   Mode,
		Paused: Paused,
	}
	err = json.Unmarshal([]byte(Configuration), &c)
	if err != nil {
		return
	}
	u.Configuration = c
	if Description != "" {
		u.Description = Description
	}
	resp, err = api.UpdateUserAgentRule(ZoneId, UserAgentId, u)
	return
}

func UpdateCustomHostnameSSL(api *cloudflare.API, ZoneId string, CustomHostnameId string, Method string, Type string) (resp interface{}, err error) {
	return
}

func CustomHostnames(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	resp, err = api.CustomHostname(ZoneId, CustomHostnameId)
	return
}

func ReprioritizeSSL(api *cloudflare.API, ZoneId string, PriorityList string) (resp interface{}, err error) {

	var p []cloudflare.ZoneCustomSSLPriority
	err = json.Unmarshal([]byte(PriorityList), &p)
	if err != nil {
		return
	}
	resp, err = api.ReprioritizeSSL(ZoneId, p)
	return
}

func ModifyLoadBalancerMonitor(api *cloudflare.API, MonitorId string, ExpectedCodes string, Method string, Header string, Timeout int, Path string, Interval int, Retries int, ExpectedBody string, Type string, Description string) (resp interface{}, err error) {
	return
}

func ModifyLoadBalancer(api *cloudflare.API, ZoneId string, LoadbalancerId string, Name string, FallbackPool string, DefaultPools string, Proxied bool, Ttl int) (resp interface{}, err error) {
	return
}
