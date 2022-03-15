package builder

import (
	"regexp"

	"github.com/teawithsand/ndlvr"
)

type FieldBuilder struct {
	rules []Rule
}

func NewFieldBuilder() *FieldBuilder {
	return &FieldBuilder{}
}

func (b *FieldBuilder) Clone() *FieldBuilder {
	rules := make([]Rule, len(b.rules))
	copy(rules, b.rules)

	return &FieldBuilder{
		rules: rules,
	}
}

func (b *FieldBuilder) AddRule(rule Rule) *FieldBuilder {
	b.rules = append(b.rules, rule)

	return b
}

func (b *FieldBuilder) AddRuleExt(name string, argument interface{}) *FieldBuilder {
	b.rules = append(b.rules, Rule{
		Name:     name,
		Argument: argument,
	})

	return b
}

func (b *FieldBuilder) AddSimpleRule(rule string) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     rule,
		Argument: nil,
	})
}

// Adds "like" rule and panics when regex is not valid.
func (b *FieldBuilder) MustAddLikeRule(regex string) *FieldBuilder {
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

func (b *FieldBuilder) AddListOf(fieldBuilder *FieldBuilder) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     "list_of",
		Argument: fieldBuilder.BuildRaw(),
	})
}

func (b *FieldBuilder) AddListOfObjects(builder *Builder) *FieldBuilder {
	return b.AddRule(Rule{
		Name:     "list_of_objects",
		Argument: (map[string]interface{})(builder.MustBuild().(ndlvr.RulesMap)),
	})
}

func (b *FieldBuilder) BuildRaw() interface{} {
	var rendered []interface{}
	for _, r := range b.rules {
		rendered = append(rendered, r.Render())
	}
	return rendered
}

func (b *FieldBuilder) Build() []Rule {
	return b.rules
}
