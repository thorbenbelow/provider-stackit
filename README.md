# Upjet-based Crossplane provider for STACKIT

`provider-stackit` is a [Crossplane](https://crossplane.io/) provider built with [Upjet](https://github.com/crossplane/upjet) that enables managing [STACKIT](https://www.stackit.de/en/) cloud resources declaratively using Kubernetes CRDs.

## Getting Started

### 1. Installation

Install the STACKIT provider into your Crossplane cluster:

```console
kubectl apply -f examples/install.yaml
```

Or apply the Provider manifest directly:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-stackit
spec:
  package: xpkg.upbound.io/kubee/provider-stackit:v0.1.1
```

### 2. Configure Credentials

Create a Kubernetes Secret containing your STACKIT service account key in JSON format:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: stackit-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "service_account_key": "YOUR_SERVICE_ACCOUNT_KEY"
    }
```

### 3. Configure ProviderConfig

Reference the Secret in a `ProviderConfig`:

```yaml
apiVersion: stackit.kubee.dev/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: stackit-creds
      namespace: crossplane-system
      key: credentials
```

> [!NOTE]
> Namespaced ProviderConfigs (`ClusterProviderConfig` or namespaced `ProviderConfig`) are also available under the `stackit.m.kubee.dev/v1beta1` API group. See [`./examples`](./examples) for examples.

### 4. Create Managed Resources

Provision STACKIT resources such as projects, networks, object storage buckets, or databases using declarative YAML:

```yaml
apiVersion: resourcemanager.stackit.kubee.dev/v1alpha1
kind: Project
metadata:
  name: example-project
spec:
  forProvider:
    name: "my-crossplane-project"
  providerConfigRef:
    name: default
```

## Supported Services

`provider-stackit` covers 115 STACKIT resources across:
* **Compute & Networking**: IaaS (servers, volumes, security groups, public IPs), ALB (Application Load Balancer), VPC, DNS, VPN
* **Storage & Databases**: Object Storage, SFS (Shared File Storage), PostgresFlex, MongoDBFlex, MariaDB, Redis, Valkey, OpenSearch, SQLServerFlex
* **Security & Governance**: KMS (Key Management Service), Secrets Manager, Authorization (IAM role assignments and custom roles), Resource Manager (Projects & Folders)
* **Observability & Logging**: Observability / Argus, Logs, Telemetry Router
* **Kubernetes & Cloud Native**: SKE (STACKIT Kubernetes Engine), Edge Cloud, Serverless Cloud Functions (SCF), Model Experiments & Model Serving

## Contributing & Local Development

* For the general contribution guide, see the [Upjet Contribution Guide](https://github.com/crossplane/upjet).
* If you'd like to learn how to use Upjet, see the [Usage Guide](https://github.com/crossplane/upjet/blob/main/docs/getting-started.md).

### Building Locally

To build this provider locally:

```console
make build
```

To run code generation:

```console
make generate
```

To run the provider locally against a cluster without packaging:

```console
make run
```

### Packaging & Publishing

To build the package and publish it to the Upbound Registry:

```console
XPKG_REG_ORGS="xpkg.upbound.io/kubee" \
XPKG_REG_ORGS_NO_PROMOTE="xpkg.upbound.io/kubee" \
make build.all publish BRANCH_NAME=main
```

> [!NOTE]
> `BRANCH_NAME` must match `main` to trigger package publishing.

## Add a New Resource

Follow the Upjet guide for [adding new resources](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md#add-a-new-resource). New resource configurations are added under [`./config`](./config).

## Getting Help

* For filing bugs, suggesting improvements, or requesting new resources or features, please open an [issue](https://github.com/thorbenbelow/provider-stackit/issues).

## License

The provider is released under the [Apache 2.0 license with notice](./LICENSE).
