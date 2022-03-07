package testutil

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/builtin"
	"github.com/teawithsand/ndlvr/value"
)

type E2ETests = []E2ETest

type E2ETest struct {
	Input         interface{}
	ExpectedError error

	Rules ndlvr.RulesSource
}

type AnyError struct{}

func (AnyError) Error() string {
	return "ndlvr/internal/testutil: error, which denotes, that any error may be returned"
}

func (test *E2ETest) Run(t *testing.T) {
	opts := ndlvr.Options{
		ValidationFactory: builtin.MakeBuiltinFactory(),
	}

	ctx := context.Background()

	validator, err := opts.NewEngine(ctx, test.Rules)
	if err != nil {
		t.Error(err)
		return
	}

	err = validator.Validate(context.Background(), value.MustWrap(test.Input))
	if test.ExpectedError != nil {
		// TODO(teawithsand): checking error type here
		if err == nil {
			t.Error("expected error; got nil")
			return
		}
	} else {
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func MustJSONParse(data string) interface{} {
	var res map[string]interface{}
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		panic(err)
	}

	return res
}

func MustJSONParseRules(data string) ndlvr.RulesMap {
	var res ndlvr.RulesMap
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		panic(err)
	}

	return res
}
