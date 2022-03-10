package builder

import "github.com/teawithsand/ndlvr"

// Builder to prettify process of creating LIVR rules.
type Builder struct {
	entries map[string][]Rule
}

func NewBuilder() *Builder {
	return &Builder{}
}

type Rule struct {
	Name     string
	Argument interface{}
}

func (b *Builder) AddRequired(field string) *Builder {
	return b.AddSimpleRule(field, "required")
}

func (b *Builder) AddNotEmpty(field string) *Builder {
	return b.AddSimpleRule(field, "not_empty")
}

// Like add rule but for predefined rules.
func (b *Builder) addPredefinedRule(field string, rule Rule) *Builder {
	if b.entries == nil {
		b.entries = map[string][]Rule{}
	}
	b.entries[field] = append(b.entries[field], rule)

	return b
}

// Adds arbitrary rule to builder
func (b *Builder) AddRule(field string, rule Rule) *Builder {
	if b.entries == nil {
		b.entries = map[string][]Rule{}
	}
	b.entries[field] = append(b.entries[field], rule)

	return b
}

// Adds arbitrary rule to builder
func (b *Builder) AddSimpleRule(field string, rule string) *Builder {
	return b.AddRule(field, Rule{
		Name:     rule,
		Argument: nil,
	})
}

func (b *Builder) Field(field string) *FieldBuilder {
	return &FieldBuilder{
		Builder: b,
		Field:   field,
	}
}

func (b *Builder) Build() (res ndlvr.RulesSource, err error) {
	rm := ndlvr.RulesMap{}
	for field, rules := range b.entries {
		for _, rule := range rules {
			_, ok := rm[field]
			if !ok {
				rm[field] = []interface{}{
					map[string]interface{}{
						rule.Name: rule.Argument,
					},
				}
			} else {
				rm[field] = append(rm[field].([]interface{}), map[string]interface{}{
					rule.Name: rule.Argument,
				})
			}
		}
	}

	res = rm
	return
}

func (b *Builder) MustBuild() (res ndlvr.RulesSource) {
	res, err := b.Build()
	if err != nil {
		panic(err)
	}
	return
}

type ArgumentMap map[string]interface{}

func (am ArgumentMap) Put(key string, val interface{}) {
	am[key] = val
}
