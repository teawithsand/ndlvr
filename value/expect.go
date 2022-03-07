package value

import (
	"errors"
)

var ErrExpectFiled = errors.New("ndlvr/value: some kind of expect filed")

func ExpectKeyedValueField(kv Value, fieldName interface{}, required bool) (value Value, err error) {
	skv, ok := kv.(KeyedValue)
	if !ok {
		err = ErrExpectFiled
		return
	}

	value, err = skv.GetField(fieldName)
	if err != nil {
		return
	}

	if required && value == nil {
		err = ErrExpectFiled
		return
	}

	return
}

func ExpectPrimitiveValue(val Value) (pv *PrimitiveValue, err error) {
	if val == nil {
		err = ErrExpectFiled
		return
	}

	pv, ok := val.(*PrimitiveValue)
	if !ok {
		err = ErrExpectFiled
		return
	}
	return
}

func ExpectStringValue(val Value) (sv string, err error) {
	pv, err := ExpectPrimitiveValue(val)
	if err != nil {
		return
	}

	up := pv.RawUnpointered()
	sv, ok := up.(string)
	if !ok {
		err = ErrExpectFiled
		return
	}

	return
}
