package livr

import (
	"context"
	"fmt"
)

// Helper, which makes parsing incoming validation's argument simpler.
type ArgumentParser interface {
	// Parses length from argument given.
	ParseLen(ctx context.Context, arg interface{}) (sz int, err error)

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
