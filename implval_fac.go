package livr

// MakeBuiltinFactory creates ValidationFactoryRegistry, which contains all builtin validations.
func MakeBuiltinFactory() ValidationFactoryRegistry {
	fac := make(ValidationFactoryRegistry)
	// general "empty" stuff
	fac.MustPut("required", makeRequiredVF())
	fac.MustPut("not_empty", makeNotEmptyVF())

	// object/data structure stuff
	// fac.MustPut("any_object", makeAnyObjectVF())
	// fac.MustPut("not_empty_list", makeNotEmptyListVF())

	// fac.MustPut("max_length", makeLengthValidator("max_length", maxIncLVK))
	// fac.MustPut("min_length", makeLengthValidator("min_length", minIncLVK))
	return fac
}
