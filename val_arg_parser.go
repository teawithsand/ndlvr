package livr

import (
	"context"
	"fmt"

	"github.com/teawithsand/ndlvr/value"
)

// Helper, which makes parsing incoming validation's argument simpler.
type ArgumentParser interface {
	// Parses length from argument given.
	ParseLen(ctx context.Context, arg interface{}) (sz int, err error)
	ParsePrimitiveValue(ctx context.Context, arg interface{}) (pv *value.PrimitiveValue, err error)
	ParseListValue(ctx context.Context, arg interface{}) (lv value.ListValue, err error)

	// TODO(teawithsand): implement that
	// Prases argument incoming to specifeid struct pointer like json.Unmarshal
	// ParseStruct(ctx context.Context, arg interface{}, res interface{}) (err error)
}

type DefaultArgumentParser struct{}

func (parser *DefaultArgumentParser) ParseLen(ctx context.Context, v interface{}) (res int, err error) {
	switch tv := v.(type) {
	case int:
		res = tv
		return

	// TODO(teawithsand): javascript would handle all length comparison on floats
	//  so this behaviour is incompatibile with what JS would do if there is FP part in number/
	case float32:
		res = int(tv)
		return
	case float64:
		res = int(tv)
		return
	default:
		err = &ValidationCreateError{
			Msg: fmt.Sprintf("unsupported length value type: %T", v),
		}
		return
	}
}

func (parser *DefaultArgumentParser) ParsePrimitiveValue(ctx context.Context, v interface{}) (pv *value.PrimitiveValue, err error) {
	vv, err := value.Wrap(v)
	if err != nil {
		return
	}

	pv, ok := vv.(*value.PrimitiveValue)
	if !ok {
		err = &ValidationCreateError{
			Msg: fmt.Sprintf("Value is not primitive: %T", v),
		}
	}
	return
}

func (parser *DefaultArgumentParser) ParseListValue(ctx context.Context, v interface{}) (lv value.ListValue, err error) {
	vv, err := value.Wrap(v)
	if err != nil {
		return
	}

	lv, ok := vv.(value.ListValue)
	if !ok {
		err = &ValidationCreateError{
			Msg: fmt.Sprintf("Value is not list: %T", v),
		}
	}
	return
}
