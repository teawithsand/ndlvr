package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeRequiredVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(true, func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
		return
	})
	vf = ndlvr.WrapNamed("required", vf)
	return
}
