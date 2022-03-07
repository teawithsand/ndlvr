package livr

import "github.com/teawithsand/ndlvr/value"

func makeIsStringVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			_, err = value.ExpectStringValue(fieldValue)
			if err != nil {
				return
			}
			return
		})

	vf = WrapNamed("is_string", vf)
	return
}

func makeMinLengthVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			sz, err := bctx.ArgumentParser.ParseLen(
				bctx.Ctx,
				bctx.Data.Argument,
			)
			if err != nil {
				return
			}

			sv, err := value.ExpectStringValue(fieldValue)
			if err != nil {
				return
			}

			if len(sv) < sz {
				err = MakeLIVRError("input is too short", "TOO_SHORT")
				return
			}
			return
		})

	vf = WrapNamed("min_length", vf)
	return
}

func makeMaxLengthVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			sz, err := bctx.ArgumentParser.ParseLen(
				bctx.Ctx,
				bctx.Data.Argument,
			)
			if err != nil {
				return
			}

			sv, err := value.ExpectStringValue(fieldValue)
			if err != nil {
				return
			}

			if len(sv) > sz {
				err = MakeLIVRError("input is too short", "TOO_LONG")
				return
			}
			return
		})

	vf = WrapNamed("max_length", vf)
	return
}
