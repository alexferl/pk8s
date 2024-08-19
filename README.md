# pk8s

**!IN DEVELOPMENT!**

pk8s (Programmable Kubernetes) is a software development framework designed to facilitate the creation of Kubernetes manifests.
It allows developers to define, manage, and reuse Kubernetes manifests programmatically, enhancing efficiency and maintainability.

## Concepts

pk8s applications are developed using the Go programming language.

At the core of a pk8s project is the `App`, which serves as the foundational element.
Within an `App`, users have the flexibility to define multiple `Stack`, representing various environments, clusters, or team.
Each `Stack` can further encapsulate multiple `Chart`, which are organized collections of Kubernetes resources such as `Deployment`, `Service`, `Ingress`, `ReplicaSet`, and more.

## Using
### CLI
The easiest way to get started is using the CLI.

#### Quickstart
Install:
```shell
go install github.com/alexferl/pk8s/pk8s@latest
```

Init:
```shell
pk8s init infra
```

Output:
```

  Your pk8s Go project is ready! 🚀✨

    📦 pk8s export  Export Kubernetes manifests to dist/

  Deploy:
    🛠️ kubectl apply -f dist/

```

Change directory:
```shell
cd infra
```

Export:
```shell
pk8s export
```

Output:
```
Generated Kubernetes manifests:
  dev:
    - dist/dev/hello.yaml
```

And if you check the content of `dist/dev/hello.yaml`:
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-24cf642e
```

> Note: the `apiVersion` and `kind` fields are automatically populated on generation for resources that require them.

You can now edit `main.go` and add more stacks and charts and repeat the export process to generate new manifests.

### Code
If you can't or don't want to use the CLI, you can also directly use the library.

#### Quickstart
Install:
```shell
go get github.com/alexferl/pk8s
```

Create a `main.go` file with the following content:

```go
package main

import (
	"github.com/alexferl/pk8s"
	"github.com/alexferl/pk8s/k8s"
)

type ChartParams struct {
	pk8s.ChartParams
}

func NewChart(stack *pk8s.Stack, name string, params *ChartParams) *pk8s.Chart {
	var chartParams pk8s.ChartParams
	if params != nil {
		chartParams = params.ChartParams
	}
	chart := pk8s.NewChart(stack, name, &chartParams)

	deployment := &k8s.DeploymentV1{
		// ...
	}

	chart.Append(deployment)

	return chart
}

func main() {
	app := pk8s.New()
	stack := pk8s.NewStack(app, "dev")

	NewChart(stack, "hello", nil)

	app.Export()
}
```

Run it:
```shell
go run main.go
```

Output:
```
Generated Kubernetes manifests:
  dev:
    - dist/dev/hello.yaml
```

## Examples
See the [examples](examples) directory for more complex examples.

## FAQ
1. Who is this for?
    - Software developers, DevOps engineers, and Kubernetes users who prefer defining Kubernetes resources programmatically rather than using YAML.
2. Why not just use [cdk8s](https://github.com/cdk8s-team/cdk8s)?
    - cdk8s is a robust tool for managing Kubernetes configurations, but its CLI depends on Node.js since it is fundamentally a JavaScript-based project. While it provides bindings for other languages, such as Go, the generated code may occasionally result in less elegant and idiomatic implementations in those languages.
