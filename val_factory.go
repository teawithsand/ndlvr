package livr

import (
	"context"

	"github.com/teawithsand/livr4go/value"
)

type ValidationBuildData struct {
	FieldName      string
	ValidationName string
	Argument       interface{}
}

type ValidationBuildContext struct {
	Ctx context.Context

	Options *Options
	Parser  *Parser

	OPs            value.OPs
	ArgumentParser ArgumentParser

	Data ValidationBuildData
}

type ValidationFactory interface {
	BuildValidation(bctx ValidationBuildContext) (val Validation, err error)
}

type ValidationFactoryFunc func(bctx ValidationBuildContext) (val Validation, err error)

func (f ValidationFactoryFunc) BuildValidation(bctx ValidationBuildContext) (val Validation, err error) {
	return f(bctx)
}

// ValidationFactory, which returns an error when name provided in data does not match given one.
type namedValidationFactory struct {
	Name    string
	Factory ValidationFactory
}

func (nvf *namedValidationFactory) BuildValidation(bctx ValidationBuildContext) (val Validation, err error) {
	if bctx.Data.ValidationName != nvf.Name {
		err = &ValidationNameMismatchError{
			Name:         bctx.Data.ValidationName,
			ExpectedName: nvf.Name,
		}
		return
	}

	return nvf.Factory.BuildValidation(bctx)
}

// Wraps specified factory, in one which verifies validaton name and returns error on mismatch.
func WrapNamed(name string, factory ValidationFactory) ValidationFactory {
	return &namedValidationFactory{
		Name:    name,
		Factory: factory,
	}
}