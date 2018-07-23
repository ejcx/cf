package cmd

import cloudflare "github.com/cloudflare/cloudflare-go"

func Zone(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Dns(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func User(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Ssl(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Pagerule(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Cache(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Firewall(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Organization(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Ratelimit(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func Loadbalancer(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func ListZones(api *cloudflare.API, ZoneNameFilter string) (resp interface{}, err error) {
	return
}
func ListDnsRecords(api *cloudflare.API, ZoneId string, Type string, Name string, Content string) (resp interface{}, err error) {
	return
}
func CreateDnsRecord(api *cloudflare.API, ZoneId string, Type string, Name string, Content string, Ttl int, NotProxied bool, Priority int) (resp interface{}, err error) {
	return
}
func DeleteDnsRecord(api *cloudflare.API, ZoneId string, RecordId string) (resp interface{}, err error) {
	return
}
func DeleteZone(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func CreateZone(api *cloudflare.API, Name string, Jumpstart bool, OrganizationId string) (resp interface{}, err error) {
	return
}
func ShowDnsRecord(api *cloudflare.API, ZoneId string, RecordId string) (resp interface{}, err error) {
	return
}
func ListRatelimits(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ListLoadbalancers(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ListOrganizations(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func ListPagerules(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ListCustomCerts(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ListUserAgentRules(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	return
}
func ListWafPackages(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ListWafRules(api *cloudflare.API, ZoneId string, PackageId string) (resp interface{}, err error) {
	return
}
func ListZoneLockdowns(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	return
}
func DescribeZoneLockdown(api *cloudflare.API, ZoneId string, LockdownId string) (resp interface{}, err error) {
	return
}
func EditZonePaused(api *cloudflare.API, ZoneId string, Paused bool) (resp interface{}, err error) {
	return
}
func EditZoneVanityNs(api *cloudflare.API, ZoneId string, VanityNs string) (resp interface{}, err error) {
	return
}
func SetVanityNs(api *cloudflare.API, ZoneId string, VanityNs string) (resp interface{}, err error) {
	return
}
func EditDnsRecord(api *cloudflare.API, Proxied bool, ZoneId string, RecordId string, Type string, Name string, Content string, Ttl int) (resp interface{}, err error) {
	return
}
func ListLoadbalancerMonitors(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func ListLoadbalancerPools(api *cloudflare.API) (resp interface{}, err error) {
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
func ListAvailableRatePlans(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ConnectZoneRailgun(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}
func CreateCustomHostname(api *cloudflare.API, ZoneId string, Hostname string, Method string, Type string) (resp interface{}, err error) {
	return
}
func CreateLoadbalancerMonitor(api *cloudflare.API, ExpectedCodes string, Method string, Header string, Timeout int, Path string, Interval int, Retries int, ExpectedBody string, Type string, Description string) (resp interface{}, err error) {
	return
}
func CreateLoadbalancer(api *cloudflare.API, ZoneId string, Name string, FallbackPool string, DefaultPools string, Proxied bool, Ttl int) (resp interface{}, err error) {
	return
}
func PurgeEverything(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func ActivationCheck(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func DescribeZone(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func GetIdByName(api *cloudflare.API, ZoneName string) (resp interface{}, err error) {
	return
}
func ListZoneSslSettings(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func GetZoneSettings(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func Details(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func BillingProfile(api *cloudflare.API) (resp interface{}, err error) {
	return
}
func DescribeVirtualDns(api *cloudflare.API, VirtualDnsId string) (resp interface{}, err error) {
	return
}
func DeleteVirtualDns(api *cloudflare.API, VirtualDnsId string) (resp interface{}, err error) {
	return
}
func DescribePagerule(api *cloudflare.API, ZoneId string, PageruleId string) (resp interface{}, err error) {
	return
}
func DescribeLoadbalancer(api *cloudflare.API, ZoneId string, LoadbalancerId string) (resp interface{}, err error) {
	return
}
func DescribeLoadbalancerMonitor(api *cloudflare.API, MonitorId string) (resp interface{}, err error) {
	return
}
func DescribeLoadbalancerPool(api *cloudflare.API, PoolId string) (resp interface{}, err error) {
	return
}
func DescribeOrganization(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}
func GetOrganizationInvites(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}
func GetOrganizationMembers(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}
func GetOrganizationRoles(api *cloudflare.API, OrganizationId string) (resp interface{}, err error) {
	return
}
func ListOriginCerts(api *cloudflare.API, ZoneId string) (resp interface{}, err error) {
	return
}
func DescribeOriginCert(api *cloudflare.API, CertificateId string) (resp interface{}, err error) {
	return
}
func DescribeZoneOriginCert(api *cloudflare.API, CertificateId string, ZoneId string) (resp interface{}, err error) {
	return
}
func DescribeRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}
func GetRailgunZones(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}
func DescribeRatelimit(api *cloudflare.API, ZoneId string, RatelimitId string) (resp interface{}, err error) {
	return
}
func RevokeOriginCert(api *cloudflare.API, CertificateId string) (resp interface{}, err error) {
	return
}
func TestRailgunConnection(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}
func DescribeZoneRailgun(api *cloudflare.API, ZoneId string, RailgunId string) (resp interface{}, err error) {
	return
}
func DescribeCustomHostname(api *cloudflare.API, ZoneId string, CustomHostnameId string) (resp interface{}, err error) {
	return
}
func DescribeCustomHostnameByName(api *cloudflare.API, ZoneId string, Name string) (resp interface{}, err error) {
	return
}
func SetPaused(api *cloudflare.API, ZoneId string, Paused bool) (resp interface{}, err error) {
	return
}
func DeletePagerule(api *cloudflare.API, ZoneId string, PageruleId string) (resp interface{}, err error) {
	return
}
func DeleteRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}
func DisableRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}
func DisconnectRailgun(api *cloudflare.API, RailgunId string, ZoneId string) (resp interface{}, err error) {
	return
}
func EnableRailgun(api *cloudflare.API, RailgunId string) (resp interface{}, err error) {
	return
}
func DeleteRatelimit(api *cloudflare.API, ZoneId string, RatelimitId string) (resp interface{}, err error) {
	return
}
func DeleteCustomCert(api *cloudflare.API, ZoneId string, CertificateId string) (resp interface{}, err error) {
	return
}
func DeleteCustomHostname(api *cloudflare.API, ZoneId string, CustomHostnameId string) (resp interface{}, err error) {
	return
}
func DeleteLoadbalancer(api *cloudflare.API, ZoneId string, LoadbalancerId string) (resp interface{}, err error) {
	return
}
func DeleteLoadbalancerMonitor(api *cloudflare.API, MonitorId string) (resp interface{}, err error) {
	return
}
func DeleteLoadbalancerPool(api *cloudflare.API, PoolId string) (resp interface{}, err error) {
	return
}
func DeleteOrganizationAccessRule(api *cloudflare.API, OrganizationId string, AccessRuleId string) (resp interface{}, err error) {
	return
}
func CreateRailgun(api *cloudflare.API, Name string) (resp interface{}, err error) {
	return
}
func DeleteUserAccessRule(api *cloudflare.API, AccessRuleId string) (resp interface{}, err error) {
	return
}
func DeleteUserAgentRule(api *cloudflare.API, UserAgentId string, ZoneId string) (resp interface{}, err error) {
	return
}
func DeleteZoneAccessRule(api *cloudflare.API, AccessRuleId string, ZoneId string) (resp interface{}, err error) {
	return
}
func DeleteZoneLockdown(api *cloudflare.API, ZoneId string, LockdownId string) (resp interface{}, err error) {
	return
}
func AnalyticsByColo(api *cloudflare.API, ZoneId string, Since string, Until string, Continuous bool) (resp interface{}, err error) {
	return
}
func AnalyticsDashboard(api *cloudflare.API, ZoneId string, Since string, Until string, Continuous bool) (resp interface{}, err error) {
	return
}
func EditUser(api *cloudflare.API, FirstName string, LastName string, Telephone string, Country string, Zipcode string) (resp interface{}, err error) {
	return
}
func CreateLoadbalancerPool(api *cloudflare.API, Name string, Origins string, Description string, Disabled bool, MinimumOrigins int, Monitor string, NotificationEmail string) (resp interface{}, err error) {
	return
}
func UpdateLoadbalancerPool(api *cloudflare.API, PoolId string, Name string, Origins string, Description string, Disabled bool, MinimumOrigins int, Monitor string, NotificationEmail string) (resp interface{}, err error) {
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
func CreateVirtualDns(api *cloudflare.API, Name string, OriginIps string, MinimumCacheTtl int, MaximumCacheTtl int, DeprecateAnyRequest bool) (resp interface{}, err error) {
	return
}
func UpdateVirtualDns(api *cloudflare.API, VirtualDnsId string, OriginIps string, MinimumCacheTtl int, MaximumCacheTtl int, DeprecateAnyRequest bool) (resp interface{}, err error) {
	return
}
func CreatePagerule(api *cloudflare.API, ZoneId string, Targets string, Actions string, Priority int, Status string) (resp interface{}, err error) {
	return
}
func UpdatePagerule(api *cloudflare.API, ZoneId string, PageruleId string, Targets string, Actions string, Priority int, Status string) (resp interface{}, err error) {
	return
}
func CreateOrganizationAccessRule(api *cloudflare.API, OrganizationId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	return
}
func CreateOriginCert(api *cloudflare.API, Hostnames string, RequestValidity int, RequestType string, Csr string) (resp interface{}, err error) {
	return
}
func CreateRatelimit(api *cloudflare.API, ZoneId string, Match string, Threshold int, Period int, Action string, Enabled bool, Description string, Bypass string) (resp interface{}, err error) {
	return
}
func UpdateRatelimit(api *cloudflare.API, ZoneId string, LimitId string, Match string, Threshold int, Period int, Action string, Enabled bool, Description string, Bypass string) (resp interface{}, err error) {
	return
}
func CreateUserAccessRule(api *cloudflare.API, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	return
}
func UpdateUserAccessRule(api *cloudflare.API, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	return
}
func UpdateZoneAccessRule(api *cloudflare.API, ZoneId string, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	return
}
func UpdateOrganizationAccessRule(api *cloudflare.API, OrganizationId string, AccessRuleId string, Mode string, Configuration string, Notes string) (resp interface{}, err error) {
	return
}
func ListZoneAccessRules(api *cloudflare.API, ZoneId string, Notes string, Mode string, Page int) (resp interface{}, err error) {
	return
}
func UploadCustomCert(api *cloudflare.API, ZoneId string, Certificate string, PrivateKey string, BundleMethod string) (resp interface{}, err error) {
	return
}
func UpdateCustomCert(api *cloudflare.API, ZoneId string, CertificateId string, Certificate string, PrivateKey string, BundleMethod string) (resp interface{}, err error) {
	return
}
func Purge(api *cloudflare.API, ZoneId string, Files string, Tags string, Hosts string) (resp interface{}, err error) {
	return
}
func CreateUserAgentRule(api *cloudflare.API, ZoneId string, Mode string, Configuration string, Description string, Paused bool) (resp interface{}, err error) {
	return
}
func UpdateUserAgentRule(api *cloudflare.API, ZoneId string, UserAgentId string, Mode string, Configuration string, Description string, Paused bool) (resp interface{}, err error) {
	return
}
func UpdateCustomHostname(api *cloudflare.API, ZoneId string, CustomHostnameId string, Method string, Type string) (resp interface{}, err error) {
	return
}
func ListCustomHostnames(api *cloudflare.API, ZoneId string, Page int) (resp interface{}, err error) {
	return
}
func ReprioritizeCerts(api *cloudflare.API, ZoneId string, PriorityList string) (resp interface{}, err error) {
	return
}
func UpdateLoadbalancerMonitor(api *cloudflare.API, MonitorId string, ExpectedCodes string, Method string, Header string, Timeout int, Path string, Interval int, Retries int, ExpectedBody string, Type string, Description string) (resp interface{}, err error) {
	return
}
func UpdateLoadbalancer(api *cloudflare.API, ZoneId string, LoadbalancerId string, Name string, FallbackPool string, DefaultPools string, Proxied bool, Ttl int) (resp interface{}, err error) {
	return
}
