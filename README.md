# gopass-external-secrets
<p align="center">
  <a href="https://github.com/lucas-ingemar/gopass-external-secrets/releases" style="text-decoration: none">
    <img alt="GitHub release" src="https://img.shields.io/github/v/release/lucas-ingemar/gopass-external-secrets?style=flat-square&label=Latest%20Version">
  </a>

  <img alt="GitHub Stars" src="https://img.shields.io/github/stars/lucas-ingemar/gopass-external-secrets?style=flat-square&logo=github&label=Stars">
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/lucas-ingemar/gopass-external-secrets?style=flat-square&logo=go&label=Version">
  <img alt="License" src="https://img.shields.io/github/license/lucas-ingemar/gopass-external-secrets?style=flat-square&label=License">

  <a href="https://goreportcard.com/badge/github.com/lucas-ingemar/gopass-external-secrets" style="text-decoration: none">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/lucas-ingemar/gopass-external-secrets?style=flat-square">
  </a>

  <img alt="Tests" src="https://img.shields.io/github/actions/workflow/status/lucas-ingemar/gopass-external-secrets/go-test.yaml?style=flat-square&logo=github&label=Tests">

  <img alt="Build" src="https://img.shields.io/github/actions/workflow/status/lucas-ingemar/gopass-external-secrets/publish-container.yaml?style=flat-square&logo=github&label=Build">
</p>

Use [External Secrets Operator](https://external-secrets.io) together with [gopass](https://www.gopass.pw/). This app lets you access passwords stored in a gopass repository and use them in a Kubernetes cluster. You simply create ExternalSecrets manifests with references to keys to be looked up in your gopass store.
This makes it easy and safe to store secrets in your commited manifests.

## Installation
In the [deployment folder](deployment) you can find manifests to use for deploying in Kubernetes. The concept with this app is that you inject a sidecar to the external-secrets deployment. This makes it possible for the External Secrets Operator to communicate over localhost, that leads to the benefit of not having to open any port to gopass. This makes it impossible for other apps to gain access to your secrets.

## Configuration
You need to deploy one secret manually to set the necessary secret parameters for gopass. If you want to use the example deployment that manifest must look like this:

``` yaml
apiVersion: v1
kind: Secret
metadata:
  name: gopass-envs
  namespace: external-secrets
type: Opaque
data:
  # All the following commands needs to be base64 encoded before added to this file. (some of them are encoded twice)
  gpgSecret: <ASCII ARMOURED PGP PRIVATE KEY> # Command to generate: gpg --export-secret-keys -a <KEY_ID> | base64 -w 0
  gpgKeyId: <PGP PUBLIC KEY ID>
  secretsRepoUrl: <GIT CLONE HTTP-URL TO SECRETS REPO>
```

* `gpgSecret`: The pgp key used to extract the secrets. Use the following command to genererate the input for this variable: `gpg --export-secret-keys -a <KEY_ID> | base64 -w 0`
* `gpgKeyId`: The ID of the pgp key.
* `secretsRepoUrl`: The URL for the repo containing the gopass secrets. Note: cannot be accessed over SSH. 

## Other configuration
In the [kustomization file](deployment/kustomization.yaml) you have a few other configuration parameters. 
* `GOPASS_PREFIX`: Must be the same string as `prefix` defined in [Gopass settings](#Gopass settings). `default: external-secrets`
* `AUTH_ACTIVE`: Set if basic auth should be enabled for every request. Not needed when running in sidecar mode. `default: false`
* `GIT_COOLDOWN`: When a parameter does not exist in the local gopass store, a pull will be performed if there hasn't been a pull done for `GIT_COOLDOWN` minutes. `default: 5`
* `GIT_PULL_CRON`: Cron schedule for for pulling new passwords. `default: */15 * * * *`
* `LOG_LEVEL`: Log level in the container. `default: info`

## Gopass settings
There is a [file](.gopass/create/2-external-secret.yml) with a predefined template of creating new secrets with `gopass create --store <your external-secrets store>`. Just place the file in `<gopass-repository>/.gopass/create/`. Evrery time you create a new password you can use this template to create secrets on the correct format.

## ExternalSecret example
``` yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: webhook-example
  namespace: testing
spec:
  refreshInterval: "15s"
  secretStoreRef:
    name: gopass
    kind: ClusterSecretStore
  target:
    name: webhook-example
  data:
  - secretKey: username
    remoteRef:
      key: testing/webhook-example
      property: username
  - secretKey: password
    remoteRef:
      key: testing/webhook-example
      property: password
```
The design of this app is that the key has the format `namespace/secret` and everything stored in the secret is `property`. For this particular example the output of `gopass ls` looks the following:

``` bash
  gopass ls
gopass
└── k8s_test_secrets (/home/username/.local/share/gopass/stores/k8s_test_secrets)
    └── external-secrets/
        └── testing/
            └── webhook-example

```


