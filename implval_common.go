package livr

import (
	"github.com/teawithsand/livr4go/value"
)

func makeRequiredVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(true, func(bctx ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
		return
	})
	vf = WrapNamed("required", vf)
	return
}

func makeNotEmptyVF() (vf ValidationFactory) {
	vf = SimpleFieldValidation(
		true,
		func(bctx ValidationBuildContext, vv, fieldValue value.Value) (err error) {
			if bctx.OPs.IsEmpty(fieldValue) {
				err = MakeLIVRError("Field must not be empty", "NOT_EMPTY")
				return
			}
			return
		})

	vf = WrapNamed("required", vf)
	return
}
