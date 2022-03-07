package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeMinLengthVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
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
				err = ndlvr.MakeNDLVRError("input is too short", "TOO_SHORT")
				return
			}
			return
		})

	vf = ndlvr.WrapNamed("min_length", vf)
	return
}

func makeMaxLengthVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
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
				err = ndlvr.MakeNDLVRError("input is too short", "TOO_LONG")
				return
			}
			return
		})

	vf = ndlvr.WrapNamed("max_length", vf)
	return
}
