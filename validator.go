package livr

import (
	"context"
)

type Options struct {
	Rules             RulesSource
	ValidationFactory ValidationFactory

	// Ignores type juggling when it comes to comparing stuff.
	// Things like max_length stop working on numbers.
	//
	// All lengths may not be floats.
	//
	// In general nice stuff, but incompatible with LIVR standard, so disabled by default.
	// Think of it as something like JS strict mode.
	//
	// Also: for now it's NIY in some places
	//
	// Also: use structures instead of maps, they take care of some problems already.
	IgnoreTypeJuggling bool
}

type Validator interface {
	Validate(ctx context.Context, value Value) (err error)
}

type validatorImpl struct {
	validations []Validation
}

func (v *validatorImpl) Validate(ctx context.Context, value Value) (err error) {
	var bag ErrorBag
	for _, validation := range v.validations {
		err = validation.Validate(ctx, value)
		if err != nil {
			bag.Errors = append(bag.Errors, err)
		}
	}

	if bag.IsEmpty() {
		err = nil
	} else {
		err = &bag
	}
	return
}
