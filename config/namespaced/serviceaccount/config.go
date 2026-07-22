package serviceaccount

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_service_account", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
		r.Kind = "ServiceAccount"
		r.References["project_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
	p.AddResourceConfigurator("stackit_service_account_federated_identity_provider", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
		r.Kind = "FederatedIdentityProvider"
		r.References["project_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
	p.AddResourceConfigurator("stackit_service_account_key", func(r *config.Resource) {
		r.ShortGroup = "serviceaccount"
		r.Kind = "Key"
		r.References["project_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
		r.References["service_account_email"] = config.Reference{
			TerraformName: "stackit_service_account",
		}
	})
}
