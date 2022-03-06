package livr

import (
	"context"

	"github.com/teawithsand/livr4go/value"
)

func (opts *Options) makeValidationsMap(ctx context.Context, rules RulesSource) (res []Validation, err error) {
	p := &Parser{}

	ops := &value.DefaultOPs{}
	argumentParser := &DefaultArgumentParser{}

	err = rules.GetRules(func(fieldName string, rawRule interface{}) (err error) {
		err = p.ParseTopLevelEntry(fieldName, rawRule, func(rd RuleData) (err error) {
			bctx := ValidationBuildContext{
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
			}
			validation, err := opts.ValidationFactory.BuildValidation(bctx)
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

func (opts *Options) NewEngine(ctx context.Context, rules RulesSource) (v Engine, err error) {
	validations, err := opts.makeValidationsMap(ctx, rules)
	if err != nil {
		return
	}

	v = &engineImpl{
		validations: validations,
	}
	return
}
