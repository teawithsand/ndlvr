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

	return fmt.Sprintf("livr: given factory supports validation '%s' but '%s' was provided", e.Name, e.ExpectedName)
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