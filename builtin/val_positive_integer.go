package builtin

import (
	"math"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makePositiveIntegerVF() (vf ndlvr.ValidationFactory) {
	vf = ndlvr.SimpleFieldValidation(
		true,
		func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
			pv, err := value.ExpectPrimitiveValue(fieldValue)
			if err != nil {
				return
			}

			retErr := ndlvr.MakeNDLVRError("Value must be positive integer", "NOT_POSITIVE")
			switch tv := pv.RawUnpointered().(type) {
			case uint8:
				if tv <= 0 {
					err = retErr
					return
				}
			case uint16:
				if tv <= 0 {
					err = retErr
					return
				}
			case uint32:
				if tv <= 0 {
					err = retErr
					return
				}
			case uint64:
				if tv <= 0 {
					err = retErr
					return
				}
			case int8:
				if tv <= 0 {
					err = retErr
					return
				}
			case int16:
				if tv <= 0 {
					err = retErr
					return
				}
			case int32:
				if tv <= 0 {
					err = retErr
					return
				}
			case int64:
				if tv <= 0 {
					err = retErr
					return
				}
			case float32:
				if tv <= 0 || math.IsInf(float64(tv), 0) || math.IsNaN(float64(tv)) {
					err = retErr
					return
				}
				if math.Round(float64(tv)) != float64(tv) {
					err = ndlvr.MakeNDLVRError("Value must be positive integer", "NOT_INTEGER")
					return
				}
			case float64:
				if tv <= 0 || math.IsInf(tv, 0) || math.IsNaN(tv) {
					err = retErr
					return
				}
				if math.Round(float64(tv)) != float64(tv) {
					err = ndlvr.MakeNDLVRError("Value must be positive integer", "NOT_INTEGER")
					return
				}
			// TODO(teawithsand): use big.Int here?
			// json.Number.String()
			// case json.Number:
			default:
				err = value.ErrExpectFiled
				return
			}

			return
		})

	vf = ndlvr.WrapNamed("positive_integer", vf)
	return
}
