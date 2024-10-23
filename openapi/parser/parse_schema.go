package parser

import (
	"github.com/go-faster/errors"

	"github.com/jghiloni/ogen"
	"github.com/jghiloni/ogen/jsonpointer"
	"github.com/jghiloni/ogen/jsonschema"
)

func (p *parser) parseSchema(schema *ogen.Schema, ctx *jsonpointer.ResolveCtx) (*jsonschema.Schema, error) {
	s, err := p.schemaParser.Parse(schema.ToJSONSchema(), ctx)
	if err != nil {
		return nil, errors.Wrap(err, "parse schema")
	}
	return s, nil
}
