package value

import (
	"reflect"
)

func MustWrap(data interface{}) (v Value) {
	v, err := Wrap(data)
	if err != nil {
		panic(err)
	}
	return
}

// Util function, which converts go native type to Value.
func Wrap(data interface{}) (v Value, err error) {
	if data == nil {
		v = nil
		return
	}

	switch tdata := data.(type) {
	case Value:
		v = tdata
		return
	case string:
		v = &PrimitiveValue{tdata}
		return
	case float64:
		v = &PrimitiveValue{tdata}
		return
	case int:
		v = &PrimitiveValue{tdata}
		return
	case reflect.Value: // wrapping reflect is nono
		err = &InvalidValueError{
			Data: data,
		}
		return
	// TODO(teawithsand): add more primitive types here
	default:
		refVal := reflect.ValueOf(data)
		innerRefVal := refVal
		for innerRefVal.Kind() == reflect.Ptr {
			if innerRefVal.IsNil() {
				v = nil
				return
			}
			innerRefVal = innerRefVal.Elem()
		}

		// not nil, so we can operate on it
		// nil maps are not allowed, since we can't assign to them
		// we could construct one
		// but it's not really what we want from Wrap function

		if innerRefVal.Kind() == reflect.Map || innerRefVal.Kind() == reflect.Struct {
			v = &mutableReflectKeyedValue{
				reflectKeyedValue: reflectKeyedValue{
					val: refVal,
				},
			}
			return
		} else if innerRefVal.Kind() == reflect.Slice || innerRefVal.Kind() == reflect.Array {
			v = &defaultMutableListValue{
				defaultListValue: defaultListValue{
					val: refVal,
				},
			}
			return
		} else if refVal.Kind() == reflect.Ptr {
			return Wrap(innerRefVal.Interface())
		} else {
			err = &InvalidValueError{
				Data: data,
			}
			return
		}
	}
}
