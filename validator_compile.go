package livr

import (
	"context"

	"github.com/teawithsand/livr4go/value"
)

func (opts *Options) makeValidationsMap(ctx context.Context) (res []Validation, err error) {
	p := &Parser{}

	ops := &value.DefaultOPs{}
	argumentParser := &DefaultArgumentParser{}

	err = opts.Rules.GetRules(func(fieldName string, rawRule interface{}) (err error) {
		err = p.ParseTopLevelEntry(fieldName, rawRule, func(rd RuleData) (err error) {
			validation, err := opts.ValidationFactory.BuildValidation(ValidationBuildContext{
				Ctx:            ctx,
				Parser:         p,
				Options:        opts,
				OPs:            ops,
				ArgumentParser: argumentParser,
				Data: ValidationBuildData{
					FieldName:      fieldName,
					ValidationName: rd.ValidationName,
					Argument:       rd.ValidationArgument,
				},
			})
			if err != nil {
				return
			}

			res = append(res, validation)
			return
		})
		return
	})
	if err != nil {
		return
	}

	return
}

func (opts *Options) NewValidator(ctx context.Context) (v Validator, err error) {
	validations, err := opts.makeValidationsMap(ctx)
	if err != nil {
		return
	}

	v = &validatorImpl{
		validations: validations,
	}
	return
}
