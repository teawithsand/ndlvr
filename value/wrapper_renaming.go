package value

// Wrapper, which translates name specified to new name.
type FieldAliasWrapper interface {
	Wrapper
	GetAlias(v KeyedValue, name interface{}) (alias interface{}, err error)
}
