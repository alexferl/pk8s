package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"sigs.k8s.io/yaml"
)

type APIObject struct {
	v any
}

// JSON returns the JSON encoding of the API object.
func (a *APIObject) JSON() ([]byte, error) {
	return json.Marshal(a.v)
}

// YAML returns the YAML encoding of the API object.
func (a *APIObject) YAML() ([]byte, error) {
	return yaml.Marshal(a.v)
}

// String returns a YAML string of the API object.
func (a *APIObject) String() string {
	b, err := yaml.Marshal(a.v)
	if err != nil {
		log.Fatal().Msgf("failed marshalling API object: %v", err)
	}
	return string(b)
}
