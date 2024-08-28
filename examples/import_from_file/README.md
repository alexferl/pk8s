# import-from-file

This example demonstrate how to import custom resource definitions (CRDs) from a file. See the `imports/` directory for the expected result.

## Using
CLI:
```shell
cd examples/import
pk8s import managedcertificates-crd.yaml
```

Using a custom package name:
```shell
cd examples/import
pk8s import my_name:=managedcertificates-crd.yaml
```

Output:
```
Importing custom resource definitions, this may take a few moments...
networking.gke.io
  networking.gke.io/managedcertificate
```
