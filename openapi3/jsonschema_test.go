package openapi3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SchemaRef(t *testing.T) {
	loader := NewLoader()
	loader.IsExternalRefsAllowed = true

	var err error
	sf := NewSchemaRef("testdata/jsonschema.json", nil)
	err = loader.ResolveSchemaRef(sf)
	require.NoError(t, err)
	require.NotEmpty(t, sf.Value.Properties["data"].Value.Properties)
}
