package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeListOfValidator() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.ValidationFactoryFunc(func(bctx ndlvr.ValidationBuildContext) (val ndlvr.Validation, err error) {
		rules, err := bctx.ArgumentParser.ParseRulesSource(bctx.Ctx, bctx.Data.Argument)
		if err != nil {
			return
		}

		engine, err := bctx.Options.NewEngineWithTarget(bctx.Ctx, rules, ndlvr.ValidationTarget{
			IsListValue: true,
		})
		if err != nil {
			return
		}

		val, err = ndlvr.SimpleFieldValidation(
			false,
			func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
				if fieldValue == nil { // value is not required, use required validator for that
					return
				}

				// not-list should cause error here, not for each child
				list, err := value.ExpectListValue(fieldValue)
				if err != nil {
					return
				}

				err = engine.Validate(bctx.Ctx, list)
				if err != nil {
					return
				}

				return
			},
		).BuildValidation(bctx)
		return
	})
	vf = ndlvr.WrapNamed("list_of", vf)
	return
}
