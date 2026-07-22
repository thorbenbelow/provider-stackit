package authorization

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_authorization_folder_custom_role", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "FolderCustomRole"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_folder",
		}
	})
	p.AddResourceConfigurator("stackit_authorization_folder_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "FolderRoleAssignment"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_folder",
		}
	})
	p.AddResourceConfigurator("stackit_authorization_organization_custom_role", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "OrganizationCustomRole"
	})
	p.AddResourceConfigurator("stackit_authorization_organization_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "OrganizationRoleAssignment"
	})
	p.AddResourceConfigurator("stackit_authorization_project_custom_role", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "ProjectCustomRole"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
	p.AddResourceConfigurator("stackit_authorization_project_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "ProjectRoleAssignment"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
	p.AddResourceConfigurator("stackit_authorization_service_account_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "ServiceAccountRoleAssignment"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_service_account",
		}
	})
}
