package value

import (
	"reflect"
)

// reflectKeyedValue wraps any map or struct into value.
type reflectKeyedValue struct {
	val reflect.Value
}

var _ KeyedValue = &reflectKeyedValue{}

// returns element of pointer if any.
func (rkv *reflectKeyedValue) getInnerValue() (res reflect.Value) {
	res = rkv.val
	for res.Kind() == reflect.Ptr {
		res = res.Elem()
	}
	return
}

func (rkv *reflectKeyedValue) Raw() interface{} {
	return rkv.val.Interface()
}

func isReflectZero(v reflect.Value) bool {
	return v == reflect.Value{}
}

// Panics when no such field.
// Must not return nil in that case.
//
// Returns nil value if field was not found.
func (rkv *reflectKeyedValue) GetField(key interface{}) (res Value, err error) {
	iv := rkv.getInnerValue()

	if iv.Kind() == reflect.Map {
		v := iv.MapIndex(reflect.ValueOf(key))
		if isReflectZero(v) {
			return nil, nil
		}

		res, err = Wrap(v.Interface())
		return
	} else {
		skey, ok := key.(string)
		if !ok {
			return
		}
		v := iv.FieldByName(skey)
		if isReflectZero(v) {
			return nil, nil
		}

		res, err = Wrap(v.Interface())
		return
	}
}

// Returns true if given field exists in value, false otherwise.
func (rkv *reflectKeyedValue) HasField(key interface{}) bool {
	iv := rkv.getInnerValue()

	if iv.Kind() == reflect.Map {
		return !isReflectZero(iv.MapIndex(reflect.ValueOf(key)))
	} else {
		sname, ok := key.(string)
		if !ok {
			return false
		}
		return !isReflectZero(iv.FieldByName(sname))
	}
}

// Iteration must stop when non-nil error is returned.
// This error must be returned from top-level function.
//
// Note: field name yielded here is not value but primitive go type, like string or int.
func (rkv *reflectKeyedValue) ListFields(recv func(name interface{}) (err error)) (err error) {
	iv := rkv.getInnerValue()

	if iv.Kind() == reflect.Map {
		for _, k := range iv.MapKeys() {
			err = recv(k.Interface())
			if err != nil {
				return
			}
		}
	} else {
		sz := iv.NumField()
		for i := 0; i < sz; i++ {
			err = recv(iv.Field(i).Interface())
			if err != nil {
				return
			}
		}
	}

	return
}

// Returns number of fields.
func (rkv *reflectKeyedValue) Len() int {
	iv := rkv.getInnerValue()

	if iv.Kind() == reflect.Map {
		return iv.Len()
	} else {
		return iv.NumField()
	}
}

type mutableReflectKeyedValue struct {
	reflectKeyedValue
}

func (mrkv *mutableReflectKeyedValue) getReflectField(key interface{}) reflect.Value {
	iv := mrkv.getInnerValue()

	if iv.Kind() == reflect.Map {
		v := iv.MapIndex(reflect.ValueOf(key))
		if isReflectZero(v) {
			return reflect.Value{}
		}

		return v
	} else {
		skey, ok := key.(string)
		if !ok {
			return reflect.Value{}
		}
		return iv.FieldByName(skey)
	}
}

func (mrkv *mutableReflectKeyedValue) IsAssignable(key interface{}, value Value) bool {
	if !mrkv.HasField(key) {
		return false
	}

	iv := mrkv.getInnerValue()
	if iv.Kind() == reflect.Map {
		return isAssignable(iv.Type().Elem(), value)
	} else {
		fieldType := mrkv.getReflectField(key).Type()

		return isAssignable(fieldType, value)
	}
}

func (mrkv *mutableReflectKeyedValue) SetField(key interface{}, value Value) (err error) {
	if !mrkv.HasField(key) {
		err = &NoFieldError{
			Value: mrkv,
			Field: key,
		}
		return
	}

	iv := mrkv.getInnerValue()
	if iv.Kind() == reflect.Map {
		err = assignMap(iv, key, value)
		if err != nil {
			return
		}
		return
	} else {
		fieldRef := mrkv.getReflectField(key)
		err = assignValue(fieldRef, value)
		if err != nil {
			return
		}
	}

	return
}
