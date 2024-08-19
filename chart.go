package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/alexferl/pk8s/k8s"
)

type Chart struct {
	Name      string
	Namespace string
	Params    *ChartParams

	apiObjects []APIObject
}

type ChartParams struct {
	// DisableNameHashes removes the generated stable hash that is suffixed to the resource name.
	// Optional. Default: false.
	DisableNameHashes bool

	// Labels to apply to all the resources in this chart.
	// Optional. Default: nil.
	Labels *map[string]string

	// Namespace to apply to all the resources in this chart. This will only apply to objects that don't have a
	// namespace already defined on them.
	// Optional. Default: nil. (usually means "default" in Kubernetes)
	Namespace *string
}

func NewChart(s *Stack, name string, params *ChartParams) *Chart {
	if params == nil {
		params = &ChartParams{}
	}
	c := &Chart{
		Name:       name,
		Params:     params,
		apiObjects: make([]APIObject, 0),
	}
	s.charts = append(s.charts, c)
	return s.charts[len(s.charts)-1]
}

// Append appends k8s objects.
func (c *Chart) Append(objs ...any) {
	for _, i := range objs {
		obj := APIObject{v: i}

		k8s.SetApiVersionAndKind(i)

		b, err := obj.JSON()
		if err != nil {
			log.Fatal().Msgf("failed hashing API object: %v", err)
		}

		h := sha256.New()
		h.Write(b)
		bs := h.Sum(nil)

		name := c.Name
		if !c.Params.DisableNameHashes {
			name = fmt.Sprintf("%s-%x", c.Name, bs[:4])
		}
		err = modifyMetadataField(i, "Name", &name)
		if err != nil {
			log.Fatal().Msgf("failed to modify Name metadata field: %v", err)
		}

		if c.Params.Namespace != nil {
			err = modifyMetadataField(i, "Namespace", c.Params.Namespace)
			if err != nil {
				log.Fatal().Msgf("failed to modify Namespace metadata field: %v", err)
			}
		}

		if c.Params.Labels != nil {
			err = modifyMetadataField(i, "Labels", c.Params.Labels)
			if err != nil {
				log.Fatal().Msgf("failed to modify Labels metadata field: %v", err)
			}
		}

		c.apiObjects = append(c.apiObjects, obj)
	}
}

func (c *Chart) String() string {
	var sb strings.Builder
	sb.WriteString("---\n")
	for j, obj := range c.apiObjects {
		b, err := obj.YAML()
		if err != nil {
			log.Fatal().Msgf("failed converting to YAML: %v", err)
		}

		sb.Write(b)

		// Check if this is the last element
		if !(j == len(c.apiObjects)-1) {
			sb.WriteString("---\n")
		}
	}

	return sb.String()
}

func modifyMetadataField(s interface{}, fieldName string, newValue any) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct or a pointer to a struct")
	}

	metadataField := val.FieldByName("Metadata")
	if !metadataField.IsValid() {
		return fmt.Errorf("no Metadata field found")
	}

	if metadataField.IsNil() {
		metadataField.Set(reflect.New(metadataField.Type().Elem()))
	}

	metadata := metadataField.Elem()
	field := metadata.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no %s field found in Metadata", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set %s field", fieldName)
	}

	if fieldName == "Namespace" && !isZeroValue(field) {
		return nil
	}

	newValueVal := reflect.ValueOf(newValue)
	if newValueVal.Type().ConvertibleTo(field.Type()) {
		field.Set(newValueVal.Convert(field.Type()))
	} else {
		return fmt.Errorf("cannot convert %v to %s", newValueVal.Type(), field.Type())
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
