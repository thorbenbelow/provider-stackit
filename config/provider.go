package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "stackit"
	modulePath     = "github.com/thorbenbelow/provider-stackit"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

type refSpec struct {
	target string
	field  string
}

type resourceSpec struct {
	group string
	kind  string
	refs  map[string]refSpec
}

func paramRef(target, field string) ujconfig.Reference {
	return ujconfig.Reference{
		TerraformName: target,
		Extractor:     fmt.Sprintf(`github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(%q,true)`, field),
	}
}

var resourceSpecs = map[string]resourceSpec{
	// 1. Core & IAM
	"stackit_resourcemanager_folder":  {group: "resourcemanager", kind: "Folder"},
	"stackit_resourcemanager_project": {group: "resourcemanager", kind: "Project", refs: map[string]refSpec{"parent_container_id": {"stackit_resourcemanager_folder", "container_id"}}},

	"stackit_authorization_folder_custom_role":              {group: "authorization", kind: "FolderCustomRole", refs: map[string]refSpec{"resource_id": {"stackit_resourcemanager_folder", "folder_id"}}},
	"stackit_authorization_folder_role_assignment":          {group: "authorization", kind: "FolderRoleAssignment", refs: map[string]refSpec{"resource_id": {"stackit_resourcemanager_folder", "folder_id"}}},
	"stackit_authorization_organization_custom_role":        {group: "authorization", kind: "OrganizationCustomRole"},
	"stackit_authorization_organization_role_assignment":    {group: "authorization", kind: "OrganizationRoleAssignment"},
	"stackit_authorization_project_custom_role":             {group: "authorization", kind: "ProjectCustomRole", refs: map[string]refSpec{"resource_id": {"stackit_resourcemanager_project", "project_id"}}},
	"stackit_authorization_project_role_assignment":         {group: "authorization", kind: "ProjectRoleAssignment", refs: map[string]refSpec{"resource_id": {"stackit_resourcemanager_project", "project_id"}}},
	"stackit_authorization_service_account_role_assignment": {group: "authorization", kind: "ServiceAccountRoleAssignment", refs: map[string]refSpec{"resource_id": {"stackit_service_account", "service_account_id"}}},

	"stackit_service_account":                             {group: "serviceaccount", kind: "ServiceAccount"},
	"stackit_service_account_federated_identity_provider": {group: "serviceaccount", kind: "FederatedIdentityProvider", refs: map[string]refSpec{"service_account_email": {"stackit_service_account", "email"}}},
	"stackit_service_account_key":                         {group: "serviceaccount", kind: "Key", refs: map[string]refSpec{"service_account_email": {"stackit_service_account", "email"}}},

	// 2. IaaS Compute, Storage & Networking
	"stackit_affinity_group":                  {group: "iaas", kind: "AffinityGroup"},
	"stackit_image":                           {group: "iaas", kind: "Image"},
	"stackit_key_pair":                        {group: "iaas", kind: "KeyPair"},
	"stackit_network":                         {group: "iaas", kind: "Network", refs: map[string]refSpec{"routing_table_id": {"stackit_routing_table", "routing_table_id"}, "vpc_id": {"stackit_vpc", "vpc_id"}}},
	"stackit_network_area":                    {group: "iaas", kind: "NetworkArea"},
	"stackit_network_area_region":             {group: "iaas", kind: "NetworkAreaRegion", refs: map[string]refSpec{"network_area_id": {"stackit_network_area", "network_area_id"}}},
	"stackit_network_area_route":              {group: "iaas", kind: "NetworkAreaRoute", refs: map[string]refSpec{"network_area_id": {"stackit_network_area", "network_area_id"}}},
	"stackit_network_interface":               {group: "iaas", kind: "NetworkInterface", refs: map[string]refSpec{"network_id": {"stackit_network", "network_id"}}},
	"stackit_public_ip":                       {group: "iaas", kind: "PublicIP", refs: map[string]refSpec{"network_interface_id": {"stackit_network_interface", "network_interface_id"}}},
	"stackit_public_ip_associate":             {group: "iaas", kind: "PublicIPAssociate", refs: map[string]refSpec{"public_ip_id": {"stackit_public_ip", "public_ip_id"}, "network_interface_id": {"stackit_network_interface", "network_interface_id"}}},
	"stackit_routing_table":                   {group: "iaas", kind: "RoutingTable", refs: map[string]refSpec{"network_area_id": {"stackit_network_area", "network_area_id"}}},
	"stackit_routing_table_route":             {group: "iaas", kind: "RoutingTableRoute", refs: map[string]refSpec{"network_area_id": {"stackit_network_area", "network_area_id"}, "routing_table_id": {"stackit_routing_table", "routing_table_id"}}},
	"stackit_security_group":                  {group: "iaas", kind: "SecurityGroup"},
	"stackit_security_group_rule":             {group: "iaas", kind: "SecurityGroupRule", refs: map[string]refSpec{"security_group_id": {"stackit_security_group", "security_group_id"}, "remote_security_group_id": {"stackit_security_group", "security_group_id"}}},
	"stackit_server":                          {group: "iaas", kind: "Server", refs: map[string]refSpec{"affinity_group": {"stackit_affinity_group", "affinity_group_id"}, "image_id": {"stackit_image", "image_id"}, "keypair_name": {"stackit_key_pair", "name"}}},
	"stackit_server_backup_enable":            {group: "iaas", kind: "ServerBackupEnable", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}}},
	"stackit_server_backup_schedule":          {group: "iaas", kind: "ServerBackupSchedule", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}}},
	"stackit_server_network_interface_attach": {group: "iaas", kind: "ServerNetworkInterfaceAttach", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}, "network_interface_id": {"stackit_network_interface", "network_interface_id"}}},
	"stackit_server_service_account_attach":   {group: "iaas", kind: "ServerServiceAccountAttach", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}, "service_account_email": {"stackit_service_account", "email"}}},
	"stackit_server_update_enable":            {group: "iaas", kind: "ServerUpdateEnable", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}}},
	"stackit_server_update_schedule":          {group: "iaas", kind: "ServerUpdateSchedule", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}}},
	"stackit_server_volume_attach":            {group: "iaas", kind: "ServerVolumeAttach", refs: map[string]refSpec{"server_id": {"stackit_server", "server_id"}, "volume_id": {"stackit_volume", "volume_id"}}},
	"stackit_volume":                          {group: "iaas", kind: "Volume"},
	"stackit_volume_automation":               {group: "iaas", kind: "VolumeAutomation"},

	"stackit_vpc":                            {group: "vpc", kind: "VPC"},
	"stackit_vpc_network_range":              {group: "vpc", kind: "NetworkRange", refs: map[string]refSpec{"vpc_id": {"stackit_vpc", "vpc_id"}}},
	"stackit_vpc_region":                     {group: "vpc", kind: "Region", refs: map[string]refSpec{"vpc_id": {"stackit_vpc", "vpc_id"}}},
	"stackit_vpc_routing_table":              {group: "vpc", kind: "RoutingTable", refs: map[string]refSpec{"vpc_id": {"stackit_vpc", "vpc_id"}}},
	"stackit_vpc_routing_table_static_route": {group: "vpc", kind: "RoutingTableStaticRoute", refs: map[string]refSpec{"vpc_id": {"stackit_vpc", "vpc_id"}, "routing_table_id": {"stackit_vpc_routing_table", "routing_table_id"}}},

	"stackit_vpn_connection": {group: "vpn", kind: "Connection", refs: map[string]refSpec{"gateway_id": {"stackit_vpn_gateway", "gateway_id"}}},
	"stackit_vpn_gateway":    {group: "vpn", kind: "Gateway"},

	// 3. Kubernetes, Load Balancing, DNS & CDN
	"stackit_ske_cluster":    {group: "ske", kind: "Cluster"},
	"stackit_ske_kubeconfig": {group: "ske", kind: "Kubeconfig", refs: map[string]refSpec{"cluster_name": {"stackit_ske_cluster", "name"}}},

	"stackit_edgecloud_instance":   {group: "edgecloud", kind: "Instance"},
	"stackit_edgecloud_kubeconfig": {group: "edgecloud", kind: "Kubeconfig", refs: map[string]refSpec{"instance_id": {"stackit_edgecloud_instance", "instance_id"}, "instance_name": {"stackit_edgecloud_instance", "display_name"}}},
	"stackit_edgecloud_token":      {group: "edgecloud", kind: "Token", refs: map[string]refSpec{"instance_id": {"stackit_edgecloud_instance", "instance_id"}, "instance_name": {"stackit_edgecloud_instance", "display_name"}}},

	"stackit_application_load_balancer": {group: "alb", kind: "ApplicationLoadBalancer"},
	"stackit_alb_certificate":           {group: "alb", kind: "Certificate"},
	"stackit_alb_waf_configuration":     {group: "alb", kind: "WAFConfiguration", refs: map[string]refSpec{"custom_rule_group_name": {"stackit_alb_waf_custom_rule_group", "name"}, "managed_rule_set_name": {"stackit_alb_waf_managed_rule_set", "name"}}},
	"stackit_alb_waf_custom_rule_group": {group: "alb", kind: "WAFCustomRuleGroup"},
	"stackit_alb_waf_managed_rule_set":  {group: "alb", kind: "WAFManagedRuleSet"},

	"stackit_loadbalancer":                          {group: "loadbalancer", kind: "LoadBalancer"},
	"stackit_loadbalancer_observability_credential": {group: "loadbalancer", kind: "ObservabilityCredential"},

	"stackit_dns_zone":       {group: "dns", kind: "Zone"},
	"stackit_dns_record_set": {group: "dns", kind: "RecordSet", refs: map[string]refSpec{"zone_id": {"stackit_dns_zone", "zone_id"}}},

	"stackit_cdn_distribution":  {group: "cdn", kind: "Distribution"},
	"stackit_cdn_custom_domain": {group: "cdn", kind: "CustomDomain", refs: map[string]refSpec{"distribution_id": {"stackit_cdn_distribution", "distribution_id"}}},

	// 4. Object Storage, File Storage, KMS & Secrets
	"stackit_objectstorage_bucket":            {group: "objectstorage", kind: "Bucket"},
	"stackit_objectstorage_compliance_lock":   {group: "objectstorage", kind: "ComplianceLock"},
	"stackit_objectstorage_credentials_group": {group: "objectstorage", kind: "CredentialsGroup"},
	"stackit_objectstorage_credential":        {group: "objectstorage", kind: "Credential", refs: map[string]refSpec{"credentials_group_id": {"stackit_objectstorage_credentials_group", "credentials_group_id"}}},
	"stackit_objectstorage_default_retention": {group: "objectstorage", kind: "DefaultRetention", refs: map[string]refSpec{"bucket_name": {"stackit_objectstorage_bucket", "name"}}},

	"stackit_sfs_export_policy": {group: "sfs", kind: "ExportPolicy"},
	"stackit_sfs_project_lock":  {group: "sfs", kind: "ProjectLock"},
	"stackit_sfs_resource_pool": {group: "sfs", kind: "ResourcePool"},
	"stackit_sfs_share":         {group: "sfs", kind: "Share", refs: map[string]refSpec{"resource_pool_id": {"stackit_sfs_resource_pool", "resource_pool_id"}, "export_policy": {"stackit_sfs_export_policy", "name"}}},

	"stackit_kms_keyring":      {group: "kms", kind: "Keyring"},
	"stackit_kms_key":          {group: "kms", kind: "Key", refs: map[string]refSpec{"keyring_id": {"stackit_kms_keyring", "keyring_id"}}},
	"stackit_kms_wrapping_key": {group: "kms", kind: "WrappingKey", refs: map[string]refSpec{"keyring_id": {"stackit_kms_keyring", "keyring_id"}}},

	"stackit_secretsmanager_instance":                     {group: "secretsmanager", kind: "Instance"},
	"stackit_secretsmanager_user":                         {group: "secretsmanager", kind: "User", refs: map[string]refSpec{"instance_id": {"stackit_secretsmanager_instance", "instance_id"}}},
	"stackit_secretsmanager_instance_role_binding_v1":     {group: "secretsmanager", kind: "InstanceRoleBindingV1", refs: map[string]refSpec{"resource_id": {"stackit_secretsmanager_instance", "instance_id"}}},
	"stackit_secretsmanager_secret_group_role_binding_v1": {group: "secretsmanager", kind: "SecretGroupRoleBindingV1"},

	// 5. Managed Databases & Data Services
	"stackit_postgresflex_instance": {group: "postgresflex", kind: "Instance"},
	"stackit_postgresflex_database": {group: "postgresflex", kind: "Database", refs: map[string]refSpec{"instance_id": {"stackit_postgresflex_instance", "instance_id"}}},
	"stackit_postgresflex_user":     {group: "postgresflex", kind: "User", refs: map[string]refSpec{"instance_id": {"stackit_postgresflex_instance", "instance_id"}}},

	"stackit_sqlserverflex_instance": {group: "sqlserverflex", kind: "Instance"},
	"stackit_sqlserverflex_database": {group: "sqlserverflex", kind: "Database", refs: map[string]refSpec{"instance_id": {"stackit_sqlserverflex_instance", "instance_id"}}},
	"stackit_sqlserverflex_user":     {group: "sqlserverflex", kind: "User", refs: map[string]refSpec{"instance_id": {"stackit_sqlserverflex_instance", "instance_id"}}},

	"stackit_mariadb_instance":   {group: "mariadb", kind: "Instance"},
	"stackit_mariadb_credential": {group: "mariadb", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_mariadb_instance", "instance_id"}}},

	"stackit_mongodbflex_instance": {group: "mongodbflex", kind: "Instance"},
	"stackit_mongodbflex_user":     {group: "mongodbflex", kind: "User", refs: map[string]refSpec{"instance_id": {"stackit_mongodbflex_instance", "instance_id"}}},

	"stackit_redis_instance":   {group: "redis", kind: "Instance"},
	"stackit_redis_credential": {group: "redis", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_redis_instance", "instance_id"}}},

	"stackit_valkey_instance":   {group: "valkey", kind: "Instance"},
	"stackit_valkey_credential": {group: "valkey", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_valkey_instance", "instance_id"}}},

	"stackit_opensearch_instance":   {group: "opensearch", kind: "Instance"},
	"stackit_opensearch_credential": {group: "opensearch", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_opensearch_instance", "instance_id"}}},

	"stackit_rabbitmq_instance":   {group: "rabbitmq", kind: "Instance"},
	"stackit_rabbitmq_credential": {group: "rabbitmq", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_rabbitmq_instance", "instance_id"}}},

	"stackit_dremio_instance": {group: "dremio", kind: "Instance"},
	"stackit_dremio_user":     {group: "dremio", kind: "User", refs: map[string]refSpec{"instance_id": {"stackit_dremio_instance", "instance_id"}}},

	// 6. Observability, Logging & Telemetry
	"stackit_observability_instance":       {group: "observability", kind: "Instance"},
	"stackit_observability_credential":     {group: "observability", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_observability_instance", "instance_id"}}},
	"stackit_observability_alertgroup":     {group: "observability", kind: "Alertgroup", refs: map[string]refSpec{"instance_id": {"stackit_observability_instance", "instance_id"}}},
	"stackit_observability_logalertgroup":  {group: "observability", kind: "Logalertgroup", refs: map[string]refSpec{"instance_id": {"stackit_observability_instance", "instance_id"}}},
	"stackit_observability_scrapeconfig":   {group: "observability", kind: "Scrapeconfig", refs: map[string]refSpec{"instance_id": {"stackit_observability_instance", "instance_id"}}},
	"stackit_logs_instance":                {group: "logs", kind: "Instance"},
	"stackit_logs_access_token":            {group: "logs", kind: "AccessToken", refs: map[string]refSpec{"instance_id": {"stackit_logs_instance", "instance_id"}}},
	"stackit_logme_instance":               {group: "logme", kind: "Instance"},
	"stackit_logme_credential":             {group: "logme", kind: "Credential", refs: map[string]refSpec{"instance_id": {"stackit_logme_instance", "instance_id"}}},
	"stackit_telemetryrouter_instance":     {group: "telemetryrouter", kind: "Instance"},
	"stackit_telemetryrouter_destination":  {group: "telemetryrouter", kind: "Destination", refs: map[string]refSpec{"instance_id": {"stackit_telemetryrouter_instance", "instance_id"}}},
	"stackit_telemetryrouter_access_token": {group: "telemetryrouter", kind: "AccessToken", refs: map[string]refSpec{"instance_id": {"stackit_telemetryrouter_instance", "instance_id"}}},
	"stackit_telemetrylink":                {group: "telemetryrouter", kind: "TelemetryLink", refs: map[string]refSpec{"telemetry_router_id": {"stackit_telemetryrouter_instance", "instance_id"}}},
	"stackit_intake_runner":                {group: "intake", kind: "Runner"},

	// 7. Platform, AI & DevTools
	"stackit_scf_organization":          {group: "scf", kind: "Organization"},
	"stackit_scf_organization_manager":  {group: "scf", kind: "OrganizationManager", refs: map[string]refSpec{"org_id": {"stackit_scf_organization", "org_id"}}},
	"stackit_git":                       {group: "git", kind: "Instance"},
	"stackit_modelexperiments_instance": {group: "modelexperiments", kind: "Instance"},
	"stackit_modelexperiments_token":    {group: "modelexperiments", kind: "Token", refs: map[string]refSpec{"instance_id": {"stackit_modelexperiments_instance", "instance_id"}}},
	"stackit_modelserving_token":        {group: "modelserving", kind: "Token"},
}

func resourceConfigurations() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if r.TerraformResource != nil {
			if s, ok := r.TerraformResource.Schema["project_id"]; ok && (s.Required || s.Optional) {
				r.References["project_id"] = paramRef("stackit_resourcemanager_project", "project_id")
			}
		}
		if spec, ok := resourceSpecs[r.Name]; ok {
			r.ShortGroup = spec.group
			r.Kind = spec.kind
			for field, ref := range spec.refs {
				r.References[field] = paramRef(ref.target, ref.field)
			}
		}
	}
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.kubee.dev"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			resourceConfigurations(),
		))

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.m.kubee.dev"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			resourceConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	pc.ConfigureResources()
	return pc
}
