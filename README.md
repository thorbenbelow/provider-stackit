# Upjet-based Crossplane provider for STACKIT

`provider-stackit` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/crossplane/upjet) code generation tools and exposes XRM-conformant managed resources for [STACKIT](https://www.stackit.de/en/).

## Getting Started

Follow the quick start guide in the [official Crossplane documentation](https://docs.crossplane.io/).

For details on configuring authentication via Secret and `ProviderConfig`, see the manifests in [`./examples`](./examples).

To install the provider package:

```console
kubectl apply -f examples/install.yaml
```

## Contributing

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

To build the package image and publish it to your container registry:

```console
REGISTRY_ORGS="ghcr.io/<your-org>" \
XPKG_REG_ORGS="ghcr.io/<your-org>" \
XPKG_REG_ORGS_NO_PROMOTE="" \
make build.all publish BRANCH_NAME=main
```

> Note: `BRANCH_NAME` is set to `main` (even if you are on another branch) to allow publishing images to your registry account.

### Installing to a Local Cluster

To install the provider into a local Kubernetes cluster with Crossplane already installed, apply the following manifest:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-stackit
spec:
  package: ghcr.io/<your-org>/provider-stackit:<version>
```

You can also run the provider locally against a cluster without packaging:

```console
make run
```

## Add a New Resource

Follow the Upjet guide for [adding new resources](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md#add-a-new-resource). New resource configurations are added under [`./config`](./config).

## Getting Help

* For filing bugs, suggesting improvements, or requesting new resources or features, please open an [issue](https://github.com/thorbenbelow/provider-stackit/issues).

## License

The provider is released under the [Apache 2.0 license with notice](./LICENSE).
