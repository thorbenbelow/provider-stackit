package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`

	// DefaultRegion will be used as the default location for regional services.
	// +optional
	DefaultRegion *string `json:"defaultRegion,omitempty"`

	// EnableBetaResources enables beta resources.
	// +optional
	EnableBetaResources *bool `json:"enableBetaResources,omitempty"`

	// Experiments enables experiments.
	// +optional
	Experiments []string `json:"experiments,omitempty"`

	// AlbCustomEndpoint is the custom endpoint for the Application Load Balancer service.
	// +optional
	AlbCustomEndpoint *string `json:"albCustomEndpoint,omitempty"`

	// AuthorizationCustomEndpoint is the custom endpoint for the Membership service.
	// +optional
	AuthorizationCustomEndpoint *string `json:"authorizationCustomEndpoint,omitempty"`

	// CdnCustomEndpoint is the custom endpoint for the CDN service.
	// +optional
	CdnCustomEndpoint *string `json:"cdnCustomEndpoint,omitempty"`

	// DnsCustomEndpoint is the custom endpoint for the DNS service.
	// +optional
	DnsCustomEndpoint *string `json:"dnsCustomEndpoint,omitempty"`

	// EdgecloudCustomEndpoint is the custom endpoint for the Edge Cloud service.
	// +optional
	EdgecloudCustomEndpoint *string `json:"edgecloudCustomEndpoint,omitempty"`

	// GitCustomEndpoint is the custom endpoint for the Git service.
	// +optional
	GitCustomEndpoint *string `json:"gitCustomEndpoint,omitempty"`

	// IaasCustomEndpoint is the custom endpoint for the IaaS service.
	// +optional
	IaasCustomEndpoint *string `json:"iaasCustomEndpoint,omitempty"`

	// MariadbCustomEndpoint is the custom endpoint for the MariaDB service.
	// +optional
	MariadbCustomEndpoint *string `json:"mariadbCustomEndpoint,omitempty"`

	// MongodbCustomEndpoint is the custom endpoint for the MongoDB service.
	// +optional
	MongodbCustomEndpoint *string `json:"mongodbCustomEndpoint,omitempty"`

	// MysqlCustomEndpoint is the custom endpoint for the MySQL service.
	// +optional
	MysqlCustomEndpoint *string `json:"mysqlCustomEndpoint,omitempty"`

	// ObjectStorageCustomEndpoint is the custom endpoint for the Object Storage service.
	// +optional
	ObjectStorageCustomEndpoint *string `json:"objectStorageCustomEndpoint,omitempty"`

	// ObservabilityCustomEndpoint is the custom endpoint for the Observability service.
	// +optional
	ObservabilityCustomEndpoint *string `json:"observabilityCustomEndpoint,omitempty"`

	// OpensearchCustomEndpoint is the custom endpoint for the OpenSearch service.
	// +optional
	OpensearchCustomEndpoint *string `json:"opensearchCustomEndpoint,omitempty"`

	// PostgresqlCustomEndpoint is the custom endpoint for the PostgreSQL service.
	// +optional
	PostgresqlCustomEndpoint *string `json:"postgresqlCustomEndpoint,omitempty"`

	// RabbitmqCustomEndpoint is the custom endpoint for the RabbitMQ service.
	// +optional
	RabbitmqCustomEndpoint *string `json:"rabbitmqCustomEndpoint,omitempty"`

	// RedisCustomEndpoint is the custom endpoint for the Redis service.
	// +optional
	RedisCustomEndpoint *string `json:"redisCustomEndpoint,omitempty"`

	// ResourcemanagerCustomEndpoint is the custom endpoint for the Resource Manager service.
	// +optional
	ResourcemanagerCustomEndpoint *string `json:"resourcemanagerCustomEndpoint,omitempty"`

	// SecretsmanagerCustomEndpoint is the custom endpoint for the Secrets Manager service.
	// +optional
	SecretsmanagerCustomEndpoint *string `json:"secretsmanagerCustomEndpoint,omitempty"`

	// SkeCustomEndpoint is the custom endpoint for the SKE service.
	// +optional
	SkeCustomEndpoint *string `json:"skeCustomEndpoint,omitempty"`

	// SqlserverCustomEndpoint is the custom endpoint for the SQL Server service.
	// +optional
	SqlserverCustomEndpoint *string `json:"sqlserverCustomEndpoint,omitempty"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a STACKIT provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,stackit}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,stackit}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
