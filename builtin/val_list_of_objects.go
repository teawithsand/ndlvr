package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeListOfObjectsValidator() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.ValidationFactoryFunc(func(bctx ndlvr.ValidationBuildContext) (val ndlvr.Validation, err error) {
		rules, err := bctx.ArgumentParser.ParseTopRulesSource(bctx.Ctx, bctx.Data.Argument)
		if err != nil {
			return
		}

		engine, err := bctx.Options.NewEngine(bctx.Ctx, rules)
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

				listLen := list.Len()
				for i := 0; i < listLen; i++ {
					var valueAtIndex value.Value
					valueAtIndex, err = list.GetIndex(i)
					if err != nil {
						return
					}

					err = engine.Validate(bctx.Ctx, valueAtIndex)
					if err != nil {
						return
					}
				}

				return
			},
		).BuildValidation(bctx)
		return
	})
	vf = ndlvr.WrapNamed("list_of_objects", vf)
	return
}
