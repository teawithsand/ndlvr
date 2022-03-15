package builtin_test

import (
	"testing"

	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/internal/testutil"
	"github.com/teawithsand/ndlvr/value"
)

func Test_ListOfObjects(t *testing.T) {
	var tests testutil.E2ETests

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": []
		}`),
		Rules: builder.NewBuilder().
			AddFieldBuilder(
				"field",
				builder.NewFieldBuilder().AddListOfObjects(
					builder.NewBuilder().
						AddFieldBuilder("asdf", builder.NewFieldBuilder().AddSimpleRule("positive_integer")),
				),
			).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": [{
				"asdf": 1234
			}]
		}`),
		Rules: builder.NewBuilder().
			AddFieldBuilder(
				"field",
				builder.NewFieldBuilder().AddListOfObjects(
					builder.NewBuilder().
						AddFieldBuilder("asdf", builder.NewFieldBuilder().AddSimpleRule("positive_integer")),
				),
			).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": [{
				"asdf": "fdsa"
			}]
		}`),
		ExpectedError: testutil.AnyError{},
		Rules: builder.NewBuilder().
			AddFieldBuilder(
				"field",
				builder.NewFieldBuilder().AddListOfObjects(
					builder.NewBuilder().
						AddFieldBuilder("asdf", builder.NewFieldBuilder().AddSimpleRule("positive_integer")),
				),
			).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": [{
				"asdf": 1234
			}, {
				"asdf": "fdsa"
			}]
		}`),
		ExpectedError: testutil.AnyError{},
		Rules: builder.NewBuilder().
			AddFieldBuilder(
				"field",
				builder.NewFieldBuilder().AddListOfObjects(
					builder.NewBuilder().
						AddFieldBuilder("asdf", builder.NewFieldBuilder().AddSimpleRule("positive_integer")),
				),
			).
			MustBuild(),
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
