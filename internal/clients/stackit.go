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
		creds := map[string]any{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		configKeys := []string{
			"alb_certificates_custom_endpoint",
			"alb_custom_endpoint",
			"alb_waf_custom_endpoint",
			"authorization_custom_endpoint",
			"automation_custom_endpoint",
			"cdn_custom_endpoint",
			"credentials_path",
			"default_region",
			"dns_custom_endpoint",
			"dremio_custom_endpoint",
			"edgecloud_custom_endpoint",
			"enable_beta_resources",
			"experiments",
			"git_custom_endpoint",
			"iaas_custom_endpoint",
			"intake_custom_endpoint",
			"kms_custom_endpoint",
			"loadbalancer_custom_endpoint",
			"logme_custom_endpoint",
			"logs_custom_endpoint",
			"mariadb_custom_endpoint",
			"modelexperiments_custom_endpoint",
			"modelserving_custom_endpoint",
			"mongodbflex_custom_endpoint",
			"objectstorage_custom_endpoint",
			"observability_custom_endpoint",
			"oidc_request_token",
			"oidc_request_url",
			"opensearch_custom_endpoint",
			"postgresflex_custom_endpoint",
			"private_key",
			"private_key_path",
			"rabbitmq_custom_endpoint",
			"redis_custom_endpoint",
			"resourcemanager_custom_endpoint",
			"scf_custom_endpoint",
			"secretsmanager_custom_endpoint",
			"server_backup_custom_endpoint",
			"server_update_custom_endpoint",
			"service_account_custom_endpoint",
			"service_account_email",
			"service_account_federated_token",
			"service_account_federated_token_path",
			"service_account_key",
			"service_account_key_path",
			"service_account_token",
			"service_connection_id",
			"service_enablement_custom_endpoint",
			"sfs_custom_endpoint",
			"ske_custom_endpoint",
			"sqlserverflex_custom_endpoint",
			"telemetrylink_custom_endpoint",
			"telemetryrouter_custom_endpoint",
			"token_custom_endpoint",
			"use_oidc",
			"valkey_custom_endpoint",
			"vpn_custom_endpoint",
		}

		ps.Configuration = make(map[string]any)
		for _, key := range configKeys {
			if val, ok := creds[key]; ok {
				ps.Configuration[key] = val
			}
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
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
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
