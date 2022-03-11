package builtin_test

import (
	"testing"

	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/internal/testutil"
	"github.com/teawithsand/ndlvr/value"
)

func Test_Like(t *testing.T) {
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
			AddFieldBuilder("field", builder.NewFieldBuilder().MustAddLikeRule("^asdf$")).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "fdsa"
		}`),
		Rules: builder.NewBuilder().
			AddFieldBuilder("field", builder.NewFieldBuilder().MustAddLikeRule("^asdf$")).
			MustBuild(),
		ExpectedError: testutil.AnyError{},
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
