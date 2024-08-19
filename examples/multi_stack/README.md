# multi-stack

This example demonstrate how to create multiple stacks in one app.

## Using
CLI:
```shell
cd examples/mutli_stack
pk8s export
```

Code:
```shell
cd examples/mutli_stack
go run main.go
```

Output:
```
Generated Kubernetes manifests:
  dev:
    - dist/dev/hello.yaml
  test:
    - dist/test/hello.yaml
```
