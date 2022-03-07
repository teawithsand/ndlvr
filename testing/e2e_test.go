package testing_test

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/teawithsand/ndlvr"
	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/builtin"
	"github.com/teawithsand/ndlvr/value"
)

type E2ETest struct {
	Input          interface{}
	ExpectedOutput interface{}
	ExpectedError  bool

	Rules ndlvr.RulesSource
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
	if test.ExpectedError {
		if err == nil {
			t.Error("expected error; got nil")
			return
		}
	} else {
		if err != nil {
			t.Error(err)
			return
		}

		if !reflect.DeepEqual(test.ExpectedOutput, test.Input) {
			t.Error("output mistmatch", test.ExpectedOutput)
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

func TestLIVR(t *testing.T) {
	test := E2ETest{
		Input: MustJSONParse(`
		{
			"first_name": "Vasya",
			"last_name": "Pupkin",
			"middle_name": "Some",
			"age": "25",
			"salary": 0
		}`),
		ExpectedOutput: MustJSONParse(`
		{
			"first_name": "Vasya",
			"last_name": "Pupkin",
			"middle_name": "Some",
			"age": "25",
			"salary": 0
		}`),
		Rules: (&builder.Builder{}).
			AddRule("first_name", builder.Rule{Name: "required"}).
			AddRule("last_name", builder.Rule{Name: "required"}).
			AddRule("middle_name", builder.Rule{Name: "required"}).
			AddRule("salary", builder.Rule{Name: "required"}).
			MustBuild(),
	}

	test.Run(t)
}
