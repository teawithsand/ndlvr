package ndlvr

// MakeBuiltinFactory creates ValidationFactoryRegistry, which contains all builtin validations.
func MakeBuiltinFactory() ValidationFactoryRegistry {
	fac := make(ValidationFactoryRegistry)
	// general "empty" stuff
	fac.MustPut("required", makeRequiredVF())
	fac.MustPut("not_empty", makeNotEmptyVF())

	// string stuff
	fac.MustPut("is_string", makeIsStringVF())
	fac.MustPut("min_length", makeMinLengthVF())
	fac.MustPut("max_length", makeMaxLengthVF())

	fac.MustPut("eq", makeEqVF())
	fac.MustPut("one_of", makeOneOfVF())

	// equal stuff
	return fac
}
