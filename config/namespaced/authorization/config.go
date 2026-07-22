package authorization

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_authorization_project_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "authorization"
		r.Kind = "ProjectRoleAssignment"
		r.References["resource_id"] = config.Reference{
			TerraformName: "stackit_resourcemanager_project",
		}
	})
}
