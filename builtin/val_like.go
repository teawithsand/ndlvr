package builtin

import (
	"regexp"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func makeLikeValidator(posix bool) (vf ndlvr.ValidationFactory) {
	vf = ndlvr.ValidationFactoryFunc(func(bctx ndlvr.ValidationBuildContext) (val ndlvr.Validation, err error) {
		regex, err := bctx.ArgumentParser.ParsePrimitiveValue(
			bctx.Ctx,
			bctx.Data.Argument,
		)
		if err != nil {
			return
		}

		regexString, err := value.ExpectStringValue(regex)
		if err != nil {
			return
		}

		var expr *regexp.Regexp
		if posix {
			expr, err = regexp.CompilePOSIX(regexString)
		} else {
			expr, err = regexp.Compile(regexString)
		}
		if err != nil {
			return
		}

		val, err = ndlvr.SimpleFieldValidation(
			false,
			func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
				if fieldValue == nil { // value is not required, use required validator for that
					return
				}

				testString, err := value.ExpectStringValue(fieldValue)
				if err != nil {
					return
				}

				if !expr.MatchString(testString) {
					err = ndlvr.MakeNDLVRError("input does not match regex", "NOT_LIKE")
					return
				}

				return
			},
		).BuildValidation(bctx)
		return
	})
	vf = ndlvr.WrapNamed("like", vf)
	return
}
