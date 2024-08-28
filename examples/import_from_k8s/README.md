# import-from-k8s

This example demonstrate how to import custom resource definitions (CRDs) from an existing Kubernetes cluster (GKE). See the `imports/` directory for the expected result.

## Using
CLI:
```shell
cd examples/import
kubectl get crds managedcertificates.networking.gke.io -o yaml | go run ./pk8s import
```

Custom package name:
```shell
cd examples/import
kubectl get crds managedcertificates.networking.gke.io -o yaml | go run ./pk8s import my_name:=
```

Output:
```
Importing custom resource definitions, this may take a few moments...
networking.gke.io
  networking.gke.io/managedcertificate
```
