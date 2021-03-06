package builtin_test

import (
	"testing"

	"github.com/teawithsand/ndlvr/builder"
	"github.com/teawithsand/ndlvr/internal/testutil"
	"github.com/teawithsand/ndlvr/value"
)

func Test_MinLength(t *testing.T) {
	var tests testutil.E2ETests

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "asdf"
		}`),
		Rules: builder.NewBuilder().
			AddFieldBuilder("field", builder.NewFieldBuilder().AddMinLength(4)).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "asdf"
		}`),
		ExpectedError: testutil.AnyError{},
		Rules: builder.NewBuilder().
			AddFieldBuilder("field", builder.NewFieldBuilder().AddMinLength(5)).
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

func Test_MaxLength(t *testing.T) {
	var tests testutil.E2ETests

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "asdf"
		}`),
		Rules: builder.NewBuilder().
			AddFieldBuilder("field", builder.NewFieldBuilder().AddMaxLength(4)).
			MustBuild(),
	})

	tests = append(tests, testutil.E2ETest{
		Input: testutil.MustJSONParse(`
		{
			"field": "ffffffff"
		}`),
		ExpectedError: testutil.AnyError{},
		Rules: builder.NewBuilder().
			AddFieldBuilder("field", builder.NewFieldBuilder().AddMaxLength(4)).
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
