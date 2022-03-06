package value

import (
	"errors"
	"reflect"
)

var ErrNotStringable = errors.New("livr/value: specified value can't be converted to string")

type NotAssignableValueError struct {
	To    reflect.Type
	Value Value
}

func (e *NotAssignableValueError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "livr/value: given value can't be assigned to given field"
}

type NotSettableValueError struct {
	Data interface{}
}

func (e *NotSettableValueError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "livr/value: given value is not addressable and can't be assigned. If this is struct consider passing pointer instead of struct."
}

type NoFieldError struct {
	Value Value
	Field interface{}
}

func (e *NoFieldError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "livr/value: field with given name does not exist"
}

type InvalidValueError struct {
	Data interface{}
}

func (e *InvalidValueError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return "livr/value: given data can't be wrapped as livr value"
}
