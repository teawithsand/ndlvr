package builtin_test

import (
	"testing"

	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/internal/testutil"
)

func Test_Required(t *testing.T) {
	var tests testutil.E2ETests

	type FieldStruct struct {
		Field string `json:"field"`
	}

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "asdf"
		}`),
		Rules: builder.NewBuilder().
			AddSimpleRule("field", "required").
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": { "asdf": "fdsa" }
		}`),
		Rules: builder.NewBuilder().
			AddSimpleRule("field", "required").
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: FieldStruct{},
		Rules: builder.NewBuilder().
			AddSimpleRule("Field", "required").
			MustBuild(),
	})

	for _, test := range tests {
		test.Run(t)
		if t.Failed() {
			break
		}
	}
}
