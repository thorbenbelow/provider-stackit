# STACKIT Provider for Crossplane

`provider-stackit` is a [Crossplane](https://crossplane.io/) provider built with [Upjet](https://github.com/crossplane/upjet) that enables managing [STACKIT](https://www.stackit.de/en/) cloud resources declaratively using Kubernetes CRDs.

## Getting Started

### 1. Installation

Install the STACKIT provider into your Crossplane cluster:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-stackit
spec:
  package: xpkg.upbound.io/kubee/provider-stackit:v0.1.1
```

Or apply directly:

```console
kubectl apply -f examples/install.yaml
```

### 2. Configure Credentials

Create a Kubernetes Secret containing your STACKIT service account credentials in JSON format:

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
apiVersion: stackit.crossplane.io/v1beta1
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

`provider-stackit` covers STACKIT services including:
* **Compute & Networking**: IaaS (servers, volumes, security groups, public IPs), ALB (Application Load Balancer), VPC, DNS, VPN
* **Storage & Databases**: Object Storage, SFS (Shared File Storage), PostgresFlex, MongoDBFlex, MariaDB, Redis, Valkey, OpenSearch, SQLServerFlex
* **Security & Governance**: KMS (Key Management Service), Secrets Manager, Authorization (IAM role assignments and custom roles), Resource Manager (Projects)
* **Observability & Logging**: Observability / Argus, Logs, Telemetry Router
* **Kubernetes & Cloud Native**: SKE (STACKIT Kubernetes Engine), Edge Cloud, Serverless Cloud Functions (SCF), Model Experiments & Model Serving

## Repository & Contributing

* GitHub Repository: [thorbenbelow/provider-stackit](https://github.com/thorbenbelow/provider-stackit)
* Issues & Feature Requests: [Issue Tracker](https://github.com/thorbenbelow/provider-stackit/issues)
* License: [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0)
