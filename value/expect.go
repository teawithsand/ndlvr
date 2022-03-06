package value

import (
	"errors"
)

var ErrExpectFiled = errors.New("livr/value: some kind of expect filed")

func ExpectKeyedValueField(kv Value, fieldName interface{}, required bool) (value Value, err error) {
	skv, ok := value.(KeyedValue)
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
