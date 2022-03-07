package value

import "reflect"

func isAssignable(ty reflect.Type, val Value) (ok bool) {
	return reflect.TypeOf(val.Raw()).AssignableTo(ty) || reflect.PtrTo(reflect.TypeOf(val.Raw())).AssignableTo(ty)
}

// assignValue
func assignValue(fieldRef reflect.Value, value Value) (err error) {
	if !fieldRef.CanSet() {
		return &NotSettableValueError{
			Data: value,
		}
	}

	refValue := reflect.ValueOf(value.Raw())

	if refValue.Type().AssignableTo(fieldRef.Type()) {
		fieldRef.Set(refValue)
		return
	} else if reflect.PtrTo(refValue.Type()).AssignableTo(fieldRef.Type()) {
		var setValue reflect.Value
		if !refValue.CanAddr() {
			setValue = reflect.New(refValue.Type())
			// copying is allowed when we receive non-ptr type
			// since primitive types are always non-ptr
			setValue.Elem().Set(refValue)
		} else {
			setValue = refValue.Addr()
		}

		fieldRef.Set(setValue)
		return
	}

	err = &NotAssignableValueError{
		To:    fieldRef.Type(),
		Value: value,
	}

	return
}

func assignMap(mapRefVal reflect.Value, key interface{}, value Value) (err error) {
	if mapRefVal.Kind() != reflect.Map {
		panic("ndlvr/value: required map reflect value")
	}

	refValue := reflect.ValueOf(value.Raw())

	if refValue.Type().AssignableTo(mapRefVal.Type().Elem()) {
		mapRefVal.SetMapIndex(reflect.ValueOf(key), refValue)
		return
	} else if reflect.PtrTo(refValue.Type()).AssignableTo(mapRefVal.Type().Elem()) {
		var setValue reflect.Value
		if !refValue.CanAddr() {
			setValue = reflect.New(refValue.Type())
			// copying is allowed when we receive non-ptr type
			// since primitive types are always non-ptr
			setValue.Elem().Set(refValue)
		} else {
			setValue = refValue.Addr()
		}

		mapRefVal.SetMapIndex(reflect.ValueOf(key), setValue)
		return
	} else {
		err = &NotAssignableValueError{
			To:    mapRefVal.Type().Elem(),
			Value: value,
		}
		return
	}
}

func assignList(listVal reflect.Value, i int, value Value) (err error) {
	if listVal.Kind() != reflect.Array && listVal.Kind() != reflect.Slice {
		panic("ndlvr/value: required array/slice reflect value")
	}

	if i > listVal.Len() || i < 0 {
		// TODO(teawithsand): OOB index handling here
	}

	err = assignValue(listVal.Index(i), value)
	return
}
