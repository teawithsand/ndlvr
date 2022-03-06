package livr

type Options struct {
	Rules             RulesSource
	ValidationFactory ValidationFactory

	// Ignores type juggling when it comes to comparing stuff.
	// Things like max_length stop working on not numbers.
	//
	// All lengths may not be floats.
	//
	// In general nice stuff, but incompatible with LIVR standard, so disabled by default.
	// Think of it as something like JS strict mode.
	//
	// Also: for now it's NIY in some places
	//
	// Also: use structures instead of maps, they take care of some problems already.
	IgnoreTypeJuggling bool
}
