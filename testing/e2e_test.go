package testing_test

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	livr "github.com/teawithsand/livr4go"
	"github.com/teawithsand/livr4go/builder"
	"github.com/teawithsand/livr4go/value"
)

type E2ETest struct {
	Input          interface{}
	ExpectedOutput interface{}
	ExpectedError  bool

	Rules livr.RulesSource
}

func (test *E2ETest) Run(t *testing.T) {
	opts := livr.Options{
		ValidationFactory: livr.MakeBuiltinFactory(),
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

func MustJSONParseRules(data string) livr.RulesMap {
	var res livr.RulesMap
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
