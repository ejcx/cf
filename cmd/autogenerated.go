package cmd

import (
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var (
	ZoneNameFilter      string
	ZoneId              string
	Type                string
	Name                string
	Content             string
	Ttl                 int
	NotProxied          bool
	Priority            int
	RecordId            string
	Jumpstart           bool
	OrganizationId      string
	Page                int
	PackageId           string
	LockdownId          string
	Paused              bool
	VanityNs            string
	Proxied             bool
	Notes               string
	Mode                string
	RailgunId           string
	Hostname            string
	Method              string
	ExpectedCodes       string
	Header              string
	Timeout             int
	Path                string
	Interval            int
	Retries             int
	ExpectedBody        string
	Description         string
	FallbackPool        string
	DefaultPools        string
	ZoneName            string
	VirtualDnsId        string
	PageruleId          string
	LoadbalancerId      string
	MonitorId           string
	PoolId              string
	CertificateId       string
	RatelimitId         string
	CustomHostnameId    string
	AccessRuleId        string
	UserAgentId         string
	Since               string
	Until               string
	Continuous          bool
	FirstName           string
	LastName            string
	Telephone           string
	Country             string
	Zipcode             string
	Origins             string
	Disabled            bool
	MinimumOrigins      int
	Monitor             string
	NotificationEmail   string
	ZoneSettingsObject  string
	Configuration       string
	Urls                string
	OriginIps           string
	MinimumCacheTtl     int
	MaximumCacheTtl     int
	DeprecateAnyRequest bool
	Targets             string
	Actions             string
	Status              string
	Hostnames           string
	RequestValidity     int
	RequestType         string
	Csr                 string
	Match               string
	Threshold           int
	Period              int
	Action              string
	Enabled             bool
	Bypass              string
	LimitId             string
	Certificate         string
	PrivateKey          string
	BundleMethod        string
	Files               string
	Tags                string
	Hosts               string
	PriorityList        string
	Pattern             string
	Disable             bool
	RouteId             string
	Script              string
)

func init() {
	var ListZones = &cobra.Command{
		Use:   "list-zones",
		Short: "Command for listing zones",
		Long:  `  This is a meaty description of the list-zones`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZones")
		},
	}

	ListZones.Flags().StringVar(&ZoneNameFilter, "zone-name-filter", "", "string for filtering by name")

	var ListDnsRecords = &cobra.Command{
		Use:   "list-dns-records",
		Short: "Command for listing dns-records",
		Long:  `  List DNS Records associated with a given zone-id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListDnsRecords")
		},
	}

	ListDnsRecords.Flags().StringVar(&ZoneId, "zone-id", "", "zone id used for filtering")
	ListDnsRecords.MarkFlagRequired("zone-id")

	ListDnsRecords.Flags().StringVar(&Type, "type", "", "DNS Record type used for filter")

	ListDnsRecords.Flags().StringVar(&Name, "name", "", "DNS Record name used for filter")

	ListDnsRecords.Flags().StringVar(&Content, "content", "", "DNS Record content used for filter")

	var CreateDnsRecord = &cobra.Command{
		Use:   "create-dns-record",
		Short: "Command DNS Record",
		Long:  `Create DNS record associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateDnsRecord")
		},
	}

	CreateDnsRecord.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the new dns record")
	CreateDnsRecord.MarkFlagRequired("zone-id")

	CreateDnsRecord.Flags().StringVar(&Type, "type", "", "valid values: A, AAAA, CNAME, TXT, SRV, LOC, MX, NS, SPF, CERT, DNSKEY, DS, NAPTR, SMIMEA, SSHFP, TLSA, URI read only")
	CreateDnsRecord.MarkFlagRequired("type")

	CreateDnsRecord.Flags().StringVar(&Name, "name", "", "DNS Record name (example: example.com), max length: 255")
	CreateDnsRecord.MarkFlagRequired("name")

	CreateDnsRecord.Flags().StringVar(&Content, "content", "", "DNS Record content used for filter")
	CreateDnsRecord.MarkFlagRequired("content")

	CreateDnsRecord.Flags().IntVar(&Ttl, "ttl", 0, "Time to live for DNS record. Value of 1 is 'automatic', min value:120 max value:2147483647")

	CreateDnsRecord.Flags().BoolVar(&NotProxied, "not-proxied", false, "Whether the record is receiving the performance and security benefits of Cloudflare")

	CreateDnsRecord.Flags().IntVar(&Priority, "priority", 0, "Used with some records like MX and SRV to determine priority. If you do not supply a priority for an MX record, a default value of 0 will be set. min value:0 max value:65535.")

	var DeleteDnsRecord = &cobra.Command{
		Use:   "delete-dns-record",
		Short: "Delete DNS Record",
		Long:  `Delete DNS record associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteDnsRecord")
		},
	}

	DeleteDnsRecord.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the record you wish to delete")
	DeleteDnsRecord.MarkFlagRequired("zone-id")

	DeleteDnsRecord.Flags().StringVar(&RecordId, "record-id", "", "The record id associated with the dns record you wish to delete")
	DeleteDnsRecord.MarkFlagRequired("record-id")

	var DeleteZone = &cobra.Command{
		Use:   "delete-zone",
		Short: "Delete zone",
		Long:  `Delete a zone associated with your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteZone")
		},
	}

	DeleteZone.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id that will be deleted")
	DeleteZone.MarkFlagRequired("zone-id")

	var CreateZone = &cobra.Command{
		Use:   "create-zone",
		Short: "Create zone",
		Long:  `Create a zone associated with your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateZone")
		},
	}

	CreateZone.Flags().StringVar(&Name, "name", "", "The zone name that will be added to your account")
	CreateZone.MarkFlagRequired("name")

	CreateZone.Flags().BoolVar(&Jumpstart, "jumpstart", false, "Should the zone DNS be pre-populated")

	CreateZone.Flags().StringVar(&OrganizationId, "organization-id", "", "The organizationID associated with the zone")

	var ShowDnsRecord = &cobra.Command{
		Use:   "show-dns-record",
		Short: "Show DNS Record",
		Long:  `Show a single DNS record associated with a zone ID and record ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DNSRecord")
		},
	}

	ShowDnsRecord.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the DNS Record")
	ShowDnsRecord.MarkFlagRequired("zone-id")

	ShowDnsRecord.Flags().StringVar(&RecordId, "record-id", "", "*Required:* The recordID associated with the DNS Record")
	ShowDnsRecord.MarkFlagRequired("record-id")

	var ListRatelimits = &cobra.Command{
		Use:   "list-ratelimits",
		Short: "Show Ratelimits",
		Long:  `Returns all Rate Limits for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListAllRateLimits")
		},
	}

	ListRatelimits.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the Ratelimits")
	ListRatelimits.MarkFlagRequired("zone-id")

	var ListLoadbalancers = &cobra.Command{
		Use:   "list-loadbalancers",
		Short: "Show LoadBalancers",
		Long:  `Returns all LoadBalancers for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListLoadBalancers")
		},
	}

	ListLoadbalancers.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the Ratelimits")
	ListLoadbalancers.MarkFlagRequired("zone-id")

	var ListOrganizations = &cobra.Command{
		Use:   "list-organizations",
		Short: "Show Organizations",
		Long:  `Returns all Organizations associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListOrganizations")
		},
	}

	var ListPagerules = &cobra.Command{
		Use:   "list-pagerules",
		Short: "Show Page Rules",
		Long:  `Returns all page rules associated with a given zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListPageRules")
		},
	}

	ListPagerules.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the pagerules")
	ListPagerules.MarkFlagRequired("zone-id")

	var ListCustomCerts = &cobra.Command{
		Use:   "list-custom-certs",
		Short: "Show Custom Certs",
		Long:  `Returns all custom certs for a given zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListCustomCerts")
		},
	}

	ListCustomCerts.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the custom certs")
	ListCustomCerts.MarkFlagRequired("zone-id")

	var ListUserAgentRules = &cobra.Command{
		Use:   "list-user-agent-rules",
		Short: "List User-Agent rules",
		Long:  `Returns all User-Agent rules for a specific zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListUserAgentRules")
		},
	}

	ListUserAgentRules.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the user-agent rule.")
	ListUserAgentRules.MarkFlagRequired("zone-id")

	ListUserAgentRules.Flags().IntVar(&Page, "page", 0, "Pagination for user-agent rules")

	var ListWafPackages = &cobra.Command{
		Use:   "list-waf-packages",
		Short: "List WAF Packages",
		Long:  `Return the WAF Packages associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListWAFPackages")
		},
	}

	ListWafPackages.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the WAF packages.")
	ListWafPackages.MarkFlagRequired("zone-id")

	var ListWafRules = &cobra.Command{
		Use:   "list-waf-rules",
		Short: "List WAF Rules",
		Long:  `Return the WAF Rules associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListWAFRules")
		},
	}

	ListWafRules.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the WAF configuration.")
	ListWafRules.MarkFlagRequired("zone-id")

	ListWafRules.Flags().StringVar(&PackageId, "package-id", "", "The package ID associated with the displayed WAF rules.")
	ListWafRules.MarkFlagRequired("package-id")

	var ListZoneLockdowns = &cobra.Command{
		Use:   "list-zone-lockdowns",
		Short: "List Zone Lockdowns",
		Long:  `Return the lockdowns associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZoneLockdowns")
		},
	}

	ListZoneLockdowns.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the WAF configuration.")
	ListZoneLockdowns.MarkFlagRequired("zone-id")

	ListZoneLockdowns.Flags().IntVar(&Page, "page", 0, "Pagination for zone lockdowns.")

	var DescribeZoneLockdown = &cobra.Command{
		Use:   "describe-zone-lockdown",
		Short: "Get detailed zone lockdown information",
		Long:  `Return the detailed information about a lockdown associated with a given zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneLockdown")
		},
	}

	DescribeZoneLockdown.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the zone lockdown.")
	DescribeZoneLockdown.MarkFlagRequired("zone-id")

	DescribeZoneLockdown.Flags().StringVar(&LockdownId, "lockdown-id", "", "The lockdown ID associated with the lockdown")
	DescribeZoneLockdown.MarkFlagRequired("lockdown-id")

	var EditZonePaused = &cobra.Command{
		Use:   "edit-zone-paused",
		Short: "Edit a given zone",
		Long:  `Edit a given zone's properties.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "EditZonePaused")
		},
	}

	EditZonePaused.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the zone being updated")
	EditZonePaused.MarkFlagRequired("zone-id")

	EditZonePaused.Flags().BoolVar(&Paused, "paused", false, "Set to pause the zone while editing the zone")

	var EditZoneVanityNs = &cobra.Command{
		Use:   "edit-zone-vanity-ns",
		Short: "Edit a given zone",
		Long:  `Edit a given zone's properties.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "EditZoneVanityNS")
		},
	}

	EditZoneVanityNs.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the zone being updated")
	EditZoneVanityNs.MarkFlagRequired("zone-id")

	EditZoneVanityNs.Flags().StringVar(&VanityNs, "vanity-ns", "", "Comma delimited list of vanity nameservers")
	EditZoneVanityNs.MarkFlagRequired("vanity-ns")

	var SetVanityNs = &cobra.Command{
		Use:   "set-vanity-ns",
		Short: "Set zone's vanity nameservers",
		Long:  `Set a given zone's vanity nameservers.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneSetVanityNS")
		},
	}

	SetVanityNs.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the zone being updated")
	SetVanityNs.MarkFlagRequired("zone-id")

	SetVanityNs.Flags().StringVar(&VanityNs, "vanity-ns", "", "Comma delimited list of vanity nameservers")
	SetVanityNs.MarkFlagRequired("vanity-ns")

	var EditDnsRecord = &cobra.Command{
		Use:   "edit-dns-record",
		Short: "Edit proxy status for dns record",
		Long:  `Edit an individual dns record's proxied status.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "EditDNSRecord")
		},
	}

	EditDnsRecord.Flags().BoolVar(&Proxied, "proxied", false, "Set this flag is you wish to proxy through Cloudflare, otherwise do not set")

	EditDnsRecord.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the dns record")
	EditDnsRecord.MarkFlagRequired("zone-id")

	EditDnsRecord.Flags().StringVar(&RecordId, "record-id", "", "The record ID that indicates the dns record")
	EditDnsRecord.MarkFlagRequired("record-id")

	EditDnsRecord.Flags().StringVar(&Type, "type", "", "valid values: A, AAAA, CNAME, TXT, SRV, LOC, MX, NS, SPF, CERT, DNSKEY, DS, NAPTR, SMIMEA, SSHFP, TLSA, URI read only")
	EditDnsRecord.MarkFlagRequired("type")

	EditDnsRecord.Flags().StringVar(&Name, "name", "", "DNS Record name (example: example.com), max length: 255")
	EditDnsRecord.MarkFlagRequired("name")

	EditDnsRecord.Flags().StringVar(&Content, "content", "", "DNS Record content used for filter")
	EditDnsRecord.MarkFlagRequired("content")

	EditDnsRecord.Flags().IntVar(&Ttl, "ttl", 0, "Time to live for DNS record. Value of 1 is 'automatic', min value:120 max value:2147483647")

	var ListLoadbalancerMonitors = &cobra.Command{
		Use:   "list-loadbalancer-monitors",
		Short: "Show LoadBalancer Monitors",
		Long:  `Returns all LoadBalancer Monitors. Not specific to a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListLoadBalancerMonitors")
		},
	}

	var ListLoadbalancerPools = &cobra.Command{
		Use:   "list-loadbalancer-pools",
		Short: "Show LoadBalancer Pools",
		Long:  `Returns all LoadBalancer Pools. Not specific to a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListLoadBalancerPools")
		},
	}

	var ListOrganizationAccessRules = &cobra.Command{
		Use:   "list-organization-access-rules",
		Short: "List Organization Access Rules",
		Long:  `Returns all Organizations associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListOrganizationAccessRules")
		},
	}

	ListOrganizationAccessRules.Flags().StringVar(&OrganizationId, "organization-id", "", "The Organization ID")
	ListOrganizationAccessRules.MarkFlagRequired("organization-id")

	ListOrganizationAccessRules.Flags().StringVar(&Notes, "notes", "", "Matching any string within previously created access rules with the notes")

	ListOrganizationAccessRules.Flags().StringVar(&Mode, "mode", "", "valid values: block, challenge, whitelist, js_challenge")

	ListOrganizationAccessRules.Flags().IntVar(&Page, "page", 0, "Requested page within paginated list of results")

	var ListRailguns = &cobra.Command{
		Use:   "list-railguns",
		Short: "List Railguns associated with the account",
		Long:  `Returns all Railguns associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListRailguns")
		},
	}

	var ListZoneRailguns = &cobra.Command{
		Use:   "list-zone-railguns",
		Short: "List all Railguns associated with a zone",
		Long:  `Returns all Railguns associated with a given zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZoneRailguns")
		},
	}

	var ListUserAccessRules = &cobra.Command{
		Use:   "list-user-access-rules",
		Short: "List User Access Rules",
		Long:  `Returns all access rules associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListUserAccessRules")
		},
	}

	ListUserAccessRules.Flags().StringVar(&Notes, "notes", "", "Matching any string within previously created access rules with the notes")

	ListUserAccessRules.Flags().StringVar(&Mode, "mode", "", "valid values: block, challenge, whitelist, js_challenge")

	ListUserAccessRules.Flags().IntVar(&Page, "page", 0, "Requested page within paginated list of results")

	var ListVirtualDns = &cobra.Command{
		Use:   "list-virtual-dns",
		Short: "List Virtual DNS clusters",
		Long:  `Returns all Virtual DNS clusters associated with an account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListVirtualDns")
		},
	}

	var ListAvailableRatePlans = &cobra.Command{
		Use:   "list-available-rate-plans",
		Short: "List all available zone rate plans",
		Long:  `List all rate plans the zone can subscribe to.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "AvailableZoneRatePlans")
		},
	}

	ListAvailableRatePlans.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the rate plans")
	ListAvailableRatePlans.MarkFlagRequired("zone-id")

	var ConnectZoneRailgun = &cobra.Command{
		Use:   "connect-zone-railgun",
		Short: "Connect or disconnect a Railgun",
		Long:  `Connect or disconnect a Railgun`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ConnectZoneRailgun")
		},
	}

	ConnectZoneRailgun.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID to be associated with the railgun")
	ConnectZoneRailgun.MarkFlagRequired("zone-id")

	ConnectZoneRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun ID to be associated with the zone")
	ConnectZoneRailgun.MarkFlagRequired("railgun-id")

	var CreateCustomHostname = &cobra.Command{
		Use:   "create-custom-hostname",
		Short: "Create a custom hostname for an associated zone.",
		Long:  `Add a new custom hostname and request that an SSL certificate be issued for it. One of three validation methods—http, cname, email—should be used, with 'http' recommended if the CNAME is already in place (or will be soon). Specifying 'email' will send an email to the WHOIS contacts on file for the base domain plus hostmaster, postmaster, webmaster, admin, administrator. Specifying 'cname' will return a CNAME that needs to be placed. If http is used and the domain is not already pointing to the Managed CNAME host, the PATCH method must be used once it is (to complete validation).`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateCustomHostname")
		},
	}

	CreateCustomHostname.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the custom hostname")
	CreateCustomHostname.MarkFlagRequired("zone-id")

	CreateCustomHostname.Flags().StringVar(&Hostname, "hostname", "", "The custom hostname that will point to your hostname via CNAME.")
	CreateCustomHostname.MarkFlagRequired("hostname")

	CreateCustomHostname.Flags().StringVar(&Method, "method", "", "The SSL Verification method. valid values: http, email, cname.")
	CreateCustomHostname.MarkFlagRequired("method")

	CreateCustomHostname.Flags().StringVar(&Type, "type", "", "The type of SSL certificate valid values: dv only")
	CreateCustomHostname.MarkFlagRequired("type")

	var CreateLoadbalancerMonitor = &cobra.Command{
		Use:   "create-loadbalancer-monitor",
		Short: "Create a configured monitor",
		Long:  `Create a configured monitor`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateLoadBalancerMonitor")
		},
	}

	CreateLoadbalancerMonitor.Flags().StringVar(&ExpectedCodes, "expected-codes", "", "The expected http response code in the healthcheck")
	CreateLoadbalancerMonitor.MarkFlagRequired("expected-codes")

	CreateLoadbalancerMonitor.Flags().StringVar(&Method, "method", "", "The HTTP method to use for the health check. default value: GET")

	CreateLoadbalancerMonitor.Flags().StringVar(&Header, "header", "", "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Example: {\"Host\": [\"example.com\"],\"X-App-ID\": [\"abc123\"]}")

	CreateLoadbalancerMonitor.Flags().IntVar(&Timeout, "timeout", 0, "The timeout (in seconds) before marking the health check as failed. default value: 5")
	CreateLoadbalancerMonitor.MarkFlagRequired("timeout")

	CreateLoadbalancerMonitor.Flags().StringVar(&Path, "path", "", "The endpoint path to health check against. default value: /")
	CreateLoadbalancerMonitor.MarkFlagRequired("path")

	CreateLoadbalancerMonitor.Flags().IntVar(&Interval, "interval", 0, "The interval between each health check. Shorter intervals may improve failover time, but will increase load. default value 60")
	CreateLoadbalancerMonitor.MarkFlagRequired("interval")

	CreateLoadbalancerMonitor.Flags().IntVar(&Retries, "retries", 0, "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. default value 2")
	CreateLoadbalancerMonitor.MarkFlagRequired("retries")

	CreateLoadbalancerMonitor.Flags().StringVar(&ExpectedBody, "expected-body", "", "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.")

	CreateLoadbalancerMonitor.Flags().StringVar(&Type, "type", "", "The protocol to use for the healthcheck. Currently supported protocols are 'HTTP' and 'HTTPS'. default value: http")
	CreateLoadbalancerMonitor.MarkFlagRequired("type")

	CreateLoadbalancerMonitor.Flags().StringVar(&Description, "description", "", "Object description")
	CreateLoadbalancerMonitor.MarkFlagRequired("description")

	var CreateLoadbalancer = &cobra.Command{
		Use:   "create-loadbalancer",
		Short: "Create a configured monitor",
		Long:  `Create a configured monitor`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateLoadBalancer")
		},
	}

	CreateLoadbalancer.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the loadbalancer")
	CreateLoadbalancer.MarkFlagRequired("zone-id")

	CreateLoadbalancer.Flags().StringVar(&Name, "name", "", "The DNS hostname to associate with your Load Balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.")
	CreateLoadbalancer.MarkFlagRequired("name")

	CreateLoadbalancer.Flags().StringVar(&FallbackPool, "fallback-pool", "", "The pool ID to use when all other pools are detected as unhealthy. max length: 32")
	CreateLoadbalancer.MarkFlagRequired("fallback-pool")

	CreateLoadbalancer.Flags().StringVar(&DefaultPools, "default-pools", "", "A comma separated list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.")
	CreateLoadbalancer.MarkFlagRequired("default-pools")

	CreateLoadbalancer.Flags().BoolVar(&Proxied, "proxied", false, "Whether the hostname should be gray clouded (false) or orange clouded (true). default value: false")

	CreateLoadbalancer.Flags().IntVar(&Ttl, "ttl", 0, "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.")

	var PurgeEverything = &cobra.Command{
		Use:   "purge-everything",
		Short: "Purge Everything",
		Long:  `Purge all items in this zones cache.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "PurgeEverything")
		},
	}

	PurgeEverything.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID that will be purged.")
	PurgeEverything.MarkFlagRequired("zone-id")

	var ActivationCheck = &cobra.Command{
		Use:   "activation-check",
		Short: "Initiate another zone activation check",
		Long:  `Initiates another zone activation check for newly-created zones`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ActivationCheck")
		},
	}

	ActivationCheck.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the activation check")
	ActivationCheck.MarkFlagRequired("zone-id")

	var DescribeZone = &cobra.Command{
		Use:   "describe-zone",
		Short: "Fetches information about a zone.",
		Long:  `Fetches information about a zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneDetails")
		},
	}

	DescribeZone.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID that will be purged.")
	DescribeZone.MarkFlagRequired("zone-id")

	var GetIdByName = &cobra.Command{
		Use:   "get-id-by-name",
		Short: "Get the zone id by name",
		Long:  `Get the zone id by the zone name.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "GetIDByName")
		},
	}

	GetIdByName.Flags().StringVar(&ZoneName, "zone-name", "", "The zoneName that you want the ID of")
	GetIdByName.MarkFlagRequired("zone-name")

	var ListZoneSslSettings = &cobra.Command{
		Use:   "list-zone-ssl-settings",
		Short: "Fetch zone ssl settings",
		Long:  `Get the ssl settings associated with a zone.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneSSLSettings")
		},
	}

	ListZoneSslSettings.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID you wish to fetch the SSL settings for")
	ListZoneSslSettings.MarkFlagRequired("zone-id")

	var GetZoneSettings = &cobra.Command{
		Use:   "get-zone-settings",
		Short: "Get zone specific settings",
		Long:  `Get zone specific settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneSettings")
		},
	}

	GetZoneSettings.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID you wish to fetch the zone settings for")
	GetZoneSettings.MarkFlagRequired("zone-id")

	var Details = &cobra.Command{
		Use:   "details",
		Short: "Get user specific settings",
		Long:  `Get user specific settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UserDetails")
		},
	}

	var BillingProfile = &cobra.Command{
		Use:   "billing-profile",
		Short: "Get the billing profile",
		Long:  `Get the user billing profile.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UserBillingProfile")
		},
	}

	var DescribeVirtualDns = &cobra.Command{
		Use:   "describe-virtual-dns",
		Short: "Get virtual dns details",
		Long:  `Get the details about a virtual dns instance.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "VirtualDNS")
		},
	}

	DescribeVirtualDns.Flags().StringVar(&VirtualDnsId, "virtual-dns-id", "", "The virtualDNS ID you wish to fetch the details of.")
	DescribeVirtualDns.MarkFlagRequired("virtual-dns-id")

	var DeleteVirtualDns = &cobra.Command{
		Use:   "delete-virtual-dns",
		Short: "Delete virtual dns instance",
		Long:  `Delete a specific virtual dns instance.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteVirtualDNS")
		},
	}

	DeleteVirtualDns.Flags().StringVar(&VirtualDnsId, "virtual-dns-id", "", "The virtualDNS ID you wish to delete.")
	DeleteVirtualDns.MarkFlagRequired("virtual-dns-id")

	var DescribePagerule = &cobra.Command{
		Use:   "describe-pagerule",
		Short: "Get page rule details",
		Long:  `Get the details of a specific page rule`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "PageRule")
		},
	}

	DescribePagerule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the page rule")
	DescribePagerule.MarkFlagRequired("zone-id")

	DescribePagerule.Flags().StringVar(&PageruleId, "pagerule-id", "", "The pagerule ID you wish to fetch the details of.")
	DescribePagerule.MarkFlagRequired("pagerule-id")

	var DescribeLoadbalancer = &cobra.Command{
		Use:   "describe-loadbalancer",
		Short: "Get loadbalancer details",
		Long:  `Get loadbalancer details.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "LoadBalancerDetails")
		},
	}

	DescribeLoadbalancer.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the page rule")
	DescribeLoadbalancer.MarkFlagRequired("zone-id")

	DescribeLoadbalancer.Flags().StringVar(&LoadbalancerId, "loadbalancer-id", "", "The loadbalancer id that you wish to view the details of.")
	DescribeLoadbalancer.MarkFlagRequired("loadbalancer-id")

	var DescribeLoadbalancerMonitor = &cobra.Command{
		Use:   "describe-loadbalancer-monitor",
		Short: "Get loadbalancer monitor details",
		Long:  `Get loadbalancer monitor details.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "LoadBalancerMonitorDetails")
		},
	}

	DescribeLoadbalancerMonitor.Flags().StringVar(&MonitorId, "monitor-id", "", "The loadbalancer monitor id that you wish to view the details of.")
	DescribeLoadbalancerMonitor.MarkFlagRequired("monitor-id")

	var DescribeLoadbalancerPool = &cobra.Command{
		Use:   "describe-loadbalancer-pool",
		Short: "Get loadbalancer pool details",
		Long:  `Get loadbalancer pool details.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "LoadBalancerPoolDetails")
		},
	}

	DescribeLoadbalancerPool.Flags().StringVar(&PoolId, "pool-id", "", "The loadbalancer pool id that you wish to view the details of.")
	DescribeLoadbalancerPool.MarkFlagRequired("pool-id")

	var DescribeOrganization = &cobra.Command{
		Use:   "describe-organization",
		Short: "Get organization details",
		Long:  `Get organization details.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OrganizationDetails")
		},
	}

	DescribeOrganization.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id that you wish to view the details of.")
	DescribeOrganization.MarkFlagRequired("organization-id")

	var GetOrganizationInvites = &cobra.Command{
		Use:   "get-organization-invites",
		Short: "Get organization invites",
		Long:  `Get organization invites.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OrganizationInvites")
		},
	}

	GetOrganizationInvites.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id that you wish to view the invites for.")
	GetOrganizationInvites.MarkFlagRequired("organization-id")

	var GetOrganizationMembers = &cobra.Command{
		Use:   "get-organization-members",
		Short: "Get organization members",
		Long:  `Get organization members.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OrganizationMembers")
		},
	}

	GetOrganizationMembers.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id that you wish to view the members of")
	GetOrganizationMembers.MarkFlagRequired("organization-id")

	var GetOrganizationRoles = &cobra.Command{
		Use:   "get-organization-roles",
		Short: "Get organization roles",
		Long:  `Get organization roles.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OrganizationRoles")
		},
	}

	GetOrganizationRoles.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id that you wish to view the member roles of")
	GetOrganizationRoles.MarkFlagRequired("organization-id")

	var ListOriginCerts = &cobra.Command{
		Use:   "list-origin-certs",
		Short: "List origin certificates",
		Long:  `List Origin Certificates associated with a given zone ID`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OriginCertificates")
		},
	}

	ListOriginCerts.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id that you wish to view the origin certs of")
	ListOriginCerts.MarkFlagRequired("zone-id")

	var DescribeOriginCert = &cobra.Command{
		Use:   "describe-origin-cert",
		Short: "Get origin cert details",
		Long:  `Get detailed information about a specific origin certificate`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "OriginCertificate")
		},
	}

	DescribeOriginCert.Flags().StringVar(&CertificateId, "certificate-id", "", "The origin certificate id that you wish to view detailed information about")
	DescribeOriginCert.MarkFlagRequired("certificate-id")

	var DescribeZoneOriginCert = &cobra.Command{
		Use:   "describe-zone-origin-cert",
		Short: "Get zone's origin cert details",
		Long:  `Get detailed information about a specific zone's origin certificate`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "SSLDetails")
		},
	}

	DescribeZoneOriginCert.Flags().StringVar(&CertificateId, "certificate-id", "", "The zone's certificate id that you wish to view detailed information about")
	DescribeZoneOriginCert.MarkFlagRequired("certificate-id")

	DescribeZoneOriginCert.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id that you wish to view detailed information about")
	DescribeZoneOriginCert.MarkFlagRequired("zone-id")

	var DescribeRailgun = &cobra.Command{
		Use:   "describe-railgun",
		Short: "Get railgun instance details",
		Long:  `Get detailed information about a specific railgun`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "RailgunDetails")
		},
	}

	DescribeRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun id that you wish to view detailed information about")
	DescribeRailgun.MarkFlagRequired("railgun-id")

	var GetRailgunZones = &cobra.Command{
		Use:   "get-railgun-zones",
		Short: "Get railgun instance zone details",
		Long:  `Get detailed information about a specific railgun's zones`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "RailgunZones")
		},
	}

	GetRailgunZones.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun id that you wish to view associated zones")
	GetRailgunZones.MarkFlagRequired("railgun-id")

	var DescribeRatelimit = &cobra.Command{
		Use:   "describe-ratelimit",
		Short: "Get detailed information about a zone",
		Long:  `Get detailed information about a specific zone's ratelimits`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "RateLimit")
		},
	}

	DescribeRatelimit.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the ratelimit")
	DescribeRatelimit.MarkFlagRequired("zone-id")

	DescribeRatelimit.Flags().StringVar(&RatelimitId, "ratelimit-id", "", "The ratelimit id that you wish to view detailed information about")
	DescribeRatelimit.MarkFlagRequired("ratelimit-id")

	var RevokeOriginCert = &cobra.Command{
		Use:   "revoke-origin-cert",
		Short: "Revoke origin certificate",
		Long:  `Revoke a specific origin certificate`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "RevokeOriginCertificate")
		},
	}

	RevokeOriginCert.Flags().StringVar(&CertificateId, "certificate-id", "", "The certificate id that is being revoked")
	RevokeOriginCert.MarkFlagRequired("certificate-id")

	var TestRailgunConnection = &cobra.Command{
		Use:   "test-railgun-connection",
		Short: "List User Access Rules",
		Long:  `Returns all access rules associated with your account`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "TestRailgunConnection")
		},
	}

	TestRailgunConnection.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the railgun connection you wish to test")
	TestRailgunConnection.MarkFlagRequired("zone-id")

	TestRailgunConnection.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun id associated with the railgun connection you wish to test")
	TestRailgunConnection.MarkFlagRequired("railgun-id")

	var DescribeZoneRailgun = &cobra.Command{
		Use:   "describe-zone-railgun",
		Short: "Get zone railgun details",
		Long:  `Returns all railgun details for an associated zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneRailgunDetails")
		},
	}

	DescribeZoneRailgun.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the railguns you wish to get")
	DescribeZoneRailgun.MarkFlagRequired("zone-id")

	DescribeZoneRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun id associated with the railguns you wish to get")
	DescribeZoneRailgun.MarkFlagRequired("railgun-id")

	var DescribeCustomHostname = &cobra.Command{
		Use:   "describe-custom-hostname",
		Short: "Custom hostname details",
		Long:  `Returns details associated with the custom hostname id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CustomHostname")
		},
	}

	DescribeCustomHostname.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the custom hostname being returned")
	DescribeCustomHostname.MarkFlagRequired("zone-id")

	DescribeCustomHostname.Flags().StringVar(&CustomHostnameId, "custom-hostname-id", "", "The custom hostname id associated with the custom hostname you wish to describe")
	DescribeCustomHostname.MarkFlagRequired("custom-hostname-id")

	var DescribeCustomHostnameByName = &cobra.Command{
		Use:   "describe-custom-hostname-by-name",
		Short: "Custom hostname details",
		Long:  `Returns details associated with the custom hostname id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CustomHostnameIDByName")
		},
	}

	DescribeCustomHostnameByName.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the custom hostname id being returned")
	DescribeCustomHostnameByName.MarkFlagRequired("zone-id")

	DescribeCustomHostnameByName.Flags().StringVar(&Name, "name", "", "The custom hostname associated with the custom hostname id you wish to return")
	DescribeCustomHostnameByName.MarkFlagRequired("name")

	var SetPaused = &cobra.Command{
		Use:   "set-paused",
		Short: "Pause or unpause a zone",
		Long:  `Set or unset a zone from being paused`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneSetPaused")
		},
	}

	SetPaused.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id to either pause or unpause")
	SetPaused.MarkFlagRequired("zone-id")

	SetPaused.Flags().BoolVar(&Paused, "paused", false, "This flag is unset for unpaused. Set for paused.")

	var DeletePagerule = &cobra.Command{
		Use:   "delete-pagerule",
		Short: "Delete a specific page rule",
		Long:  `Delete a page rule associated with a specific zone and pagerule id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeletePageRule")
		},
	}

	DeletePagerule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the page rule")
	DeletePagerule.MarkFlagRequired("zone-id")

	DeletePagerule.Flags().StringVar(&PageruleId, "pagerule-id", "", "The pagerule ID associated with the pagerule you wish to delete")
	DeletePagerule.MarkFlagRequired("pagerule-id")

	var DeleteRailgun = &cobra.Command{
		Use:   "delete-railgun",
		Short: "Delete a specific railgun by its id",
		Long:  `Delete a railgun associated with a specific railgun id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteRailgun")
		},
	}

	DeleteRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun ID associated with the railgun being deleted")
	DeleteRailgun.MarkFlagRequired("railgun-id")

	var DisableRailgun = &cobra.Command{
		Use:   "disable-railgun",
		Short: "Delete a specific railgun by its id",
		Long:  `Disable a railgun associated with a specific railgun id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DisableRailgun")
		},
	}

	DisableRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun ID associated with the railgun being disabled")
	DisableRailgun.MarkFlagRequired("railgun-id")

	var DisconnectRailgun = &cobra.Command{
		Use:   "disconnect-railgun",
		Short: "Disconnect a railgun from a zone",
		Long:  `Disconnect a railgun associated with a specific zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DisconnectZoneRailgun")
		},
	}

	DisconnectRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun ID associated with the railgun being disconnected")
	DisconnectRailgun.MarkFlagRequired("railgun-id")

	DisconnectRailgun.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID that the railgun instance is being disconnected from")
	DisconnectRailgun.MarkFlagRequired("zone-id")

	var EnableRailgun = &cobra.Command{
		Use:   "enable-railgun",
		Short: "Enable a specific disabled railgun by its id",
		Long:  `Enable a railgun associated with a specific railgun id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "EnableRailgun")
		},
	}

	EnableRailgun.Flags().StringVar(&RailgunId, "railgun-id", "", "The railgun ID associated with the railgun being enabled ")
	EnableRailgun.MarkFlagRequired("railgun-id")

	var DeleteRatelimit = &cobra.Command{
		Use:   "delete-ratelimit",
		Short: "Delete a specific ratelimit by its id",
		Long:  `Delete a ratelimit associated with a specific ratelimit id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteRateLimit")
		},
	}

	DeleteRatelimit.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the ratelimit being deleted")
	DeleteRatelimit.MarkFlagRequired("zone-id")

	DeleteRatelimit.Flags().StringVar(&RatelimitId, "ratelimit-id", "", "The ratelimit ID associated with the ratelimit being deleted")
	DeleteRatelimit.MarkFlagRequired("ratelimit-id")

	var DeleteCustomCert = &cobra.Command{
		Use:   "delete-custom-cert",
		Short: "Delete a specific certificate by its id",
		Long:  `Delete a custom certificate associated with a specific certificate id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteSSL")
		},
	}

	DeleteCustomCert.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the certificate being deleted")
	DeleteCustomCert.MarkFlagRequired("zone-id")

	DeleteCustomCert.Flags().StringVar(&CertificateId, "certificate-id", "", "The certificate ID associated with the certificate being deleted")
	DeleteCustomCert.MarkFlagRequired("certificate-id")

	var DeleteCustomHostname = &cobra.Command{
		Use:   "delete-custom-hostname",
		Short: "Delete a specific custom hostname",
		Long:  `Delete a custom hostname associated with a specific zone id`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteCustomHostname")
		},
	}

	DeleteCustomHostname.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the custom hostname being deleted")
	DeleteCustomHostname.MarkFlagRequired("zone-id")

	DeleteCustomHostname.Flags().StringVar(&CustomHostnameId, "custom-hostname-id", "", "The custom hostname ID associated with the custom hostname being deleted")
	DeleteCustomHostname.MarkFlagRequired("custom-hostname-id")

	var DeleteLoadbalancer = &cobra.Command{
		Use:   "delete-loadbalancer",
		Short: "Delete a specific loadbalancer",
		Long:  `Delete a loadbalancer associated with a specific zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteLoadBalancer")
		},
	}

	DeleteLoadbalancer.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the loadbalancer")
	DeleteLoadbalancer.MarkFlagRequired("zone-id")

	DeleteLoadbalancer.Flags().StringVar(&LoadbalancerId, "loadbalancer-id", "", "The loadbalancer-id associated with the custom hostname being deleted")
	DeleteLoadbalancer.MarkFlagRequired("loadbalancer-id")

	var DeleteLoadbalancerMonitor = &cobra.Command{
		Use:   "delete-loadbalancer-monitor",
		Short: "Delete a specific loadbalancer monitor",
		Long:  `Delete a specific loadbalancer monitor`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteLoadBalancerMonitor")
		},
	}

	DeleteLoadbalancerMonitor.Flags().StringVar(&MonitorId, "monitor-id", "", "The load balancer monitor ID associated with the monitor being deleted")
	DeleteLoadbalancerMonitor.MarkFlagRequired("monitor-id")

	var DeleteLoadbalancerPool = &cobra.Command{
		Use:   "delete-loadbalancer-pool",
		Short: "Delete a specific loadbalancer pool",
		Long:  `Delete a specific loadbalancer pool`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteLoadBalancerPool")
		},
	}

	DeleteLoadbalancerPool.Flags().StringVar(&PoolId, "pool-id", "", "The load balancer pool ID associated with the pool being deleted")
	DeleteLoadbalancerPool.MarkFlagRequired("pool-id")

	var DeleteOrganizationAccessRule = &cobra.Command{
		Use:   "delete-organization-access-rule",
		Short: "Delete a specific access rule",
		Long:  `Delete an access rule associated with a specific organization`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteOrganizationAccessRule")
		},
	}

	DeleteOrganizationAccessRule.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization ID associated with the access rule being deleted")
	DeleteOrganizationAccessRule.MarkFlagRequired("organization-id")

	DeleteOrganizationAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The access rule ID associated with the access rule being deleted")
	DeleteOrganizationAccessRule.MarkFlagRequired("access-rule-id")

	var CreateRailgun = &cobra.Command{
		Use:   "create-railgun",
		Short: "Create a railgun",
		Long:  `Create a railgun`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateRailgun")
		},
	}

	CreateRailgun.Flags().StringVar(&Name, "name", "", "The name you are assigning to the newly created railgun")
	CreateRailgun.MarkFlagRequired("name")

	var DeleteUserAccessRule = &cobra.Command{
		Use:   "delete-user-access-rule",
		Short: "Delete a user access rule",
		Long:  `Delete a specific user access rule`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteUserAccessRule")
		},
	}

	DeleteUserAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The access rule id associated with the user access rule you are deleting")
	DeleteUserAccessRule.MarkFlagRequired("access-rule-id")

	var DeleteUserAgentRule = &cobra.Command{
		Use:   "delete-user-agent-rule",
		Short: "Delete a user agent rule",
		Long:  `Delete a specific user agent rule`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteUserAgentRule")
		},
	}

	DeleteUserAgentRule.Flags().StringVar(&UserAgentId, "user-agent-id", "", "The user agent rule id associated with the user agent rule being deleted")
	DeleteUserAgentRule.MarkFlagRequired("user-agent-id")

	DeleteUserAgentRule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the user agent rule being deleted")
	DeleteUserAgentRule.MarkFlagRequired("zone-id")

	var DeleteZoneAccessRule = &cobra.Command{
		Use:   "delete-zone-access-rule",
		Short: "Delete a zone access rule",
		Long:  `Delete a specific zone access rule`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteZoneAccessRule")
		},
	}

	DeleteZoneAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The zone access rule id associated with the zone access rule being deleted")
	DeleteZoneAccessRule.MarkFlagRequired("access-rule-id")

	DeleteZoneAccessRule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the zone access rule being deleted")
	DeleteZoneAccessRule.MarkFlagRequired("zone-id")

	var DeleteZoneLockdown = &cobra.Command{
		Use:   "delete-zone-lockdown",
		Short: "Delete a zone lockdown",
		Long:  `Delete a specific zone lockdown`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteZoneLockdown")
		},
	}

	DeleteZoneLockdown.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the specific zone lockdown rule")
	DeleteZoneLockdown.MarkFlagRequired("zone-id")

	DeleteZoneLockdown.Flags().StringVar(&LockdownId, "lockdown-id", "", "The zone lockdown id associated with the zone lockdown being deleted")
	DeleteZoneLockdown.MarkFlagRequired("lockdown-id")

	var AnalyticsByColo = &cobra.Command{
		Use:   "analytics-by-colo",
		Short: "Get analytics by colo",
		Long:  `Retrieve zone analytics structured by colocation`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneAnalyticsByColocation")
		},
	}

	AnalyticsByColo.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id that analytics are being retreived for")
	AnalyticsByColo.MarkFlagRequired("zone-id")

	AnalyticsByColo.Flags().StringVar(&Since, "since", "", "String timestamp of the analytics start time")

	AnalyticsByColo.Flags().StringVar(&Until, "until", "", "String timestamp of the analytics end time")

	AnalyticsByColo.Flags().BoolVar(&Continuous, "continuous", false, "When continuous is true and since or until is set, the api will only return completely aggregated results")

	var AnalyticsDashboard = &cobra.Command{
		Use:   "analytics-dashboard",
		Short: "Get an analytics overview by zone",
		Long:  `Retrieve zone analytics overview`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ZoneAnalyticsDashboard")
		},
	}

	AnalyticsDashboard.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id that analytics are being retreived for")
	AnalyticsDashboard.MarkFlagRequired("zone-id")

	AnalyticsDashboard.Flags().StringVar(&Since, "since", "", "String timestamp of the analytics start time")

	AnalyticsDashboard.Flags().StringVar(&Until, "until", "", "String timestamp of the analytics end time")

	AnalyticsDashboard.Flags().BoolVar(&Continuous, "continuous", false, "When continuous is true and since or until is set, the api will only return completely aggregated results")

	var EditUser = &cobra.Command{
		Use:   "edit-user",
		Short: "Edit user account details",
		Long:  `Edit user account details`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateUser")
		},
	}

	EditUser.Flags().StringVar(&FirstName, "first-name", "", "User's first name, max length 60")

	EditUser.Flags().StringVar(&LastName, "last-name", "", "User's last name, max length 60")

	EditUser.Flags().StringVar(&Telephone, "telephone", "", "User's telephone number, max length 20")

	EditUser.Flags().StringVar(&Country, "country", "", "The country in which the user lives, example US. max length 30")

	EditUser.Flags().StringVar(&Zipcode, "zipcode", "", "The zip code or postal code in which the user lives, max length 20")

	var CreateLoadbalancerPool = &cobra.Command{
		Use:   "create-loadbalancer-pool",
		Short: "Create a loadbalancer pool",
		Long:  `Create a new loadbalancer pool`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateLoadBalancerPool")
		},
	}

	CreateLoadbalancerPool.Flags().StringVar(&Name, "name", "", "The name of the loadbalancer pool")
	CreateLoadbalancerPool.MarkFlagRequired("name")

	CreateLoadbalancerPool.Flags().StringVar(&Origins, "origins", "", "The list of origins Example: [{\"name\": \"app-server-1\", \"address\": \"0.0.0.0\", \"enabled\": true, \"weight\": 0.56}]")
	CreateLoadbalancerPool.MarkFlagRequired("origins")

	CreateLoadbalancerPool.Flags().StringVar(&Description, "description", "", "A human-readable description of the pool.")

	CreateLoadbalancerPool.Flags().BoolVar(&Disabled, "disabled", false, "By default, the pool will be enabled. Specify disabled in order to modify this default")

	CreateLoadbalancerPool.Flags().IntVar(&MinimumOrigins, "minimum-origins", 0, "The minimum number of origins that must be healthy for this pool to serve traffic. ")

	CreateLoadbalancerPool.Flags().StringVar(&Monitor, "monitor", "", "The ID of the Monitor to use for health checking origins within this pool.")

	CreateLoadbalancerPool.Flags().StringVar(&NotificationEmail, "notification-email", "", "The email address to send health status notifications to. This can be an individual mailbox or a mailing list.")

	var UpdateLoadbalancerPool = &cobra.Command{
		Use:   "update-loadbalancer-pool",
		Short: "Edit an existing loadbalancer pool",
		Long:  `Edit an existing loadbalancer pool`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ModifyLoadBalancerPool")
		},
	}

	UpdateLoadbalancerPool.Flags().StringVar(&PoolId, "pool-id", "", "The name of the loadbalancer pool")
	UpdateLoadbalancerPool.MarkFlagRequired("pool-id")

	UpdateLoadbalancerPool.Flags().StringVar(&Name, "name", "", "The name of the loadbalancer pool")
	UpdateLoadbalancerPool.MarkFlagRequired("name")

	UpdateLoadbalancerPool.Flags().StringVar(&Origins, "origins", "", "The list of origins Example: [{\"name\": \"app-server-1\", \"address\": \"0.0.0.0\", \"enabled\": true, \"weight\": 0.56}]")
	UpdateLoadbalancerPool.MarkFlagRequired("origins")

	UpdateLoadbalancerPool.Flags().StringVar(&Description, "description", "", "A human-readable description of the pool.")

	UpdateLoadbalancerPool.Flags().BoolVar(&Disabled, "disabled", false, "By default, the pool will be enabled. Specify disabled in order to modify this default")

	UpdateLoadbalancerPool.Flags().IntVar(&MinimumOrigins, "minimum-origins", 0, "The minimum number of origins that must be healthy for this pool to serve traffic. ")

	UpdateLoadbalancerPool.Flags().StringVar(&Monitor, "monitor", "", "The ID of the Monitor to use for health checking origins within this pool.")

	UpdateLoadbalancerPool.Flags().StringVar(&NotificationEmail, "notification-email", "", "The email address to send health status notifications to. This can be an individual mailbox or a mailing list.")

	var UpdateZoneSettings = &cobra.Command{
		Use:   "update-zone-settings",
		Short: "Edit a zones settings",
		Long:  `Edit a zones settings`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateZoneSettings")
		},
	}

	UpdateZoneSettings.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the settings being modified")
	UpdateZoneSettings.MarkFlagRequired("zone-id")

	UpdateZoneSettings.Flags().StringVar(&ZoneSettingsObject, "zone-settings-object", "", "One or more zone setting objects. Must contain an ID and a value. Example: [{\"id\": \"always_online\",\"value\": \"on\"}]")
	UpdateZoneSettings.MarkFlagRequired("zone-settings-object")

	var UpdateZoneLockdown = &cobra.Command{
		Use:   "update-zone-lockdown",
		Short: "Edit an existing zone lockdown",
		Long:  `Edit an existing zone lockdown`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateZoneLockdown")
		},
	}

	UpdateZoneLockdown.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the zone lockdown")
	UpdateZoneLockdown.MarkFlagRequired("zone-id")

	UpdateZoneLockdown.Flags().StringVar(&LockdownId, "lockdown-id", "", "The lockdown id associated with the zone lockdown")
	UpdateZoneLockdown.MarkFlagRequired("lockdown-id")

	UpdateZoneLockdown.Flags().StringVar(&Configuration, "configuration", "", "The new configuration associated with the lockdown - Example: [{\"target\": \"ip\",\"value\": \"198.51.100.4\"}]")
	UpdateZoneLockdown.MarkFlagRequired("configuration")

	UpdateZoneLockdown.Flags().StringVar(&Urls, "urls", "", "Comma delimited list of URLs associated with the zone lockdown")
	UpdateZoneLockdown.MarkFlagRequired("urls")

	UpdateZoneLockdown.Flags().BoolVar(&Paused, "paused", false, "Whether this zone lockdown is currently paused")

	UpdateZoneLockdown.Flags().StringVar(&Description, "description", "", "A note that you can use to describe the reason for a Lockdown rule")

	var CreateZoneLockdown = &cobra.Command{
		Use:   "create-zone-lockdown",
		Short: "Create a new zone lockdown",
		Long:  `Create a new zone lockdown`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateZoneLockdown")
		},
	}

	CreateZoneLockdown.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the zone lockdown")
	CreateZoneLockdown.MarkFlagRequired("zone-id")

	CreateZoneLockdown.Flags().StringVar(&Configuration, "configuration", "", "The new configuration associated with the lockdown - Example: [{\"target\": \"ip\",\"value\": \"198.51.100.4\"}]")
	CreateZoneLockdown.MarkFlagRequired("configuration")

	CreateZoneLockdown.Flags().StringVar(&Urls, "urls", "", "Comma delimited list of URLs associated with the zone lockdown")
	CreateZoneLockdown.MarkFlagRequired("urls")

	CreateZoneLockdown.Flags().BoolVar(&Paused, "paused", false, "Whether this zone lockdown is currently paused")

	CreateZoneLockdown.Flags().StringVar(&Description, "description", "", "A note that you can use to describe the reason for a Lockdown rule")

	var CreateVirtualDns = &cobra.Command{
		Use:   "create-virtual-dns",
		Short: "Create a new virtual dns cluster",
		Long:  `Create a new virtual dns cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateVirtualDNS")
		},
	}

	CreateVirtualDns.Flags().StringVar(&Name, "name", "", "Virtual DNS Cluster Name, max length: 160 characters")
	CreateVirtualDns.MarkFlagRequired("name")

	CreateVirtualDns.Flags().StringVar(&OriginIps, "origin-ips", "", "Comma delimited list of origin IP addresses - Example: [\"192.0.2.1\",\"198.51.100.1\",\"2001:DB8:100::CF\"]")
	CreateVirtualDns.MarkFlagRequired("origin-ips")

	CreateVirtualDns.Flags().IntVar(&MinimumCacheTtl, "minimum-cache-ttl", 0, "Minimum DNS Cache TTL. default value: 60, min value:30, max value:36000")

	CreateVirtualDns.Flags().IntVar(&MaximumCacheTtl, "maximum-cache-ttl", 0, "Maximum DNS Cache TTL. default value: 900, min value:30, max value:36000")

	CreateVirtualDns.Flags().BoolVar(&DeprecateAnyRequest, "deprecate-any-request", false, "Deprecate the response to ANY requests")

	var UpdateVirtualDns = &cobra.Command{
		Use:   "update-virtual-dns",
		Short: "Update a virtual dns cluster",
		Long:  `Update a virtual dns cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateVirtualDNS")
		},
	}

	UpdateVirtualDns.Flags().StringVar(&VirtualDnsId, "virtual-dns-id", "", "The virtual DNS id being modified")
	UpdateVirtualDns.MarkFlagRequired("virtual-dns-id")

	UpdateVirtualDns.Flags().StringVar(&OriginIps, "origin-ips", "", "Comma delimited list of origin IP addresses - Example: [\"192.0.2.1\",\"198.51.100.1\",\"2001:DB8:100::CF\"]")

	UpdateVirtualDns.Flags().IntVar(&MinimumCacheTtl, "minimum-cache-ttl", 0, "Minimum DNS Cache TTL. default value: 60, min value:30, max value:36000")

	UpdateVirtualDns.Flags().IntVar(&MaximumCacheTtl, "maximum-cache-ttl", 0, "Maximum DNS Cache TTL. default value: 900, min value:30, max value:36000")

	UpdateVirtualDns.Flags().BoolVar(&DeprecateAnyRequest, "deprecate-any-request", false, "Deprecate the response to ANY requests")

	var CreatePagerule = &cobra.Command{
		Use:   "create-pagerule",
		Short: "Create a new page rule",
		Long:  `Create a new page rule associated with a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreatePageRule")
		},
	}

	CreatePagerule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the new page rule")
	CreatePagerule.MarkFlagRequired("zone-id")

	CreatePagerule.Flags().StringVar(&Targets, "targets", "", "List of page rule targets. Examples: '[{\"target\": \"url\",\"constraint\": {\"operator\": \"matches\",\"value\": \"*example.com/images/*\"}}]'")
	CreatePagerule.MarkFlagRequired("targets")

	CreatePagerule.Flags().StringVar(&Actions, "actions", "", "List of page rule actions. Examples: '[{\"id\": \"always_online\",\"value\": \"on\"}]'")
	CreatePagerule.MarkFlagRequired("actions")

	CreatePagerule.Flags().IntVar(&Priority, "priority", 0, "A number that indicates the preference for a page rule over another. default value: 1")

	CreatePagerule.Flags().StringVar(&Status, "status", "", "Status of the page rule. default value: disabled valid values: active, disabled required = false")
	CreatePagerule.MarkFlagRequired("status")

	var UpdatePagerule = &cobra.Command{
		Use:   "update-pagerule",
		Short: "Update a new page rule",
		Long:  `Update a new page rule associated with a zone and page rule`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ChangePageRule")
		},
	}

	UpdatePagerule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the updated page rule")
	UpdatePagerule.MarkFlagRequired("zone-id")

	UpdatePagerule.Flags().StringVar(&PageruleId, "pagerule-id", "", "The pagerule id associated with the updated page rule")
	UpdatePagerule.MarkFlagRequired("pagerule-id")

	UpdatePagerule.Flags().StringVar(&Targets, "targets", "", "List of page rule targets. Examples: '[{\"target\": \"url\",\"constraint\": {\"operator\": \"matches\",\"value\": \"*example.com/images/*\"}}]'")
	UpdatePagerule.MarkFlagRequired("targets")

	UpdatePagerule.Flags().StringVar(&Actions, "actions", "", "List of page rule actions. Examples: '[{\"id\": \"always_online\",\"value\": \"on\"}]'")
	UpdatePagerule.MarkFlagRequired("actions")

	UpdatePagerule.Flags().IntVar(&Priority, "priority", 0, "A number that indicates the preference for a page rule over another. default value: 1")

	UpdatePagerule.Flags().StringVar(&Status, "status", "", "Status of the page rule. default value: disabled valid values: active, disabled required = false")
	UpdatePagerule.MarkFlagRequired("status")

	var CreateOrganizationAccessRule = &cobra.Command{
		Use:   "create-organization-access-rule",
		Short: "Create an organization access rule",
		Long:  `Make a new IP, IP range, or country access rule for all zones owned by the organization. Note: If you would like to create an access rule that applies to a specific zone only, use the zone firewall endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateOrganizationAccessRule")
		},
	}

	CreateOrganizationAccessRule.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the new access rule")
	CreateOrganizationAccessRule.MarkFlagRequired("organization-id")

	CreateOrganizationAccessRule.Flags().StringVar(&Mode, "mode", "", "The action to apply to a matched request valid values: block, challenge, whitelist, js_challenge")
	CreateOrganizationAccessRule.MarkFlagRequired("mode")

	CreateOrganizationAccessRule.Flags().StringVar(&Configuration, "configuration", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")
	CreateOrganizationAccessRule.MarkFlagRequired("configuration")

	CreateOrganizationAccessRule.Flags().StringVar(&Notes, "notes", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")

	var CreateOriginCert = &cobra.Command{
		Use:   "create-origin-cert",
		Short: "Create a Cloudflare-signed certificate",
		Long:  `Create a Cloudflare-signed certificate.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateOriginCertificate")
		},
	}

	CreateOriginCert.Flags().StringVar(&Hostnames, "hostnames", "", "Comma-delimited list of hostnames or wildcard names (e.g., *.example.com) bound to the certificate")
	CreateOriginCert.MarkFlagRequired("hostnames")

	CreateOriginCert.Flags().IntVar(&RequestValidity, "request-validity", 0, "The number of days for which the certificate should be valid. default value: 5475, valid values: 7, 30, 90, 365, 730, 1095, 5475")
	CreateOriginCert.MarkFlagRequired("request-validity")

	CreateOriginCert.Flags().StringVar(&RequestType, "request-type", "", "Signature type desired on certificate (\"origin-rsa\" (rsa), \"origin-ecc\" (ecdsa), or \"keyless-certificate\" (for Keyless SSL servers) valid values: origin-rsa, origin-ecc, keyless-certificate")
	CreateOriginCert.MarkFlagRequired("request-type")

	CreateOriginCert.Flags().StringVar(&Csr, "csr", "", "The Certificate Signing Request (CSR). Must be newline-encoded. -----BEGIN CERTIFICATE REQUEST-----\nMIICxzCCA...\n-----END CERTIFICATE REQUEST-----")
	CreateOriginCert.MarkFlagRequired("csr")

	var CreateRatelimit = &cobra.Command{
		Use:   "create-ratelimit",
		Short: "Creates a new rate limit for a zone.",
		Long:  `Create a new rate limit for a zone. See the record object definitions for required attributes for each record type`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateRateLimit")
		},
	}

	CreateRatelimit.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the newly created ratelimit")
	CreateRatelimit.MarkFlagRequired("zone-id")

	CreateRatelimit.Flags().StringVar(&Match, "match", "", "The match object described by https://api.cloudflare.com/#rate-limits-for-a-zone-create-rate-limit")
	CreateRatelimit.MarkFlagRequired("match")

	CreateRatelimit.Flags().IntVar(&Threshold, "threshold", 0, "The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period min value:2 max value:1000000")
	CreateRatelimit.MarkFlagRequired("threshold")

	CreateRatelimit.Flags().IntVar(&Period, "period", 0, "The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed. min value:1 max value:86400")
	CreateRatelimit.MarkFlagRequired("period")

	CreateRatelimit.Flags().StringVar(&Action, "action", "", "The action to be performed when the threshold of matched traffic within the period defined is exceeded '{\"mode\": \"challenge\",\"timeout\": 86400,\"response\": {\"content_type\": \"text/xml\",\"body\": \"<error>This request has been rate-limited.</error>\"}}'")
	CreateRatelimit.MarkFlagRequired("action")

	CreateRatelimit.Flags().BoolVar(&Enabled, "enabled", false, "Whether this ratelimit is currently enabled or not.")

	CreateRatelimit.Flags().StringVar(&Description, "description", "", "A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed max length: 1024")

	CreateRatelimit.Flags().StringVar(&Bypass, "bypass", "", "Criteria that would allow the rate limit to be bypassed, for example to express that you shouldn't apply a rate limit to a given set of URLs '[{\"name\": \"url\",\"value\": \"api.example.com/*\"}]'")

	var UpdateRatelimit = &cobra.Command{
		Use:   "update-ratelimit",
		Short: "Update a rate limit for a zone.",
		Long:  `Update a specific rate limit.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateRateLimit")
		},
	}

	UpdateRatelimit.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the newly created ratelimit")
	UpdateRatelimit.MarkFlagRequired("zone-id")

	UpdateRatelimit.Flags().StringVar(&LimitId, "limit-id", "", "The ratelimit id associated with the newly created ratelimit")
	UpdateRatelimit.MarkFlagRequired("limit-id")

	UpdateRatelimit.Flags().StringVar(&Match, "match", "", "The match object described by https://api.cloudflare.com/#rate-limits-for-a-zone-create-rate-limit")
	UpdateRatelimit.MarkFlagRequired("match")

	UpdateRatelimit.Flags().IntVar(&Threshold, "threshold", 0, "The threshold that triggers the rate limit mitigations, combine with period. i.e. threshold per period min value:2 max value:1000000")
	UpdateRatelimit.MarkFlagRequired("threshold")

	UpdateRatelimit.Flags().IntVar(&Period, "period", 0, "The time in seconds to count matching traffic. If the count exceeds threshold within this period the action will be performed. min value:1 max value:86400")
	UpdateRatelimit.MarkFlagRequired("period")

	UpdateRatelimit.Flags().StringVar(&Action, "action", "", "The action to be performed when the threshold of matched traffic within the period defined is exceeded '{\"mode\": \"challenge\",\"timeout\": 86400,\"response\": {\"content_type\": \"text/xml\",\"body\": \"<error>This request has been rate-limited.</error>\"}}'")
	UpdateRatelimit.MarkFlagRequired("action")

	UpdateRatelimit.Flags().BoolVar(&Enabled, "enabled", false, "Whether this ratelimit is currently enabled or not.")

	UpdateRatelimit.Flags().StringVar(&Description, "description", "", "A note that you can use to describe the reason for a rate limit. This value is sanitized and all tags are removed max length: 1024")

	UpdateRatelimit.Flags().StringVar(&Bypass, "bypass", "", "Criteria that would allow the rate limit to be bypassed, for example to express that you shouldn't apply a rate limit to a given set of URLs '[{\"name\": \"url\",\"value\": \"api.example.com/*\"}]'")

	var CreateUserAccessRule = &cobra.Command{
		Use:   "create-user-access-rule",
		Short: "Create a user access rule",
		Long:  `Make a new IP, IP range, or country access rule for all zones owned by the user. Note: If you would like to create an access rule that applies to a specific zone only, use the zone firewall endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateUserAccessRule")
		},
	}

	CreateUserAccessRule.Flags().StringVar(&Mode, "mode", "", "The action to apply to a matched request valid values: block, challenge, whitelist, js_challenge")
	CreateUserAccessRule.MarkFlagRequired("mode")

	CreateUserAccessRule.Flags().StringVar(&Configuration, "configuration", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")
	CreateUserAccessRule.MarkFlagRequired("configuration")

	CreateUserAccessRule.Flags().StringVar(&Notes, "notes", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")

	var UpdateUserAccessRule = &cobra.Command{
		Use:   "update-user-access-rule",
		Short: "Create a user access rule",
		Long:  `Update an IP, IP range, or country access rule for all zones owned by the user. Note: If you would like to create an access rule that applies to a specific zone only, use the zone firewall endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateUserAccessRule")
		},
	}

	UpdateUserAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The access rule id associated with the rule being updated")
	UpdateUserAccessRule.MarkFlagRequired("access-rule-id")

	UpdateUserAccessRule.Flags().StringVar(&Mode, "mode", "", "The action to apply to a matched request valid values: block, challenge, whitelist, js_challenge")
	UpdateUserAccessRule.MarkFlagRequired("mode")

	UpdateUserAccessRule.Flags().StringVar(&Configuration, "configuration", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")
	UpdateUserAccessRule.MarkFlagRequired("configuration")

	UpdateUserAccessRule.Flags().StringVar(&Notes, "notes", "", "Set the access rule's human readable note")

	var UpdateZoneAccessRule = &cobra.Command{
		Use:   "update-zone-access-rule",
		Short: "Create a user access rule",
		Long:  `Update an IP, IP range, or country access rule for all zones owned by the user. Note: If you would like to create an access rule that applies to a specific zone only, use the zone firewall endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateZoneAccessRule")
		},
	}

	UpdateZoneAccessRule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the rule being updated")
	UpdateZoneAccessRule.MarkFlagRequired("zone-id")

	UpdateZoneAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The access rule id associated with the rule being updated")
	UpdateZoneAccessRule.MarkFlagRequired("access-rule-id")

	UpdateZoneAccessRule.Flags().StringVar(&Mode, "mode", "", "The action to apply to a matched request valid values: block, challenge, whitelist, js_challenge")
	UpdateZoneAccessRule.MarkFlagRequired("mode")

	UpdateZoneAccessRule.Flags().StringVar(&Configuration, "configuration", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")
	UpdateZoneAccessRule.MarkFlagRequired("configuration")

	UpdateZoneAccessRule.Flags().StringVar(&Notes, "notes", "", "Set the access rule's human readable note")

	var UpdateOrganizationAccessRule = &cobra.Command{
		Use:   "update-organization-access-rule",
		Short: "Create a user access rule",
		Long:  `Update an IP, IP range, or country access rule for all zones owned by the user. Note: If you would like to create an access rule that applies to a specific zone only, use the zone firewall endpoints.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateOrganizationAccessRule")
		},
	}

	UpdateOrganizationAccessRule.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the rule being updated")
	UpdateOrganizationAccessRule.MarkFlagRequired("organization-id")

	UpdateOrganizationAccessRule.Flags().StringVar(&AccessRuleId, "access-rule-id", "", "The access rule id associated with the rule being updated")
	UpdateOrganizationAccessRule.MarkFlagRequired("access-rule-id")

	UpdateOrganizationAccessRule.Flags().StringVar(&Mode, "mode", "", "The action to apply to a matched request valid values: block, challenge, whitelist, js_challenge")
	UpdateOrganizationAccessRule.MarkFlagRequired("mode")

	UpdateOrganizationAccessRule.Flags().StringVar(&Configuration, "configuration", "", "Rule configuration. Example {\"target\": \"ip\",\"value\": \"198.51.100.4\"}")
	UpdateOrganizationAccessRule.MarkFlagRequired("configuration")

	UpdateOrganizationAccessRule.Flags().StringVar(&Notes, "notes", "", "Set the access rule's human readable note")

	var ListZoneAccessRules = &cobra.Command{
		Use:   "list-zone-access-rules",
		Short: "List Zone Access Rules",
		Long:  `Returns all access rules associated with the zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZoneAccessRules")
		},
	}

	ListZoneAccessRules.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the access rule")
	ListZoneAccessRules.MarkFlagRequired("zone-id")

	ListZoneAccessRules.Flags().StringVar(&Notes, "notes", "", "Matching any string within previously created access rules with the notes")

	ListZoneAccessRules.Flags().StringVar(&Mode, "mode", "", "valid values: block, challenge, whitelist, js_challenge")

	ListZoneAccessRules.Flags().IntVar(&Page, "page", 0, "Requested page within paginated list of results")

	var UploadCustomCert = &cobra.Command{
		Use:   "upload-custom-cert",
		Short: "Upload a custom SSL cert",
		Long:  `Upload a new SSL certificate for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateSSL")
		},
	}

	UploadCustomCert.Flags().StringVar(&ZoneId, "zone-id", "", "The zone-id associated with the custom cert")
	UploadCustomCert.MarkFlagRequired("zone-id")

	UploadCustomCert.Flags().StringVar(&Certificate, "certificate", "", "The zone's SSL certificate or certificate and the intermediate(s)")
	UploadCustomCert.MarkFlagRequired("certificate")

	UploadCustomCert.Flags().StringVar(&PrivateKey, "private-key", "", "The zone's private key. Example -----BEGIN RSA PRIVATE KEY-----...-----END RSA PRIVATE KEY-----")
	UploadCustomCert.MarkFlagRequired("private-key")

	UploadCustomCert.Flags().StringVar(&BundleMethod, "bundle-method", "", "A ubiquitous bundle is a bundle that has a higher probability of being verified everywhere, even by clients using outdated or unusual trust stores. default value: ubiquitous; valid values: ubiquitous, optimal, force  required = false")

	var UpdateCustomCert = &cobra.Command{
		Use:   "update-custom-cert",
		Short: "Upload a custom SSL cert",
		Long:  `Update a new SSL certificate for a zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateSSL")
		},
	}

	UpdateCustomCert.Flags().StringVar(&ZoneId, "zone-id", "", "The zone-id associated with the custom cert")
	UpdateCustomCert.MarkFlagRequired("zone-id")

	UpdateCustomCert.Flags().StringVar(&CertificateId, "certificate-id", "", "The certificate id associated with the custom cert")
	UpdateCustomCert.MarkFlagRequired("certificate-id")

	UpdateCustomCert.Flags().StringVar(&Certificate, "certificate", "", "The zone's SSL certificate or certificate and the intermediate(s)")
	UpdateCustomCert.MarkFlagRequired("certificate")

	UpdateCustomCert.Flags().StringVar(&PrivateKey, "private-key", "", "The zone's private key. Example -----BEGIN RSA PRIVATE KEY-----...-----END RSA PRIVATE KEY-----")
	UpdateCustomCert.MarkFlagRequired("private-key")

	UpdateCustomCert.Flags().StringVar(&BundleMethod, "bundle-method", "", "A ubiquitous bundle is a bundle that has a higher probability of being verified everywhere, even by clients using outdated or unusual trust stores. default value: ubiquitous; valid values: ubiquitous, optimal, force  required = false")

	var Purge = &cobra.Command{
		Use:   "purge",
		Short: "Purge specific items",
		Long:  `Purge specific items in the cache`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "Purge")
		},
	}

	Purge.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID that will be purged.")
	Purge.MarkFlagRequired("zone-id")

	Purge.Flags().StringVar(&Files, "files", "", "The files that will be purged.")

	Purge.Flags().StringVar(&Tags, "tags", "", "The tags that will be purged.")

	Purge.Flags().StringVar(&Hosts, "hosts", "", "The hosts that will be purged.")

	var CreateUserAgentRule = &cobra.Command{
		Use:   "create-user-agent-rule",
		Short: "Creates a new User-Agent rule",
		Long:  `Creates a new User-Agent rule for a specific zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateUserAgentRule")
		},
	}

	CreateUserAgentRule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the user-agent rule.")
	CreateUserAgentRule.MarkFlagRequired("zone-id")

	CreateUserAgentRule.Flags().StringVar(&Mode, "mode", "", "The type of action to perform. max length: 12 valid values: block, challenge, js_challenge")
	CreateUserAgentRule.MarkFlagRequired("mode")

	CreateUserAgentRule.Flags().StringVar(&Configuration, "configuration", "", "Target/Value pair to use for this rule. The value is the exact UserAgent to match {\"target\": \"ua\",  \"value\": \"Mozilla/5.0\"}")
	CreateUserAgentRule.MarkFlagRequired("configuration")

	CreateUserAgentRule.Flags().StringVar(&Description, "description", "", "Some useful information about this rule to help identify the purpose of it.")

	CreateUserAgentRule.Flags().BoolVar(&Paused, "paused", false, "Whether this UA rule is currently paused")

	var UpdateUserAgentRule = &cobra.Command{
		Use:   "update-user-agent-rule",
		Short: "Updates a new User-Agent rule",
		Long:  `Updates a new User-Agent rule for a specific zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateUserAgentRule")
		},
	}

	UpdateUserAgentRule.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the user-agent rule.")
	UpdateUserAgentRule.MarkFlagRequired("zone-id")

	UpdateUserAgentRule.Flags().StringVar(&UserAgentId, "user-agent-id", "", "The user agent ID associated with the user-agent rule.")
	UpdateUserAgentRule.MarkFlagRequired("user-agent-id")

	UpdateUserAgentRule.Flags().StringVar(&Mode, "mode", "", "The type of action to perform. max length: 12 valid values: block, challenge, js_challenge")
	UpdateUserAgentRule.MarkFlagRequired("mode")

	UpdateUserAgentRule.Flags().StringVar(&Configuration, "configuration", "", "Target/Value pair to use for this rule. The value is the exact UserAgent to match {\"target\": \"ua\",  \"value\": \"Mozilla/5.0\"}")
	UpdateUserAgentRule.MarkFlagRequired("configuration")

	UpdateUserAgentRule.Flags().StringVar(&Description, "description", "", "Some useful information about this rule to help identify the purpose of it.")

	UpdateUserAgentRule.Flags().BoolVar(&Paused, "paused", false, "Whether this UA rule is currently paused")

	var UpdateCustomHostname = &cobra.Command{
		Use:   "update-custom-hostname",
		Short: "Create a custom hostname for an associated zone.",
		Long:  `Modify SSL configuration for a custom hostname. When sent with SSL config that matches existing config, used to indicate that hostname should pass domain control validation (DCV). Can also be used to change validation type, e.g., from 'http' to 'email'.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateCustomHostnameSSL")
		},
	}

	UpdateCustomHostname.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the custom hostname")
	UpdateCustomHostname.MarkFlagRequired("zone-id")

	UpdateCustomHostname.Flags().StringVar(&CustomHostnameId, "custom-hostname-id", "", "The custom hostname id associated with the custom hostname being modified")
	UpdateCustomHostname.MarkFlagRequired("custom-hostname-id")

	UpdateCustomHostname.Flags().StringVar(&Method, "method", "", "The SSL Verification method. valid values: http, email, cname.")
	UpdateCustomHostname.MarkFlagRequired("method")

	UpdateCustomHostname.Flags().StringVar(&Type, "type", "", "The type of SSL certificate valid values: dv only")
	UpdateCustomHostname.MarkFlagRequired("type")

	var ListCustomHostnames = &cobra.Command{
		Use:   "list-custom-hostnames",
		Short: "List custom hostnames",
		Long:  `List Custom Hostnames fetches custom hostnames for the given zone, by applying filter.Hostname if not empty and scoping the result to page'th 50 items.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CustomHostnames")
		},
	}

	ListCustomHostnames.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the custom hostname")
	ListCustomHostnames.MarkFlagRequired("zone-id")

	ListCustomHostnames.Flags().StringVar(&Hostname, "hostname", "", "The custom hostname that you wish to filter by")

	ListCustomHostnames.Flags().IntVar(&Page, "page", 0, "API supports pagination. Up to 50 results per page. Default is page 1")

	var ReprioritizeCerts = &cobra.Command{
		Use:   "reprioritize-certs",
		Short: "Reprioritize SSL certs",
		Long:  `If a zone has multiple SSL certificates, you can set the order in which they should be used during a request. Higher priority numbers will be tried first.`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ReprioritizeSSL")
		},
	}

	ReprioritizeCerts.Flags().StringVar(&ZoneId, "zone-id", "", "The zone ID associated with the certs")
	ReprioritizeCerts.MarkFlagRequired("zone-id")

	ReprioritizeCerts.Flags().StringVar(&PriorityList, "priority-list", "", "Array of ordered certificates. [{\"id\": \"5a7805061c76ada191ed06f989cc3dac\",\"priority\": 2},{\"id\": \"9a7806061c88ada191ed06f989cc3dac\",\"priority\": 1}]")
	ReprioritizeCerts.MarkFlagRequired("priority-list")

	var UpdateLoadbalancerMonitor = &cobra.Command{
		Use:   "update-loadbalancer-monitor",
		Short: "Update a configured monitor",
		Long:  `Update an existing monitor`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ModifyLoadBalancerMonitor")
		},
	}

	UpdateLoadbalancerMonitor.Flags().StringVar(&MonitorId, "monitor-id", "", "The monitor id associated with the existing loadbalancer monitor")
	UpdateLoadbalancerMonitor.MarkFlagRequired("monitor-id")

	UpdateLoadbalancerMonitor.Flags().StringVar(&ExpectedCodes, "expected-codes", "", "The expected http response code in the healthcheck")
	UpdateLoadbalancerMonitor.MarkFlagRequired("expected-codes")

	UpdateLoadbalancerMonitor.Flags().StringVar(&Method, "method", "", "The HTTP method to use for the health check. default value: GET")

	UpdateLoadbalancerMonitor.Flags().StringVar(&Header, "header", "", "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Example: {\"Host\": [\"example.com\"],\"X-App-ID\": [\"abc123\"]}")

	UpdateLoadbalancerMonitor.Flags().IntVar(&Timeout, "timeout", 0, "The timeout (in seconds) before marking the health check as failed. default value: 5")
	UpdateLoadbalancerMonitor.MarkFlagRequired("timeout")

	UpdateLoadbalancerMonitor.Flags().StringVar(&Path, "path", "", "The endpoint path to health check against. default value: /")
	UpdateLoadbalancerMonitor.MarkFlagRequired("path")

	UpdateLoadbalancerMonitor.Flags().IntVar(&Interval, "interval", 0, "The interval between each health check. Shorter intervals may improve failover time, but will increase load. default value 60")
	UpdateLoadbalancerMonitor.MarkFlagRequired("interval")

	UpdateLoadbalancerMonitor.Flags().IntVar(&Retries, "retries", 0, "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. default value 2")
	UpdateLoadbalancerMonitor.MarkFlagRequired("retries")

	UpdateLoadbalancerMonitor.Flags().StringVar(&ExpectedBody, "expected-body", "", "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.")

	UpdateLoadbalancerMonitor.Flags().StringVar(&Type, "type", "", "The protocol to use for the healthcheck. Currently supported protocols are 'HTTP' and 'HTTPS'. default value: http")
	UpdateLoadbalancerMonitor.MarkFlagRequired("type")

	UpdateLoadbalancerMonitor.Flags().StringVar(&Description, "description", "", "Object description")
	UpdateLoadbalancerMonitor.MarkFlagRequired("description")

	var UpdateLoadbalancer = &cobra.Command{
		Use:   "update-loadbalancer",
		Short: "Update an existing loadbalancer",
		Long:  `Update an existing loadbalancer`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ModifyLoadBalancer")
		},
	}

	UpdateLoadbalancer.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the loadbalancer")
	UpdateLoadbalancer.MarkFlagRequired("zone-id")

	UpdateLoadbalancer.Flags().StringVar(&LoadbalancerId, "loadbalancer-id", "", "The loadbalancer id associated with the loadbalancer being modified")
	UpdateLoadbalancer.MarkFlagRequired("loadbalancer-id")

	UpdateLoadbalancer.Flags().StringVar(&Name, "name", "", "The DNS hostname to associate with your Load Balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.")
	UpdateLoadbalancer.MarkFlagRequired("name")

	UpdateLoadbalancer.Flags().StringVar(&FallbackPool, "fallback-pool", "", "The pool ID to use when all other pools are detected as unhealthy. max length: 32")
	UpdateLoadbalancer.MarkFlagRequired("fallback-pool")

	UpdateLoadbalancer.Flags().StringVar(&DefaultPools, "default-pools", "", "A comma separated list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.")
	UpdateLoadbalancer.MarkFlagRequired("default-pools")

	UpdateLoadbalancer.Flags().BoolVar(&Proxied, "proxied", false, "Whether the hostname should be gray clouded (false) or orange clouded (true). default value: false")

	UpdateLoadbalancer.Flags().IntVar(&Ttl, "ttl", 0, "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.")

	var CreateWorkerRoute = &cobra.Command{
		Use:   "create-worker-route",
		Short: "Create Worker Router",
		Long:  `Create a new Edge Worker Route Filter`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "CreateWorkerRoute")
		},
	}

	CreateWorkerRoute.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker route")
	CreateWorkerRoute.MarkFlagRequired("zone-id")

	CreateWorkerRoute.Flags().StringVar(&Pattern, "pattern", "", "The url pattern for the route, example: example.net/*")
	CreateWorkerRoute.MarkFlagRequired("pattern")

	CreateWorkerRoute.Flags().BoolVar(&Disable, "disable", false, "Set this flag to disable the worker on a given route")

	var UpdateWorkerRoute = &cobra.Command{
		Use:   "update-worker-route",
		Short: "Update Worker Route Filter",
		Long:  `Create a new Edge Worker Route Filter`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UpdateWorkerRoute")
		},
	}

	UpdateWorkerRoute.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker route")
	UpdateWorkerRoute.MarkFlagRequired("zone-id")

	UpdateWorkerRoute.Flags().StringVar(&RouteId, "route-id", "", "The routeId associated with the worker route")
	UpdateWorkerRoute.MarkFlagRequired("route-id")

	UpdateWorkerRoute.Flags().StringVar(&Pattern, "pattern", "", "The url pattern for the route, example: example.net/*")
	UpdateWorkerRoute.MarkFlagRequired("pattern")

	UpdateWorkerRoute.Flags().BoolVar(&Disable, "disable", false, "Set this flag to disable the worker on a given route")

	var ListWorkerRoutes = &cobra.Command{
		Use:   "list-worker-routes",
		Short: "List Worker Route Filters",
		Long:  `List all a new Edge Worker Route Filter`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListWorkerRoutes")
		},
	}

	ListWorkerRoutes.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker route filters")
	ListWorkerRoutes.MarkFlagRequired("zone-id")

	var UploadWorker = &cobra.Command{
		Use:   "upload-worker",
		Short: "Upload a Cloudflare Worker for your zone",
		Long:  `Upload a cloudflare worker for your zone`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UploadWorker")
		},
	}

	UploadWorker.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker")
	UploadWorker.MarkFlagRequired("zone-id")

	UploadWorker.Flags().StringVar(&Script, "script", "", "The raw worker code to run at the edge. @ to read code from a file. - to read from stdin")
	UploadWorker.MarkFlagRequired("script")

	var UploadOrganizationWorker = &cobra.Command{
		Use:   "upload-organization-worker",
		Short: "Upload a Cloudflare Worker associated with your Organization",
		Long:  `Upload an organization's cloudflare worker`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "UploadOrganizationWorker")
		},
	}

	UploadOrganizationWorker.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker")
	UploadOrganizationWorker.MarkFlagRequired("zone-id")

	UploadOrganizationWorker.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the worker")
	UploadOrganizationWorker.MarkFlagRequired("organization-id")

	UploadOrganizationWorker.Flags().StringVar(&Name, "name", "", "The worker's name")
	UploadOrganizationWorker.MarkFlagRequired("name")

	UploadOrganizationWorker.Flags().StringVar(&Script, "script", "", "The raw worker code to run at the edge. @ to read code from a file. - to read from stdin")
	UploadOrganizationWorker.MarkFlagRequired("script")

	var DeleteWorker = &cobra.Command{
		Use:   "delete-worker",
		Short: "Delete Worker",
		Long:  `Delete a cloudflare worker`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteWorker")
		},
	}

	DeleteWorker.Flags().StringVar(&ZoneId, "zone-id", "", "The zoneID associated with the worker")
	DeleteWorker.MarkFlagRequired("zone-id")

	var DeleteOrganizationWorker = &cobra.Command{
		Use:   "delete-organization-worker",
		Short: "Delete Organization Worker",
		Long:  `Delete an organization's cloudflare worker`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DeleteOrganizationWorker")
		},
	}

	DeleteOrganizationWorker.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the worker")
	DeleteOrganizationWorker.MarkFlagRequired("organization-id")

	DeleteOrganizationWorker.Flags().StringVar(&Name, "name", "", "The worker's name")
	DeleteOrganizationWorker.MarkFlagRequired("name")

	var ListWorkerScripts = &cobra.Command{
		Use:   "list-worker-scripts",
		Short: "List your worker scripts",
		Long:  `List your cloudflare worker scripts`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListWorkerScripts")
		},
	}

	ListWorkerScripts.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the workers")
	ListWorkerScripts.MarkFlagRequired("organization-id")

	var DownloadWorker = &cobra.Command{
		Use:   "download-worker",
		Short: "Download your worker script",
		Long:  `Download your cloudflare worker script`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DownloadWorker")
		},
	}

	DownloadWorker.Flags().StringVar(&ZoneId, "zone-id", "", "The zone id associated with the workers")
	DownloadWorker.MarkFlagRequired("zone-id")

	var DownloadOrganizationWorker = &cobra.Command{
		Use:   "download-organization-worker",
		Short: "Download your worker script",
		Long:  `Download a worker associated with your organization`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "DownloadOrganizationWorker")
		},
	}

	DownloadOrganizationWorker.Flags().StringVar(&OrganizationId, "organization-id", "", "The organization id associated with the workers")
	DownloadOrganizationWorker.MarkFlagRequired("organization-id")

	DownloadOrganizationWorker.Flags().StringVar(&Name, "name", "", "The worker's name")
	DownloadOrganizationWorker.MarkFlagRequired("name")

	var Zone = &cobra.Command{
		Use:   "zone",
		Short: "Interact with cloudflare zones",
		Long:  `  A Zone is a domain name along with its subdomains and other identities. Using the zone command to interact with cloudflare zones`,
	}
	Zone.AddCommand(ActivationCheck)
	Zone.AddCommand(AnalyticsByColo)
	Zone.AddCommand(AnalyticsDashboard)
	Zone.AddCommand(CreateCustomHostname)
	Zone.AddCommand(CreateZone)
	Zone.AddCommand(DeleteCustomHostname)
	Zone.AddCommand(DeleteZone)
	Zone.AddCommand(DeleteZoneAccessRule)
	Zone.AddCommand(DescribeZone)
	Zone.AddCommand(EditZonePaused)
	Zone.AddCommand(EditZoneVanityNs)
	Zone.AddCommand(DescribeCustomHostname)
	Zone.AddCommand(DescribeCustomHostnameByName)
	Zone.AddCommand(GetIdByName)
	Zone.AddCommand(GetZoneSettings)
	Zone.AddCommand(ListAvailableRatePlans)
	Zone.AddCommand(ListCustomHostnames)
	Zone.AddCommand(ListZoneAccessRules)
	Zone.AddCommand(ListZones)
	Zone.AddCommand(SetPaused)
	Zone.AddCommand(SetVanityNs)
	Zone.AddCommand(UpdateCustomHostname)
	Zone.AddCommand(UpdateZoneAccessRule)
	Zone.AddCommand(UpdateZoneSettings)

	RootCmd.AddCommand(Zone)

	var Worker = &cobra.Command{
		Use:   "worker",
		Short: "Interact with cloudflare workers api",
		Long: `  Manage your edge workers deployed to your account.
`,
	}
	Worker.AddCommand(CreateWorkerRoute)
	Worker.AddCommand(DeleteOrganizationWorker)
	Worker.AddCommand(DeleteWorker)
	Worker.AddCommand(DownloadOrganizationWorker)
	Worker.AddCommand(DownloadWorker)
	Worker.AddCommand(ListWorkerRoutes)
	Worker.AddCommand(ListWorkerScripts)
	Worker.AddCommand(UpdateWorkerRoute)
	Worker.AddCommand(UploadOrganizationWorker)
	Worker.AddCommand(UploadWorker)

	RootCmd.AddCommand(Worker)

	var Dns = &cobra.Command{
		Use:   "dns",
		Short: "Create, read, update, and delete your dns records.",
		Long:  `  Create, read, update, and delete cloudflare dns and virtual dns.`,
	}
	Dns.AddCommand(CreateDnsRecord)
	Dns.AddCommand(CreateVirtualDns)
	Dns.AddCommand(DeleteDnsRecord)
	Dns.AddCommand(DeleteVirtualDns)
	Dns.AddCommand(DescribeVirtualDns)
	Dns.AddCommand(EditDnsRecord)
	Dns.AddCommand(ListDnsRecords)
	Dns.AddCommand(ListVirtualDns)
	Dns.AddCommand(ShowDnsRecord)
	Dns.AddCommand(UpdateVirtualDns)

	RootCmd.AddCommand(Dns)

	var User = &cobra.Command{
		Use:   "user",
		Short: "Interacting with your Cloudflare account",
		Long:  `  The User subcommand will help you interact with your Cloudflare account.`,
	}
	User.AddCommand(BillingProfile)
	User.AddCommand(CreateUserAccessRule)
	User.AddCommand(DeleteUserAccessRule)
	User.AddCommand(Details)
	User.AddCommand(EditUser)
	User.AddCommand(ListUserAccessRules)
	User.AddCommand(UpdateUserAccessRule)

	RootCmd.AddCommand(User)

	var Ssl = &cobra.Command{
		Use:   "ssl",
		Short: "Control and insight in to your zone's SSL stack ",
		Long:  `  Fine grained control and deep insights in to your zone's SSL stack, as well as the ability to order new certs.`,
	}
	Ssl.AddCommand(CreateOriginCert)
	Ssl.AddCommand(DescribeOriginCert)
	Ssl.AddCommand(DescribeZoneOriginCert)
	Ssl.AddCommand(DeleteCustomCert)
	Ssl.AddCommand(ListCustomCerts)
	Ssl.AddCommand(ListOriginCerts)
	Ssl.AddCommand(ListZoneSslSettings)
	Ssl.AddCommand(ReprioritizeCerts)
	Ssl.AddCommand(RevokeOriginCert)
	Ssl.AddCommand(UpdateCustomCert)
	Ssl.AddCommand(UploadCustomCert)

	RootCmd.AddCommand(Ssl)

	var Pagerule = &cobra.Command{
		Use:   "pagerule",
		Short: "Change how cloudflare works on a URL or subdomain basis.",
		Long:  `  Gives you the ability to control how Cloudflare works on a URL or subdomain basis. Page Rules allow you to customize Cloudflare's functionality to match your domain's unique needs.`,
	}
	Pagerule.AddCommand(CreatePagerule)
	Pagerule.AddCommand(ListPagerules)
	Pagerule.AddCommand(DeletePagerule)
	Pagerule.AddCommand(DescribePagerule)
	Pagerule.AddCommand(UpdatePagerule)

	RootCmd.AddCommand(Pagerule)

	var Cache = &cobra.Command{
		Use:   "cache",
		Short: "Commands for interacting with caching and railgun APIs",
		Long:  `  Commands for the management and description of cache technologies.`,
	}
	Cache.AddCommand(ConnectZoneRailgun)
	Cache.AddCommand(CreateRailgun)
	Cache.AddCommand(DeleteRailgun)
	Cache.AddCommand(DescribeRailgun)
	Cache.AddCommand(DescribeZoneRailgun)
	Cache.AddCommand(DisableRailgun)
	Cache.AddCommand(DisconnectRailgun)
	Cache.AddCommand(EnableRailgun)
	Cache.AddCommand(GetRailgunZones)
	Cache.AddCommand(ListRailguns)
	Cache.AddCommand(ListZoneRailguns)
	Cache.AddCommand(Purge)
	Cache.AddCommand(PurgeEverything)
	Cache.AddCommand(TestRailgunConnection)

	RootCmd.AddCommand(Cache)

	var Firewall = &cobra.Command{
		Use:   "firewall",
		Short: "Commands to interact with products that block access to your site",
		Long:  `  The firewall subcommand is used to configure the WAF, user-agent rules, and zone-lockdowns.`,
	}
	Firewall.AddCommand(CreateUserAgentRule)
	Firewall.AddCommand(CreateZoneLockdown)
	Firewall.AddCommand(DeleteUserAgentRule)
	Firewall.AddCommand(DeleteZoneLockdown)
	Firewall.AddCommand(DescribeZoneLockdown)
	Firewall.AddCommand(ListUserAgentRules)
	Firewall.AddCommand(ListWafPackages)
	Firewall.AddCommand(ListWafRules)
	Firewall.AddCommand(ListZoneLockdowns)
	Firewall.AddCommand(UpdateUserAgentRule)
	Firewall.AddCommand(UpdateZoneLockdown)

	RootCmd.AddCommand(Firewall)

	var Organization = &cobra.Command{
		Use:   "organization",
		Short: "Interact with the organizations you own and have access to",
		Long:  `  Organizztions show all information about the organizations and provides the ability to modify access and display information about your specific organizations.`,
	}
	Organization.AddCommand(CreateOrganizationAccessRule)
	Organization.AddCommand(DeleteOrganizationAccessRule)
	Organization.AddCommand(DescribeOrganization)
	Organization.AddCommand(GetOrganizationInvites)
	Organization.AddCommand(GetOrganizationMembers)
	Organization.AddCommand(GetOrganizationRoles)
	Organization.AddCommand(ListOrganizationAccessRules)
	Organization.AddCommand(ListOrganizations)
	Organization.AddCommand(UpdateOrganizationAccessRule)

	RootCmd.AddCommand(Organization)

	var Ratelimit = &cobra.Command{
		Use:   "ratelimit",
		Short: "Configure, create, and view your zone's ratelimits.",
		Long:  `  Cloudflare ratelimits that ensure that the requests to your site are done at a rate you accept.`,
	}
	Ratelimit.AddCommand(CreateRatelimit)
	Ratelimit.AddCommand(DeleteRatelimit)
	Ratelimit.AddCommand(DescribeRatelimit)
	Ratelimit.AddCommand(ListRatelimits)
	Ratelimit.AddCommand(UpdateRatelimit)

	RootCmd.AddCommand(Ratelimit)

	var Loadbalancer = &cobra.Command{
		Use:   "loadbalancer",
		Short: "Manage, create, and describe your loadbalancers, pools, and monitor",
		Long:  `  The Cloudflare loadbalancer product is an enterprise ready load balancing platform.`,
	}
	Loadbalancer.AddCommand(CreateLoadbalancer)
	Loadbalancer.AddCommand(CreateLoadbalancerPool)
	Loadbalancer.AddCommand(CreateLoadbalancerMonitor)
	Loadbalancer.AddCommand(DeleteLoadbalancer)
	Loadbalancer.AddCommand(DeleteLoadbalancerMonitor)
	Loadbalancer.AddCommand(DeleteLoadbalancerPool)
	Loadbalancer.AddCommand(DescribeLoadbalancer)
	Loadbalancer.AddCommand(DescribeLoadbalancerPool)
	Loadbalancer.AddCommand(DescribeLoadbalancerMonitor)
	Loadbalancer.AddCommand(ListLoadbalancerMonitors)
	Loadbalancer.AddCommand(ListLoadbalancerPools)
	Loadbalancer.AddCommand(ListLoadbalancers)
	Loadbalancer.AddCommand(UpdateLoadbalancer)
	Loadbalancer.AddCommand(UpdateLoadbalancerMonitor)
	Loadbalancer.AddCommand(UpdateLoadbalancerPool)

	RootCmd.AddCommand(Loadbalancer)

}

func Run(cmd *cobra.Command, args []string, name string, api *cloudflare.API) (resp interface{}, err error) {
	switch name {

	case "ListZones":
		resp, err = ListZones(api, ZoneNameFilter)
	case "ListDnsRecords":
		resp, err = ListDnsRecords(api, ZoneId, Type, Name, Content)
	case "CreateDnsRecord":
		resp, err = CreateDnsRecord(api, ZoneId, Type, Name, Content, Ttl, NotProxied, Priority)
	case "DeleteDnsRecord":
		resp, err = DeleteDnsRecord(api, ZoneId, RecordId)
	case "DeleteZone":
		resp, err = DeleteZone(api, ZoneId)
	case "CreateZone":
		resp, err = CreateZone(api, Name, Jumpstart, OrganizationId)
	case "DNSRecord":
		resp, err = DNSRecord(api, ZoneId, RecordId)
	case "ListAllRateLimits":
		resp, err = ListAllRateLimits(api, ZoneId)
	case "ListLoadBalancers":
		resp, err = ListLoadBalancers(api, ZoneId)
	case "ListOrganizations":
		resp, err = ListOrganizations(api)
	case "ListPageRules":
		resp, err = ListPageRules(api, ZoneId)
	case "ListCustomCerts":
		resp, err = ListCustomCerts(api, ZoneId)
	case "ListUserAgentRules":
		resp, err = ListUserAgentRules(api, ZoneId, Page)
	case "ListWAFPackages":
		resp, err = ListWAFPackages(api, ZoneId)
	case "ListWAFRules":
		resp, err = ListWAFRules(api, ZoneId, PackageId)
	case "ListZoneLockdowns":
		resp, err = ListZoneLockdowns(api, ZoneId, Page)
	case "ZoneLockdown":
		resp, err = ZoneLockdown(api, ZoneId, LockdownId)
	case "EditZonePaused":
		resp, err = EditZonePaused(api, ZoneId, Paused)
	case "EditZoneVanityNS":
		resp, err = EditZoneVanityNS(api, ZoneId, VanityNs)
	case "ZoneSetVanityNS":
		resp, err = ZoneSetVanityNS(api, ZoneId, VanityNs)
	case "EditDNSRecord":
		resp, err = EditDNSRecord(api, Proxied, ZoneId, RecordId, Type, Name, Content, Ttl)
	case "ListLoadBalancerMonitors":
		resp, err = ListLoadBalancerMonitors(api)
	case "ListLoadBalancerPools":
		resp, err = ListLoadBalancerPools(api)
	case "ListOrganizationAccessRules":
		resp, err = ListOrganizationAccessRules(api, OrganizationId, Notes, Mode, Page)
	case "ListRailguns":
		resp, err = ListRailguns(api)
	case "ListZoneRailguns":
		resp, err = ListZoneRailguns(api)
	case "ListUserAccessRules":
		resp, err = ListUserAccessRules(api, Notes, Mode, Page)
	case "ListVirtualDns":
		resp, err = ListVirtualDns(api)
	case "AvailableZoneRatePlans":
		resp, err = AvailableZoneRatePlans(api, ZoneId)
	case "ConnectZoneRailgun":
		resp, err = ConnectZoneRailgun(api, ZoneId, RailgunId)
	case "CreateCustomHostname":
		resp, err = CreateCustomHostname(api, ZoneId, Hostname, Method, Type)
	case "CreateLoadBalancerMonitor":
		resp, err = CreateLoadBalancerMonitor(api, ExpectedCodes, Method, Header, Timeout, Path, Interval, Retries, ExpectedBody, Type, Description)
	case "CreateLoadBalancer":
		resp, err = CreateLoadBalancer(api, ZoneId, Name, FallbackPool, DefaultPools, Proxied, Ttl)
	case "PurgeEverything":
		resp, err = PurgeEverything(api, ZoneId)
	case "ActivationCheck":
		resp, err = ActivationCheck(api, ZoneId)
	case "ZoneDetails":
		resp, err = ZoneDetails(api, ZoneId)
	case "GetIDByName":
		resp, err = GetIDByName(api, ZoneName)
	case "ZoneSSLSettings":
		resp, err = ZoneSSLSettings(api, ZoneId)
	case "ZoneSettings":
		resp, err = ZoneSettings(api, ZoneId)
	case "UserDetails":
		resp, err = UserDetails(api)
	case "UserBillingProfile":
		resp, err = UserBillingProfile(api)
	case "VirtualDNS":
		resp, err = VirtualDNS(api, VirtualDnsId)
	case "DeleteVirtualDNS":
		resp, err = DeleteVirtualDNS(api, VirtualDnsId)
	case "PageRule":
		resp, err = PageRule(api, ZoneId, PageruleId)
	case "LoadBalancerDetails":
		resp, err = LoadBalancerDetails(api, ZoneId, LoadbalancerId)
	case "LoadBalancerMonitorDetails":
		resp, err = LoadBalancerMonitorDetails(api, MonitorId)
	case "LoadBalancerPoolDetails":
		resp, err = LoadBalancerPoolDetails(api, PoolId)
	case "OrganizationDetails":
		resp, err = OrganizationDetails(api, OrganizationId)
	case "OrganizationInvites":
		resp, err = OrganizationInvites(api, OrganizationId)
	case "OrganizationMembers":
		resp, err = OrganizationMembers(api, OrganizationId)
	case "OrganizationRoles":
		resp, err = OrganizationRoles(api, OrganizationId)
	case "OriginCertificates":
		resp, err = OriginCertificates(api, ZoneId)
	case "OriginCertificate":
		resp, err = OriginCertificate(api, CertificateId)
	case "SSLDetails":
		resp, err = SSLDetails(api, CertificateId, ZoneId)
	case "RailgunDetails":
		resp, err = RailgunDetails(api, RailgunId)
	case "RailgunZones":
		resp, err = RailgunZones(api, RailgunId)
	case "RateLimit":
		resp, err = RateLimit(api, ZoneId, RatelimitId)
	case "RevokeOriginCertificate":
		resp, err = RevokeOriginCertificate(api, CertificateId)
	case "TestRailgunConnection":
		resp, err = TestRailgunConnection(api, ZoneId, RailgunId)
	case "ZoneRailgunDetails":
		resp, err = ZoneRailgunDetails(api, ZoneId, RailgunId)
	case "CustomHostname":
		resp, err = CustomHostname(api, ZoneId, CustomHostnameId)
	case "CustomHostnameIDByName":
		resp, err = CustomHostnameIDByName(api, ZoneId, Name)
	case "ZoneSetPaused":
		resp, err = ZoneSetPaused(api, ZoneId, Paused)
	case "DeletePageRule":
		resp, err = DeletePageRule(api, ZoneId, PageruleId)
	case "DeleteRailgun":
		resp, err = DeleteRailgun(api, RailgunId)
	case "DisableRailgun":
		resp, err = DisableRailgun(api, RailgunId)
	case "DisconnectZoneRailgun":
		resp, err = DisconnectZoneRailgun(api, RailgunId, ZoneId)
	case "EnableRailgun":
		resp, err = EnableRailgun(api, RailgunId)
	case "DeleteRateLimit":
		resp, err = DeleteRateLimit(api, ZoneId, RatelimitId)
	case "DeleteSSL":
		resp, err = DeleteSSL(api, ZoneId, CertificateId)
	case "DeleteCustomHostname":
		resp, err = DeleteCustomHostname(api, ZoneId, CustomHostnameId)
	case "DeleteLoadBalancer":
		resp, err = DeleteLoadBalancer(api, ZoneId, LoadbalancerId)
	case "DeleteLoadBalancerMonitor":
		resp, err = DeleteLoadBalancerMonitor(api, MonitorId)
	case "DeleteLoadBalancerPool":
		resp, err = DeleteLoadBalancerPool(api, PoolId)
	case "DeleteOrganizationAccessRule":
		resp, err = DeleteOrganizationAccessRule(api, OrganizationId, AccessRuleId)
	case "CreateRailgun":
		resp, err = CreateRailgun(api, Name)
	case "DeleteUserAccessRule":
		resp, err = DeleteUserAccessRule(api, AccessRuleId)
	case "DeleteUserAgentRule":
		resp, err = DeleteUserAgentRule(api, UserAgentId, ZoneId)
	case "DeleteZoneAccessRule":
		resp, err = DeleteZoneAccessRule(api, AccessRuleId, ZoneId)
	case "DeleteZoneLockdown":
		resp, err = DeleteZoneLockdown(api, ZoneId, LockdownId)
	case "ZoneAnalyticsByColocation":
		resp, err = ZoneAnalyticsByColocation(api, ZoneId, Since, Until, Continuous)
	case "ZoneAnalyticsDashboard":
		resp, err = ZoneAnalyticsDashboard(api, ZoneId, Since, Until, Continuous)
	case "UpdateUser":
		resp, err = UpdateUser(api, FirstName, LastName, Telephone, Country, Zipcode)
	case "CreateLoadBalancerPool":
		resp, err = CreateLoadBalancerPool(api, Name, Origins, Description, Disabled, MinimumOrigins, Monitor, NotificationEmail)
	case "ModifyLoadBalancerPool":
		resp, err = ModifyLoadBalancerPool(api, PoolId, Name, Origins, Description, Disabled, MinimumOrigins, Monitor, NotificationEmail)
	case "UpdateZoneSettings":
		resp, err = UpdateZoneSettings(api, ZoneId, ZoneSettingsObject)
	case "UpdateZoneLockdown":
		resp, err = UpdateZoneLockdown(api, ZoneId, LockdownId, Configuration, Urls, Paused, Description)
	case "CreateZoneLockdown":
		resp, err = CreateZoneLockdown(api, ZoneId, Configuration, Urls, Paused, Description)
	case "CreateVirtualDNS":
		resp, err = CreateVirtualDNS(api, Name, OriginIps, MinimumCacheTtl, MaximumCacheTtl, DeprecateAnyRequest)
	case "UpdateVirtualDNS":
		resp, err = UpdateVirtualDNS(api, VirtualDnsId, OriginIps, MinimumCacheTtl, MaximumCacheTtl, DeprecateAnyRequest)
	case "CreatePageRule":
		resp, err = CreatePageRule(api, ZoneId, Targets, Actions, Priority, Status)
	case "ChangePageRule":
		resp, err = ChangePageRule(api, ZoneId, PageruleId, Targets, Actions, Priority, Status)
	case "CreateOrganizationAccessRule":
		resp, err = CreateOrganizationAccessRule(api, OrganizationId, Mode, Configuration, Notes)
	case "CreateOriginCertificate":
		resp, err = CreateOriginCertificate(api, Hostnames, RequestValidity, RequestType, Csr)
	case "CreateRateLimit":
		resp, err = CreateRateLimit(api, ZoneId, Match, Threshold, Period, Action, Enabled, Description, Bypass)
	case "UpdateRateLimit":
		resp, err = UpdateRateLimit(api, ZoneId, LimitId, Match, Threshold, Period, Action, Enabled, Description, Bypass)
	case "CreateUserAccessRule":
		resp, err = CreateUserAccessRule(api, Mode, Configuration, Notes)
	case "UpdateUserAccessRule":
		resp, err = UpdateUserAccessRule(api, AccessRuleId, Mode, Configuration, Notes)
	case "UpdateZoneAccessRule":
		resp, err = UpdateZoneAccessRule(api, ZoneId, AccessRuleId, Mode, Configuration, Notes)
	case "UpdateOrganizationAccessRule":
		resp, err = UpdateOrganizationAccessRule(api, OrganizationId, AccessRuleId, Mode, Configuration, Notes)
	case "ListZoneAccessRules":
		resp, err = ListZoneAccessRules(api, ZoneId, Notes, Mode, Page)
	case "CreateSSL":
		resp, err = CreateSSL(api, ZoneId, Certificate, PrivateKey, BundleMethod)
	case "UpdateSSL":
		resp, err = UpdateSSL(api, ZoneId, CertificateId, Certificate, PrivateKey, BundleMethod)
	case "Purge":
		resp, err = Purge(api, ZoneId, Files, Tags, Hosts)
	case "CreateUserAgentRule":
		resp, err = CreateUserAgentRule(api, ZoneId, Mode, Configuration, Description, Paused)
	case "UpdateUserAgentRule":
		resp, err = UpdateUserAgentRule(api, ZoneId, UserAgentId, Mode, Configuration, Description, Paused)
	case "UpdateCustomHostnameSSL":
		resp, err = UpdateCustomHostnameSSL(api, ZoneId, CustomHostnameId, Method, Type)
	case "CustomHostnames":
		resp, err = CustomHostnames(api, ZoneId, Hostname, Page)
	case "ReprioritizeSSL":
		resp, err = ReprioritizeSSL(api, ZoneId, PriorityList)
	case "ModifyLoadBalancerMonitor":
		resp, err = ModifyLoadBalancerMonitor(api, MonitorId, ExpectedCodes, Method, Header, Timeout, Path, Interval, Retries, ExpectedBody, Type, Description)
	case "ModifyLoadBalancer":
		resp, err = ModifyLoadBalancer(api, ZoneId, LoadbalancerId, Name, FallbackPool, DefaultPools, Proxied, Ttl)
	case "CreateWorkerRoute":
		resp, err = CreateWorkerRoute(api, ZoneId, Pattern, Disable)
	case "UpdateWorkerRoute":
		resp, err = UpdateWorkerRoute(api, ZoneId, RouteId, Pattern, Disable)
	case "ListWorkerRoutes":
		resp, err = ListWorkerRoutes(api, ZoneId)
	case "UploadWorker":
		resp, err = UploadWorker(api, ZoneId, Script)
	case "UploadOrganizationWorker":
		resp, err = UploadOrganizationWorker(api, ZoneId, OrganizationId, Name, Script)
	case "DeleteWorker":
		resp, err = DeleteWorker(api, ZoneId)
	case "DeleteOrganizationWorker":
		resp, err = DeleteOrganizationWorker(api, OrganizationId, Name)
	case "ListWorkerScripts":
		resp, err = ListWorkerScripts(api, OrganizationId)
	case "DownloadWorker":
		resp, err = DownloadWorker(api, ZoneId)
	case "DownloadOrganizationWorker":
		resp, err = DownloadOrganizationWorker(api, OrganizationId, Name)
	default:
		break
	}
	return
}
