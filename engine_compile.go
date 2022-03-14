package ndlvr

import (
	"context"

	"github.com/teawithsand/ndlvr/value"
)

func (opts *Options) makeValidationsWithTarget(ctx context.Context, rules RulesSource, target ValidationTarget) (res []Validation, err error) {
	p := &Parser{}

	ops := &value.DefaultOPs{}
	argumentParser := &DefaultArgumentParser{}

	err = rules.GetRules(func(rawRule interface{}) (err error) {
		err = p.ParseTopLevelEntry(rawRule, func(rd RuleData) (err error) {
			bctx := ValidationBuildContext{
				Ctx:            ctx,
				Parser:         p,
				Options:        opts,
				OPs:            ops,
				ArgumentParser: argumentParser,
				Data: ValidationBuildData{
					Target: target,

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

// just like makeValidationsWithTarget, but target is defined as key field name
func (opts *Options) makeTopLevelValidations(ctx context.Context, rules TopRulesSource) (res []Validation, err error) {
	p := &Parser{}

	ops := &value.DefaultOPs{}
	argumentParser := &DefaultArgumentParser{}

	err = rules.GetRules(func(fieldName string, rawRule interface{}) (err error) {
		err = p.ParseTopLevelEntry(rawRule, func(rd RuleData) (err error) {
			bctx := ValidationBuildContext{
				Ctx:            ctx,
				Parser:         p,
				Options:        opts,
				OPs:            ops,
				ArgumentParser: argumentParser,
				Data: ValidationBuildData{
					Target: ValidationTarget{
						FieldName: fieldName,
					},

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

func (opts *Options) NewEngine(ctx context.Context, rules TopRulesSource) (v Engine, err error) {
	validations, err := opts.makeTopLevelValidations(ctx, rules)
	if err != nil {
		return
	}

	v = &engineImpl{
		validations: validations,
	}
	return
}

// This is used to compile embedded validations - validations in validations.
// It *should* not be used by user code, unless user is implementing custom validation.
func (opts *Options) NewEngineWithTarget(ctx context.Context, rules RulesSource, target ValidationTarget) (v Engine, err error) {
	validations, err := opts.makeValidationsWithTarget(ctx, rules, target)
	if err != nil {
		return
	}

	v = &engineImpl{
		validations: validations,
	}
	return
}
