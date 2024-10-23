package parser

import (
	"strings"

	"github.com/jghiloni/ogen"
	"github.com/jghiloni/ogen/jsonschema"
)

type componentsResolver struct {
	components map[string]*ogen.Schema
	root       *jsonschema.RootResolver
}

func (c componentsResolver) ResolveReference(ref string) (*jsonschema.RawSchema, error) {
	const prefix = "#/components/schemas/"
	if strings.HasPrefix(ref, prefix) {
		name := strings.TrimPrefix(ref, prefix)
		s, ok := c.components[name]
		if ok {
			return s.ToJSONSchema(), nil
		}
	}
	return c.root.ResolveReference(ref)
}
