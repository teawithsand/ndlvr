package builder

import livr "github.com/teawithsand/livr4go"

// Builder to prettify process of creating LIVR rules.
type Builder struct {
	entries map[string][]Rule
}

type Rule struct {
	Name     string
	Argument interface{}
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

func (b *Builder) Build() (res livr.RulesSource, err error) {
	rm := livr.RulesMap{}
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

func (b *Builder) MustBuild() (res livr.RulesSource) {
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
