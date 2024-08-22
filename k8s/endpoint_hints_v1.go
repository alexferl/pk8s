// Code generated; DO NOT EDIT.

package k8s

// EndpointHintsV1 EndpointHints provides hints describing how an endpoint should be consumed.
type EndpointHintsV1 struct {
	// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
	ForZones []ForZoneV1 `json:"forZones,omitempty"`
}
