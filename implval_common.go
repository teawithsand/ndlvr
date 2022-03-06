package livr

import (
	"errors"

	"github.com/teawithsand/livr4go/value"
)

var ErrEmpty = errors.New("livr: field must not be empty")

func makeRequiredVF() (vf ValidationFactory) {
	vf = ValidationAsFactory(func(bctx ValidationBuildContext, vv value.Value) (err error) {
		_, err = value.ExpectKeyedValueField(vv, bctx.Data.FieldName, true)
		if err != nil {
			return
		}
		return
	})

	vf = WrapNamed("required", vf)
	return
}

func makeNotEmptyVF() (vf ValidationFactory) {
	vf = ValidationAsFactory(func(bctx ValidationBuildContext, vv value.Value) (err error) {
		fieldValue, err := value.ExpectKeyedValueField(vv, bctx.Data.FieldName, true)
		if err != nil {
			return
		}

		if bctx.OPs.IsEmpty(fieldValue) {
			err = ErrEmpty
			return
		}
		return
	})

	vf = WrapNamed("required", vf)
	return
}
