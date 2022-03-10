package builder

import "regexp"

type FieldBuilder struct {
	Builder *Builder
	Field   string
}

func (b *FieldBuilder) AddRule(rule Rule) *FieldBuilder {
	b.Builder.addPredefinedRule(b.Field, rule)

	return b
}

func (b *FieldBuilder) AddSimpleRule(rule string) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     rule,
		Argument: nil,
	})
}

// Adds "like" rule and panics when regex is not valid.
func (b *FieldBuilder) AddLikeRule(regex string) *FieldBuilder {
	regexp.MustCompile(regex)

	return b.AddRule(Rule{
		Name:     "like",
		Argument: regex,
	})
}

func (b *FieldBuilder) AddRequired() *FieldBuilder {
	return b.AddSimpleRule("required")
}

func (b *FieldBuilder) AddNotEmpty() *FieldBuilder {
	return b.AddSimpleRule("not_empty")
}

func (b *FieldBuilder) AddMaxLength(sz int) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     "max_length",
		Argument: sz,
	})
}

func (b *FieldBuilder) AddMinLength(sz int) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     "min_length",
		Argument: sz,
	})
}

func (b *FieldBuilder) Done() *Builder {
	return b.Builder
}