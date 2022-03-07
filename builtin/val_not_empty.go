package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeNotEmptyVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, vv, fieldValue value.Value) (err error) {
			if bctx.OPs.IsEmpty(fieldValue) {
				err = ndlvr.MakeNDLVRError("Field must not be empty", "NOT_EMPTY")
				return
			}
			return
		})

	vf = ndlvr.WrapNamed("required", vf)
	return
}
