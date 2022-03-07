package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeEqVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			equalTo, err := bctx.ArgumentParser.ParsePrimitiveValue(
				bctx.Ctx,
				bctx.Data.Argument,
			)
			if err != nil {
				return
			}

			if !bctx.OPs.Eq(equalTo, fieldValue) {
				err = ndlvr.MakeLIVRError("input is not equal to expected value", "NOT_EQUAL")
				return
			}
			return
		})

	vf = ndlvr.WrapNamed("eq", vf)
	return
}

func makeOneOfVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			oneOf, err := bctx.ArgumentParser.ParseListValue(
				bctx.Ctx,
				bctx.Data.Argument,
			)
			if err != nil {
				return
			}

			var foundMatch bool
			for i := 0; i < oneOf.Len(); i++ {
				var eq, pv value.Value

				eq, err = oneOf.GetIndex(i)
				if err != nil {
					return
				}

				pv, err = value.ExpectPrimitiveValue(eq)
				if err != nil {
					return
				}
				if bctx.OPs.Eq(pv, fieldValue) {
					foundMatch = true
					break
				}
			}

			if !foundMatch {
				err = ndlvr.MakeLIVRError("input is not equal to any of possible values", "NOT_ONE_OF")
				return
			}

			return
		})

	vf = ndlvr.WrapNamed("one_of", vf)
	return
}
