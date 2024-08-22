// Code generated; DO NOT EDIT.

package k8s

// IngressTLSV1 IngressTLS describes the transport layer security associated with an ingress.
type IngressTLSV1 struct {
	// hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	Hosts []string `json:"hosts,omitempty"`
	// secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the "Host" header is used for routing.
	SecretName *string `json:"secretName,omitempty"`
}
