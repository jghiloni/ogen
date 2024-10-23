package openapi

import (
	"github.com/jghiloni/ogen/jsonschema"
	"github.com/jghiloni/ogen/location"
)

// Example is an OpenAPI Example.
type Example struct {
	Ref Ref

	Summary       string
	Description   string
	Value         jsonschema.Example
	ExternalValue string

	location.Pointer `json:"-" yaml:"-"`
}
