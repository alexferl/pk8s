# import-from-url

This example demonstrate how to import custom resource definitions (CRDs) from a URL. See the `imports/` directory for the expected result.

## Using
CLI:
```shell
cd examples/import
pk8s import https://raw.githubusercontent.com/GoogleCloudPlatform/gke-managed-certs/v1.2.12/deploy/managedcertificates-crd.yaml
```

Using a custom package name:
```shell
cd examples/import
pk8s import my_name:=https://raw.githubusercontent.com/GoogleCloudPlatform/gke-managed-certs/v1.2.12/deploy/managedcertificates-crd.yaml
```

Output:
```
Importing custom resource definitions, this may take a few moments...
networking.gke.io
  networking.gke.io/managedcertificate
```
