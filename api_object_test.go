package pk8s

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestAPIObject_JSON(t *testing.T) {
	data := map[string]interface{}{
		"name":  "test",
		"value": 42,
	}

	apiObj := &APIObject{v: data}

	jsonData, err := apiObj.JSON()
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	assert.NoError(t, err)

	expected := map[string]interface{}{
		"name":  "test",
		"value": float64(42),
	}

	assert.Equal(t, expected, result)
}

func TestAPIObject_YAML(t *testing.T) {
	data := map[string]interface{}{
		"name":  "test",
		"value": 42,
	}

	apiObj := &APIObject{v: data}

	yamlData, err := apiObj.YAML()
	assert.NoError(t, err)

	var result map[string]interface{}
	err = yaml.Unmarshal(yamlData, &result)
	assert.NoError(t, err)

	expected := map[string]interface{}{
		"name":  "test",
		"value": float64(42),
	}

	assert.Equal(t, expected, result)
}

func TestAPIObject_String(t *testing.T) {
	data := map[string]interface{}{
		"name":  "test",
		"value": 42,
	}

	apiObj := &APIObject{v: data}

	yamlData := apiObj.String()

	assert.Equal(t, "name: test\nvalue: 42\n", yamlData)
}
