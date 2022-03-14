package builtin

import "github.com/teawithsand/ndlvr"

// MakeBuiltinFactory creates ValidationFactoryRegistry, which contains all builtin validations.
func MakeBuiltinFactory() ndlvr.ValidationFactoryRegistry {
	fac := make(ndlvr.ValidationFactoryRegistry)
	// general "empty" stuff
	fac.MustPut("required", makeRequiredVF())
	fac.MustPut("not_empty", makeNotEmptyVF())

	// string stuff
	fac.MustPut("is_string", makeIsStringVF())
	fac.MustPut("min_length", makeMinLengthVF())
	fac.MustPut("max_length", makeMaxLengthVF())
	fac.MustPut("email", makeEmailVF(true))

	// equal stuff
	fac.MustPut("eq", makeEqVF())
	fac.MustPut("one_of", makeOneOfVF())
	fac.MustPut("like", makeLikeValidator(false))

	// embedded structures stuff: list/objects
	fac.MustPut("list_of", makeListOfValidator())

	return fac
}
