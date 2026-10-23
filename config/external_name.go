package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"null_resource": idWithStub(),

	"stackit_affinity_group":                                config.IdentifierFromProvider,
	"stackit_alb_certificate":                               config.IdentifierFromProvider,
	"stackit_alb_waf_configuration":                         config.IdentifierFromProvider,
	"stackit_alb_waf_custom_rule_group":                     config.IdentifierFromProvider,
	"stackit_alb_waf_managed_rule_set":                      config.IdentifierFromProvider,
	"stackit_application_load_balancer":                     config.IdentifierFromProvider,
	"stackit_authorization_folder_custom_role":              config.IdentifierFromProvider,
	"stackit_authorization_folder_role_assignment":          config.IdentifierFromProvider,
	"stackit_authorization_organization_custom_role":        config.IdentifierFromProvider,
	"stackit_authorization_organization_role_assignment":    config.IdentifierFromProvider,
	"stackit_authorization_project_custom_role":             config.IdentifierFromProvider,
	"stackit_authorization_project_role_assignment":         config.IdentifierFromProvider,
	"stackit_authorization_service_account_role_assignment": config.IdentifierFromProvider,
	"stackit_cdn_custom_domain":                             config.IdentifierFromProvider,
	"stackit_cdn_distribution":                              config.IdentifierFromProvider,
	"stackit_dns_record_set":                                config.IdentifierFromProvider,
	"stackit_dns_zone":                                      config.IdentifierFromProvider,
	"stackit_dremio_instance":                               config.IdentifierFromProvider,
	"stackit_dremio_user":                                   config.IdentifierFromProvider,
	"stackit_edgecloud_instance":                            config.IdentifierFromProvider,
	"stackit_edgecloud_kubeconfig":                          config.IdentifierFromProvider,
	"stackit_edgecloud_token":                               config.IdentifierFromProvider,
	"stackit_git":                                           config.IdentifierFromProvider,
	"stackit_image":                                         config.IdentifierFromProvider,
	"stackit_intake_runner":                                 config.IdentifierFromProvider,
	"stackit_key_pair":                                      config.IdentifierFromProvider,
	"stackit_kms_key":                                       config.IdentifierFromProvider,
	"stackit_kms_keyring":                                   config.IdentifierFromProvider,
	"stackit_kms_wrapping_key":                              config.IdentifierFromProvider,
	"stackit_loadbalancer":                                  config.IdentifierFromProvider,
	"stackit_loadbalancer_observability_credential":         config.IdentifierFromProvider,
	"stackit_logme_credential":                              config.IdentifierFromProvider,
	"stackit_logme_instance":                                config.IdentifierFromProvider,
	"stackit_logs_access_token":                             config.IdentifierFromProvider,
	"stackit_logs_instance":                                 config.IdentifierFromProvider,
	"stackit_mariadb_credential":                            config.IdentifierFromProvider,
	"stackit_mariadb_instance":                              config.IdentifierFromProvider,
	"stackit_modelexperiments_instance":                     config.IdentifierFromProvider,
	"stackit_modelexperiments_token":                        config.IdentifierFromProvider,
	"stackit_modelserving_token":                            config.IdentifierFromProvider,
	"stackit_mongodbflex_instance":                          config.IdentifierFromProvider,
	"stackit_mongodbflex_user":                              config.IdentifierFromProvider,
	"stackit_network":                                       config.IdentifierFromProvider,
	"stackit_network_area":                                  config.IdentifierFromProvider,
	"stackit_network_area_region":                           config.IdentifierFromProvider,
	"stackit_network_area_route":                            config.IdentifierFromProvider,
	"stackit_network_interface":                             config.IdentifierFromProvider,
	"stackit_objectstorage_bucket":                          config.IdentifierFromProvider,
	"stackit_objectstorage_compliance_lock":                 config.IdentifierFromProvider,
	"stackit_objectstorage_credential":                      config.IdentifierFromProvider,
	"stackit_objectstorage_credentials_group":               config.IdentifierFromProvider,
	"stackit_objectstorage_default_retention":               config.IdentifierFromProvider,
	"stackit_observability_alertgroup":                      config.IdentifierFromProvider,
	"stackit_observability_credential":                      config.IdentifierFromProvider,
	"stackit_observability_instance":                        config.IdentifierFromProvider,
	"stackit_observability_logalertgroup":                   config.IdentifierFromProvider,
	"stackit_observability_scrapeconfig":                    config.IdentifierFromProvider,
	"stackit_opensearch_credential":                         config.IdentifierFromProvider,
	"stackit_opensearch_instance":                           config.IdentifierFromProvider,
	"stackit_postgresflex_database":                         config.IdentifierFromProvider,
	"stackit_postgresflex_instance":                         config.IdentifierFromProvider,
	"stackit_postgresflex_user":                             config.IdentifierFromProvider,
	"stackit_public_ip":                                     config.IdentifierFromProvider,
	"stackit_public_ip_associate":                           config.IdentifierFromProvider,
	"stackit_rabbitmq_credential":                           config.IdentifierFromProvider,
	"stackit_rabbitmq_instance":                             config.IdentifierFromProvider,
	"stackit_redis_credential":                              config.IdentifierFromProvider,
	"stackit_redis_instance":                                config.IdentifierFromProvider,
	"stackit_resourcemanager_folder":                        config.IdentifierFromProvider,
	"stackit_resourcemanager_project":                       config.IdentifierFromProvider,
	"stackit_routing_table":                                 config.IdentifierFromProvider,
	"stackit_routing_table_route":                           config.IdentifierFromProvider,
	"stackit_scf_organization":                              config.IdentifierFromProvider,
	"stackit_scf_organization_manager":                      config.IdentifierFromProvider,
	"stackit_secretsmanager_instance":                       config.IdentifierFromProvider,
	"stackit_secretsmanager_instance_role_binding_v1":       config.IdentifierFromProvider,
	"stackit_secretsmanager_secret_group_role_binding_v1":   config.IdentifierFromProvider,
	"stackit_secretsmanager_user":                           config.IdentifierFromProvider,
	"stackit_security_group":                                config.IdentifierFromProvider,
	"stackit_security_group_rule":                           config.IdentifierFromProvider,
	"stackit_server":                                        config.IdentifierFromProvider,
	"stackit_server_backup_enable":                          config.IdentifierFromProvider,
	"stackit_server_backup_schedule":                        config.IdentifierFromProvider,
	"stackit_server_network_interface_attach":               config.IdentifierFromProvider,
	"stackit_server_service_account_attach":                 config.IdentifierFromProvider,
	"stackit_server_update_enable":                          config.IdentifierFromProvider,
	"stackit_server_update_schedule":                        config.IdentifierFromProvider,
	"stackit_server_volume_attach":                          config.IdentifierFromProvider,
	"stackit_service_account":                               config.IdentifierFromProvider,
	"stackit_service_account_federated_identity_provider":   config.IdentifierFromProvider,
	"stackit_service_account_key":                           config.IdentifierFromProvider,
	"stackit_sfs_export_policy":                             config.IdentifierFromProvider,
	"stackit_sfs_project_lock":                              config.IdentifierFromProvider,
	"stackit_sfs_resource_pool":                             config.IdentifierFromProvider,
	"stackit_sfs_share":                                     config.IdentifierFromProvider,
	"stackit_ske_cluster":                                   config.IdentifierFromProvider,
	"stackit_ske_kubeconfig":                                config.IdentifierFromProvider,
	"stackit_sqlserverflex_database":                        config.IdentifierFromProvider,
	"stackit_sqlserverflex_instance":                        config.IdentifierFromProvider,
	"stackit_sqlserverflex_user":                            config.IdentifierFromProvider,
	"stackit_telemetrylink":                                 config.IdentifierFromProvider,
	"stackit_telemetryrouter_access_token":                  config.IdentifierFromProvider,
	"stackit_telemetryrouter_destination":                   config.IdentifierFromProvider,
	"stackit_telemetryrouter_instance":                      config.IdentifierFromProvider,
	"stackit_valkey_credential":                             config.IdentifierFromProvider,
	"stackit_valkey_instance":                               config.IdentifierFromProvider,
	"stackit_volume":                                        config.IdentifierFromProvider,
	"stackit_volume_automation":                             config.IdentifierFromProvider,
	"stackit_vpc":                                           config.IdentifierFromProvider,
	"stackit_vpc_network_range":                             config.IdentifierFromProvider,
	"stackit_vpc_region":                                    config.IdentifierFromProvider,
	"stackit_vpc_routing_table":                             config.IdentifierFromProvider,
	"stackit_vpc_routing_table_static_route":                config.IdentifierFromProvider,
	"stackit_vpn_connection":                                config.IdentifierFromProvider,
	"stackit_vpn_gateway":                                   config.IdentifierFromProvider,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
