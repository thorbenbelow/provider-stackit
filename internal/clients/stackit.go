package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/thorbenbelow/provider-stackit/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/thorbenbelow/provider-stackit/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal stackit credentials as JSON"

	keyServiceAccountKey             = "service_account_key"
	keyServiceAccountToken           = "service_account_token"
	keyDefaultRegion                 = "default_region"
	keyEnableBetaResources           = "enable_beta_resources"
	keyExperiments                   = "experiments"
	keyAlbCustomEndpoint             = "alb_custom_endpoint"
	keyAuthorizationCustomEndpoint   = "authorization_custom_endpoint"
	keyCdnCustomEndpoint             = "cdn_custom_endpoint"
	keyDnsCustomEndpoint             = "dns_custom_endpoint"
	keyEdgecloudCustomEndpoint       = "edgecloud_custom_endpoint"
	keyGitCustomEndpoint             = "git_custom_endpoint"
	keyIaasCustomEndpoint            = "iaas_custom_endpoint"
	keyMariadbCustomEndpoint         = "mariadb_custom_endpoint"
	keyMongodbCustomEndpoint         = "mongodb_custom_endpoint"
	keyMysqlCustomEndpoint           = "mysql_custom_endpoint"
	keyObjectStorageCustomEndpoint   = "object_storage_custom_endpoint"
	keyObservabilityCustomEndpoint   = "observability_custom_endpoint"
	keyOpensearchCustomEndpoint      = "opensearch_custom_endpoint"
	keyPostgresqlCustomEndpoint      = "postgresql_custom_endpoint"
	keyRabbitmqCustomEndpoint        = "rabbitmq_custom_endpoint"
	keyRedisCustomEndpoint           = "redis_custom_endpoint"
	keyResourcemanagerCustomEndpoint = "resourcemanager_custom_endpoint"
	keySecretsmanagerCustomEndpoint  = "secretsmanager_custom_endpoint"
	keySkeCustomEndpoint             = "ske_custom_endpoint"
	keySqlserverCustomEndpoint       = "sqlserver_custom_endpoint"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		for k, v := range creds {
			ps.Configuration[k] = v
		}

		// Override with fields from ProviderConfigSpec if set.
		if pcSpec.DefaultRegion != nil {
			ps.Configuration[keyDefaultRegion] = *pcSpec.DefaultRegion
		}
		if pcSpec.EnableBetaResources != nil {
			ps.Configuration[keyEnableBetaResources] = *pcSpec.EnableBetaResources
		}
		if len(pcSpec.Experiments) > 0 {
			ps.Configuration[keyExperiments] = pcSpec.Experiments
		}
		if pcSpec.AlbCustomEndpoint != nil {
			ps.Configuration[keyAlbCustomEndpoint] = *pcSpec.AlbCustomEndpoint
		}
		if pcSpec.AuthorizationCustomEndpoint != nil {
			ps.Configuration[keyAuthorizationCustomEndpoint] = *pcSpec.AuthorizationCustomEndpoint
		}
		if pcSpec.CdnCustomEndpoint != nil {
			ps.Configuration[keyCdnCustomEndpoint] = *pcSpec.CdnCustomEndpoint
		}
		if pcSpec.DnsCustomEndpoint != nil {
			ps.Configuration[keyDnsCustomEndpoint] = *pcSpec.DnsCustomEndpoint
		}
		if pcSpec.EdgecloudCustomEndpoint != nil {
			ps.Configuration[keyEdgecloudCustomEndpoint] = *pcSpec.EdgecloudCustomEndpoint
		}
		if pcSpec.GitCustomEndpoint != nil {
			ps.Configuration[keyGitCustomEndpoint] = *pcSpec.GitCustomEndpoint
		}
		if pcSpec.IaasCustomEndpoint != nil {
			ps.Configuration[keyIaasCustomEndpoint] = *pcSpec.IaasCustomEndpoint
		}
		if pcSpec.MariadbCustomEndpoint != nil {
			ps.Configuration[keyMariadbCustomEndpoint] = *pcSpec.MariadbCustomEndpoint
		}
		if pcSpec.MongodbCustomEndpoint != nil {
			ps.Configuration[keyMongodbCustomEndpoint] = *pcSpec.MongodbCustomEndpoint
		}
		if pcSpec.MysqlCustomEndpoint != nil {
			ps.Configuration[keyMysqlCustomEndpoint] = *pcSpec.MysqlCustomEndpoint
		}
		if pcSpec.ObjectStorageCustomEndpoint != nil {
			ps.Configuration[keyObjectStorageCustomEndpoint] = *pcSpec.ObjectStorageCustomEndpoint
		}
		if pcSpec.ObservabilityCustomEndpoint != nil {
			ps.Configuration[keyObservabilityCustomEndpoint] = *pcSpec.ObservabilityCustomEndpoint
		}
		if pcSpec.OpensearchCustomEndpoint != nil {
			ps.Configuration[keyOpensearchCustomEndpoint] = *pcSpec.OpensearchCustomEndpoint
		}
		if pcSpec.PostgresqlCustomEndpoint != nil {
			ps.Configuration[keyPostgresqlCustomEndpoint] = *pcSpec.PostgresqlCustomEndpoint
		}
		if pcSpec.RabbitmqCustomEndpoint != nil {
			ps.Configuration[keyRabbitmqCustomEndpoint] = *pcSpec.RabbitmqCustomEndpoint
		}
		if pcSpec.RedisCustomEndpoint != nil {
			ps.Configuration[keyRedisCustomEndpoint] = *pcSpec.RedisCustomEndpoint
		}
		if pcSpec.ResourcemanagerCustomEndpoint != nil {
			ps.Configuration[keyResourcemanagerCustomEndpoint] = *pcSpec.ResourcemanagerCustomEndpoint
		}
		if pcSpec.SecretsmanagerCustomEndpoint != nil {
			ps.Configuration[keySecretsmanagerCustomEndpoint] = *pcSpec.SecretsmanagerCustomEndpoint
		}
		if pcSpec.SkeCustomEndpoint != nil {
			ps.Configuration[keySkeCustomEndpoint] = *pcSpec.SkeCustomEndpoint
		}
		if pcSpec.SqlserverCustomEndpoint != nil {
			ps.Configuration[keySqlserverCustomEndpoint] = *pcSpec.SqlserverCustomEndpoint
		}

		return ps, nil
	}
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
