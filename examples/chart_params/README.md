# chart-params

This example demonstrate how to create a chart with params.

## Using
CLI:
```shell
cd examples/chart_params
pk8s export
```

Code:
```shell
cd examples/chart_params
go run main.go
```

Output:
```
Generated Kubernetes manifests:
  dev:
    - dist/dev/hello.yaml
```

Manifest:
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: hello
  name: hello
  namespace: custom
```
