package livr

import (
	"context"

	"github.com/teawithsand/livr4go/value"
)

type Validator interface {
	Validate(ctx context.Context, v value.Value) (err error)
}

type validatorImpl struct {
	validations []Validation
}

func (v *validatorImpl) Validate(ctx context.Context, validatedValue value.Value) (err error) {
	var bag ErrorBag
	for _, validation := range v.validations {
		err = validation.Validate(ctx, validatedValue)
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
