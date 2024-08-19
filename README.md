# pk8s
pk8s (programmable kubernetes)

## CLI

### Quickstart
Install:
```shell
go install github.com/alexferl/pk8s/pk8s
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

## Code
If you can't or don't want to use the CLI, you can also directly use the library:

### Quickstart
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
