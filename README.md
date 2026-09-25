# pulumi-yandex — a Pulumi provider for Yandex Cloud

A Terraform bridge over [`yandex-cloud/terraform-provider-yandex`](https://github.com/yandex-cloud/terraform-provider-yandex),
maintained for [Simple Container](https://github.com/simple-container-com/api) so that SC can
provision Yandex Cloud the same way it provisions AWS and GCP.

Forked from [`masikrus/pulumi-yandex`](https://github.com/masikrus/pulumi-yandex) (Apache-2.0).
Upstream Terraform provider is MPL-2.0; attribution is preserved.

## Why this fork exists

Two reasons, both load-bearing:

1. **`pulumi/pulumi-yandex` is archived at v0.13.0 (2022-02-22)** — it predates Serverless
   Containers entirely.
2. **The dynamic Terraform bridge (`pulumi package add terraform-provider …`) is unusable from
   Russia.** It needs `registry.opentofu.org` at both SDK-generation *and* plugin-download time,
   and both that registry and `registry.terraform.io` are geo-blocked there. A *static* bridge —
   this repo — resolves through `api.github.com` and the Go module proxy only, so it works from
   both sides.

## What this fork changes over `masikrus/pulumi-yandex` v1.0.1

- **It bridges the whole provider, not half of it.** `terraform-provider-yandex` is a *muxed*
  server: an SDKv2 provider plus a plugin-framework (protocol v6) provider. The parent fork
  bridged only the SDKv2 half, so everything Yandex has migrated to the plugin framework was
  silently missing — including `yandex_iam_service_account`, `yandex_container_registry`,
  `yandex_kms_symmetric_key`, `yandex_resourcemanager_folder_iam_member` and
  `yandex_serverless_container_iam_binding`. Without those there is no service account to run a
  container as, no registry to push its image to, and no way to grant it anything.
  **88 → 268 resources.**
- Upstream bumps: `terraform-provider-yandex` v0.187.0 → v0.229.0,
  `pulumi-terraform-bridge` v3.121.0 → v3.140.0.
- `PluginDownloadURL` points at this repo's GitHub Releases, so consumers need no
  `pulumi plugin install` step and no per-resource plugin option — codegen bakes it into the
  SDK's `PkgResourceDefaultOpts`.
- **Go SDK only.** The .NET/Node/Python SDKs were dropped rather than left to rot at 88 resources.

## Using it

```
go get github.com/simple-container-com/pulumi-yandex/sdk@latest
```

The plugin binary is resolved automatically from this repo's releases. To install it by hand:

```
pulumi plugin install resource yandex 1.1.0 \
  --server github://api.github.com/simple-container-com/pulumi-yandex
```

## Maintenance

`make tfgen && make build_sdks`, then tag `vX.Y.Z` — the release workflow builds the plugin for
linux/darwin × amd64/arm64 and publishes the archives the engine expects. Tag `sdk/vX.Y.Z` too,
for the Go SDK submodule. Because `RespectSchemaVersion` is on, the release tag and
`PROVIDER_VERSION` in the Makefile must match exactly.

**Both `cmd/pulumi-tfgen-yandex` and `cmd/pulumi-resource-yandex` must use the `pkg/pf/tfbridge`
`MainWithMuxer` entrypoints** — see the comments in those files. The plain non-`pf` `Main`
compiles and generates a schema that looks complete, then panics at the engine's first runtime
call. Nothing catches it before a real `pulumi` run.

---

# Yandex Cloud Pulumi Provider Configuration

This document describes the available configuration parameters for the Yandex Cloud Pulumi provider.

## Authentication Parameters

- **token** - Yandex Cloud OAuth token or IAM token for authentication  
  Env: `YANDEX_CLOUD_TOKEN`

- **service_account_key_file** - Path to service account key file in JSON format  
  Env: `YANDEX_CLOUD_SERVICE_ACCOUNT_KEY_FILE`

- **shared_credentials_file** - Path to shared credentials file  
  Env: `YANDEX_CLOUD_SHARED_CREDENTIALS_FILE`

- **profile** - Profile name from shared credentials file  
  Env: `YANDEX_CLOUD_PROFILE`

## Resource Location Parameters

- **folder_id** - Default folder ID for resources  
  Env: `YANDEX_CLOUD_FOLDER_ID`

- **cloud_id** - Default cloud ID for resources  
  Env: `YANDEX_CLOUD_CLOUD_ID`

- **organization_id** - Default organization ID  
  Env: `YANDEX_CLOUD_ORGANIZATION_ID`

- **region_id** - Default region (e.g., "ru-central1")  
  Env: `YANDEX_CLOUD_REGION_ID`

- **zone** - Default availability zone (e.g., "ru-central1-a")  
  Env: `YANDEX_CLOUD_ZONE`

## API Configuration

- **endpoint** - Custom API endpoint URL  
  Env: `YANDEX_CLOUD_ENDPOINT`

- **insecure** - Allow insecure connections to API (true/false)  
  Env: `YANDEX_CLOUD_INSECURE`

- **plaintext** - Disable TLS for API connections (true/false)  
  Env: `YANDEX_CLOUD_PLAINTEXT`

- **max_retries** - Maximum number of API retries  
  Env: `YANDEX_CLOUD_MAX_RETRIES`

## Storage Configuration

- **storage_endpoint** - Custom Object Storage endpoint  
  Env: `YANDEX_CLOUD_STORAGE_ENDPOINT`

- **storage_access_key** - Object Storage access key  
  Env: `YANDEX_CLOUD_STORAGE_ACCESS_KEY`

- **storage_secret_key** - Object Storage secret key  
  Env: `YANDEX_CLOUD_STORAGE_SECRET_KEY`

## Message Queue (YMQ) Configuration

- **ymq_endpoint** - Custom Yandex Message Queue endpoint  
  Env: `YANDEX_CLOUD_YMQ_ENDPOINT`

- **ymq_access_key** - YMQ access key  
  Env: `YANDEX_CLOUD_YMQ_ACCESS_KEY`

- **ymq_secret_key** - YMQ secret key  
  Env: `YANDEX_CLOUD_YMQ_SECRET_KEY`

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/yandex
```

or `yarn`:

```bash
yarn add @pulumi/yandex
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_yandex
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-yandex/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Yandex
```

## Configuration

The following configuration points are available for the `yandex` provider:

- `yandex:region` (environment: `YANDEX_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/yandex/api-docs/).
