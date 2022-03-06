package value

// TODO(teawtihsand): rest of primitives here

/*
type StringValue string

func (sv StringValue) Raw() interface{} {
	return string(sv)
}

func (sv StringValue) IsEmpty() bool {
	return sv == ""
}

type IntValue int

func (sv IntValue) IsEmpty() bool {
	return sv == 0 // ?
}

func (sv IntValue) Raw() interface{} {
	return int(sv)
}

type Float64Value float64

func (sv Float64Value) IsEmpty() bool {
	return sv == 0 // ?
}

func (sv Float64Value) Raw() interface{} {
	return float64(sv)
}

*/

type PrimitiveValue struct {
	Val interface{}
}

func (pv *PrimitiveValue) Raw() interface{} {
	if pv == nil {
		return nil
	}
	return pv.Val
}
