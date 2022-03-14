package ndlvr

import (
	"context"

	"github.com/teawithsand/ndlvr/value"
)

// Defines what this validation is running one
type ValidationTarget struct {
	FieldName        string // ignored when any other is set
	IsListValue      bool   // ignored when FT is set
	FunctionalTarget func(v value.Value, recv func(child value.Value) (err error)) (err error)
}

// func(vt *ValidationTarget) GetIterator(v value.Value, recv func(child value.Value) (err error)) (err error)

type ValidationBuildData struct {
	Target ValidationTarget

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

type ValidationAsFactory func(bctx ValidationBuildContext, value value.Value) (err error)

func (f ValidationAsFactory) BuildValidation(bctx ValidationBuildContext) (val Validation, err error) {
	val = ValidationFunc(func(ctx context.Context, value value.Value) (err error) {
		bctx.Ctx = ctx
		return f(bctx, value)
	})
	return
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
	// TODO(teawithsand): make this wrapper return errors, which contain validator name
	return &namedValidationFactory{
		Name:    name,
		Factory: factory,
	}
}

// SimpleFieldValidation, which accesses value of field passed in build data.
func SimpleFieldValidation(
	require bool,
	inner func(bctx ValidationBuildContext, parentValue value.Value, fieldValue value.Value) (err error),
) ValidationFactory {
	return ValidationAsFactory(func(bctx ValidationBuildContext, vv value.Value) (err error) {
		// We operate on list or something, so parent value is not accessible
		// TODO(teawithsand): make list valid parent value(?) instead of doing this no-parent hack
		if bctx.Data.Target.FunctionalTarget != nil {
			selector := bctx.Data.Target.FunctionalTarget

			err = selector(vv, func(child value.Value) (err error) {
				err = inner(bctx, vv, child)
				return
			})
			if err != nil {
				return
			}
		} else if bctx.Data.Target.IsListValue {
			var listValue value.ListValue
			listValue, err = value.ExpectListValue(vv)
			if err != nil {
				return
			}

			length := listValue.Len()
			for i := 0; i < length; i++ {
				var nth value.Value
				nth, err = listValue.GetIndex(i)
				if err != nil {
					return
				}

				err = inner(bctx, vv, nth)
				if err != nil {
					return
				}
			}
		} else {
			var fieldValue value.Value
			fieldValue, err = value.ExpectKeyedValueField(vv, bctx.Data.Target.FieldName, require)
			if err != nil {
				return
			}

			err = inner(bctx, vv, fieldValue)
		}
		return
	})
}
