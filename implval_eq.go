package ndlvr

import "github.com/teawithsand/ndlvr/value"

func makeEqVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			equalTo, err := bctx.ArgumentParser.ParsePrimitiveValue(
				bctx.Ctx,
				bctx.Data.Argument,
			)
			if err != nil {
				return
			}

			if !bctx.OPs.Eq(equalTo, fieldValue) {
				err = MakeLIVRError("input is not equal to expected value", "NOT_EQUAL")
				return
			}
			return
		})

	vf = WrapNamed("eq", vf)
	return
}

func makeOneOfVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
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
				err = MakeLIVRError("input is not equal to any of possible values", "NOT_ONE_OF")
				return
			}

			return
		})

	vf = WrapNamed("one_of", vf)
	return
}
