package livr

import (
	"context"

	"github.com/teawithsand/ndlvr/value"
)

type Engine interface {
	Validate(ctx context.Context, parentValue value.Value) (err error)
}

type engineImpl struct {
	validations []Validation
}

func (v *engineImpl) Validate(ctx context.Context, parentValue value.Value) (err error) {
	var bag ErrorBag
	for _, validation := range v.validations {
		err = validation.Validate(ctx, parentValue)
		if err != nil {
			// panic(err)
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
