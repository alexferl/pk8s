# multi-resources

This example demonstrate how to create a chart with multiple Kubernetes resources.

## Using
CLI:
```shell
cd examples/mutli_resources
pk8s export
```

Code:
```shell
cd examples/mutli_resources
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
  name: hello-cbb5a023
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
      - image: nginx:1.27.1
        name: nginx
        ports:
        - containerPort: 80
          name: http
          protocol: TCP
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: hello
  name: hello-7d4d7d2e
spec:
  ports:
  - port: 80
    targetPort: http
  selector:
    app: hello
  type: LoadBalancer

```
