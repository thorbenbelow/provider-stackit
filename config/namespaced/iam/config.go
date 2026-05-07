package iam

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_service_account", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "ServiceAccount"
		r.References["project_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
	p.AddResourceConfigurator("stackit_authorization_project_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.Kind = "ProjectRoleAssignment"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
}
