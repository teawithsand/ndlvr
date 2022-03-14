package testutil

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/builtin"
	"github.com/teawithsand/ndlvr/value"
)

type E2ETests []E2ETest

func (tests E2ETests) Mutate(m func(t *E2ETest)) {
	for i := range tests {
		m(&tests[i])
	}
}

type E2ETest struct {
	Input         interface{}
	ExpectedError error

	Wrapper value.Wrapper
	Rules   ndlvr.TopRulesSource
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

	var w value.Wrapper = &value.DefaultWrapper{}
	if test.Wrapper != nil {
		w = test.Wrapper
	}

	err = validator.Validate(context.Background(), value.WrapperMustWrap(w.Wrap(test.Input)))
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
