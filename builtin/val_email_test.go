package builtin_test

import (
	"testing"

	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/internal/testutil"
	"github.com/teawithsand/ndlvr/value"
)

func Test_Email(t *testing.T) {
	var tests testutil.E2ETests

	rules := builder.NewBuilder().
		AddFieldBuilder("field", builder.NewFieldBuilder().AddSimpleRule("email")).
		MustBuild()

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "email@gmail.com"
		}`),
		Rules: rules,
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@gmail.com"
		}`),
		ExpectedError: testutil.AnyError{},
		Rules:         rules,
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "garbage"
		}`),
		ExpectedError: testutil.AnyError{},
		Rules:         rules,
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "garbage"
		}`),
		ExpectedError: testutil.AnyError{},
		Rules:         rules,
	})

	tests.Mutate(func(t *testutil.E2ETest) {
		t.Wrapper = &value.DefaultWrapper{
			UseJSONNames: true,
		}
	})

	for _, test := range tests {
		test.Run(t)
		if t.Failed() {
			break
		}
	}
}
