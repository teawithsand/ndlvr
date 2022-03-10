package ndlvr

import (
	"context"

	"github.com/teawithsand/ndlvr/value"
)

// Validator is just like engine, but operates on interface{} type instead of ndlvr's value.
type Validator interface {
	Validate(ctx context.Context, val interface{}) (err error)
}

type DefaultValidator struct {
	Engine  Engine        // Required
	Wrapper value.Wrapper // Defaults to DefaultWrapper
}

func (dv *DefaultValidator) Validate(ctx context.Context, val interface{}) (err error) {
	wrapper := dv.Wrapper
	if wrapper == nil {
		wrapper = &value.DefaultWrapper{}
	}

	v, err := wrapper.Wrap(val)
	if err != nil {
		return
	}

	return dv.Engine.Validate(ctx, v)
}
