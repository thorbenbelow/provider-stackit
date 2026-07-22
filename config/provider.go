package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	resourcemanagerCluster "github.com/thorbenbelow/provider-stackit/config/cluster/resourcemanager"
	serviceaccountCluster "github.com/thorbenbelow/provider-stackit/config/cluster/serviceaccount"
	authorizationCluster "github.com/thorbenbelow/provider-stackit/config/cluster/authorization"
	resourcemanagerNamespaced "github.com/thorbenbelow/provider-stackit/config/namespaced/resourcemanager"
	serviceaccountNamespaced "github.com/thorbenbelow/provider-stackit/config/namespaced/serviceaccount"
	authorizationNamespaced "github.com/thorbenbelow/provider-stackit/config/namespaced/authorization"
)

const (
	resourcePrefix = "stackit"
	modulePath     = "github.com/thorbenbelow/provider-stackit"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stackit.kubee.dev"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		resourcemanagerCluster.Configure,
		serviceaccountCluster.Configure,
		authorizationCluster.Configure,
	} {
		configure(pc)
	}

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
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		resourcemanagerNamespaced.Configure,
		serviceaccountNamespaced.Configure,
		authorizationNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
