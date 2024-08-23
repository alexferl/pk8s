package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/ettle/strcase"
)

var version *string

var caser *strcase.Caser

// common initialisms for k8s
var initialisms = map[string]bool{
	"AWS":   true,
	"CA":    true,
	"CEL":   true,
	"CIDR":  true,
	"CSI":   true,
	"DB":    true,
	"FC":    true,
	"FS":    true,
	"GCE":   true,
	"GID":   true,
	"GMSA":  true,
	"GRPC":  true,
	"IO":    true,
	"IPC":   true,
	"ISCSI": true,
	"NFS":   true,
	"OS":    true,
	"PID":   true,
	"QOS":   true,
	"RBD":   true,
	"SE":    true,
	"SSL":   true,
	"TLS":   true,
	"TTY":   true,
}

func main() {
	version = flag.String("version", "", "Kubernetes version")
	flag.Parse()

	if *version == "" {
		log.Fatalf("-version is required")
	}

	caser = strcase.NewCaser(true, initialisms, nil)
	data := downloadSwaggerSpec(*version)

	p := New()
	p.Process(data)

	generateGroupVersionKind(data)
	generateUtils()

	err := runGoFmt("k8s")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func downloadSwaggerSpec(version string) []byte {
	url := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/%s/api/openapi-spec/swagger.json", version)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	return data
}

func runGoFmt(dir string) error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go fmt failed: %v", err)
	}
	return nil
}
