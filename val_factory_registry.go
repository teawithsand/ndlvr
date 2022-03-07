package ndlvr

import (
	"fmt"
)

// ValidationFactoryRegistry is registry of factories, which always directs same validation name to same factory.
type ValidationFactoryRegistry map[string]ValidationFactory

// MustPut puts ValidationFactory with given name.
// Fails if one with given name is already set.
// Registry must not be nil map.
func (reg ValidationFactoryRegistry) MustPut(name string, vfac ValidationFactory) {
	_, ok := reg[name]
	if ok {
		panic(fmt.Errorf("ndlvr: validation registry put filed: validation '%s' is already set", name))
	}
	reg[name] = vfac
}

// Sets ValidationFactory with given name.
// Overrides one if was set already.
// Registry must not be nil map.
func (reg ValidationFactoryRegistry) Set(name string, vfac ValidationFactory) {
	reg[name] = vfac
}

func (reg ValidationFactoryRegistry) BuildValidation(bctx ValidationBuildContext) (val Validation, err error) {
	if reg == nil {
		err = &ValidationNotSupportedError{
			Name: bctx.Data.ValidationName,
		}
		return
	}

	arg, ok := reg[bctx.Data.ValidationName]
	if !ok {
		err = &ValidationNotSupportedError{
			Name: bctx.Data.ValidationName,
		}
		return
	}

	return arg.BuildValidation(bctx)
}
