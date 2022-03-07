package ndlvr

import "fmt"

type RuleParseError struct {
	FieldName string
	Rule      interface{}
}

func (e *RuleParseError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("ndlvr: Rule parse filed for field: %s", e.FieldName)
}

type ValidationNameMismatchError struct {
	Name         string
	ExpectedName string
}

func (e *ValidationNameMismatchError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("ndlvr: given factory supports only validation '%s' but '%s' was provided", e.Name, e.ExpectedName)
}

type ValidationCreateError struct {
	Msg string
}

func (e *ValidationCreateError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("ndlvr: filed to create validation: %s", e.Msg)
}

type ValidationNotSupportedError struct {
	Name string
}

func (e *ValidationNotSupportedError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("ndlvr: validation '%s' is not supported by this factory", e.Name)
}
