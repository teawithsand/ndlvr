package testing_test

import livr "github.com/teawithsand/livr4go"

func RunE2ETest() {
	var rules livr.RulesMap

	opt := livr.Options{
		Rules:             rules,
		ValidationFactory: livr.MakeBuiltinFactory(),
	}
	_ = opt
}
