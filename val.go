package livr

import (
	"context"

	"github.com/teawithsand/livr4go/value"
)

// Validation validates specified key, but has access to whole value, so things like checking if fields are same work.
type Validation interface {
	// Note: validate MUST NOT modify value.
	Validate(ctx context.Context, parentValue value.Value) (err error)
}

// Modifier modifies value given.
// All modifiers are Validations.
// Some validations lik "to_lc" modify values.
//
// Unilike in most libraries, here validate and modify are separate stages to make code a bit cleanner.
type Modifier interface {
	Validation
	Modify(ctx context.Context, parentValue value.Value) (err error)
}

type ValidationFunc func(ctx context.Context, parentValue value.Value) (err error)

func (vf ValidationFunc) Validate(ctx context.Context, parentValue value.Value) (err error) {
	return vf(ctx, parentValue)
}
