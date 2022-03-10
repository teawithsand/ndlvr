package builtin

import (
	"net/mail"
	"regexp"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/value"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// just to make sure it will work
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func makeEmailVF(strict bool) (vf ndlvr.ValidationFactory) {
	vf = ndlvr.ValidationFactoryFunc(func(bctx ndlvr.ValidationBuildContext) (val ndlvr.Validation, err error) {
		val, err = ndlvr.SimpleFieldValidation(
			false,
			func(bctx ndlvr.ValidationBuildContext, parentValue, fieldValue value.Value) (err error) {
				if fieldValue == nil { // value is not required, use required validator for that
					return
				}

				emailString, err := value.ExpectStringValue(fieldValue)
				if err != nil {
					return
				}

				// length validator is here, since it looks like valid allows arbitrary long email
				if !valid(emailString) || len(emailString) > 256 || (!strict && !emailRegex.MatchString(emailString)) {
					err = ndlvr.MakeNDLVRError("email is not valid", "INVALID_EMAIL")
					return
				}

				return
			},
		).BuildValidation(bctx)
		return
	})
	vf = ndlvr.WrapNamed("email", vf)
	return
}
