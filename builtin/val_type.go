package builtin

import (
	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeIsStringVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			_, err = value.ExpectStringValue(fieldValue)
			if err != nil {
				return
			}
			return
		})

	vf = ndlvr.WrapNamed("is_string", vf)
	return
}
