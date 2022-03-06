package livr

import "fmt"

type RuleParseError struct {
	FieldName string
	Rule      interface{}
}

func (e *RuleParseError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("livr: Rule parse filed for field: %s", e.FieldName)
}

type ValidationNameMismatchError struct {
	Name         string
	ExpectedName string
}

func (e *ValidationNameMismatchError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("livr: given factory supports only validation '%s' but '%s' was provided", e.Name, e.ExpectedName)
}

type ValidationCreateError struct {
	Msg string
}

func (e *ValidationCreateError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("livr: filed to create validation: %s", e.Msg)
}

type ValidationNotSupportedError struct {
	Name string
}

func (e *ValidationNotSupportedError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("livr: validation '%s' is not supported by this factory", e.Name)
}
