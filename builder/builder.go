package builder

import "github.com/teawithsand/ndlvr"

type Builder struct {
	entries map[string][]Rule
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) ensureEntries() {
	if b.entries == nil {
		b.entries = map[string][]Rule{}
	}
}

func (b *Builder) Clone() *Builder {
	var entries map[string][]Rule
	if b.entries != nil {
		entries = make(map[string][]Rule)
		for k, v := range b.entries {
			entries[k] = v
		}
	}

	return &Builder{
		entries: entries,
	}
}

// Merges other builder into current one.
// Does not modify other builder.
// Rules of other builder are appended to rules of current one for each field.
func (b *Builder) MergeBuilder(other *Builder) *Builder {
	b.ensureEntries()
	other.ensureEntries()

	for k, v := range other.entries {
		// TODO(teawithsand): make rules unique?
		// email for instance shouldn't appear twice
		b.entries[k] = append(b.entries[k], v...)
	}

	return b
}

// Merges other builder into current one.
// Does not modify other builder.
// Rules of other builder override rules of current builder, for each field other builder sets any rules.
func (b *Builder) ConsumeBuilder(other *Builder) *Builder {
	b.ensureEntries()
	other.ensureEntries()

	for k, v := range other.entries {
		b.entries[k] = v
	}

	return b
}

func (b *Builder) AddFieldRules(fieldName string, rules []Rule) *Builder {
	b.ensureEntries()
	b.entries[fieldName] = append(b.entries[fieldName], rules...)
	return b
}

func (b *Builder) AddFieldRule(fieldName string, rule Rule) *Builder {
	b.ensureEntries()
	b.entries[fieldName] = append(b.entries[fieldName], rule)
	return b
}

func (b *Builder) AddFieldBuilder(fieldName string, fb *FieldBuilder) *Builder {
	return b.AddFieldRules(fieldName, fb.Build())
}

func (b *Builder) MustBuild() (res ndlvr.TopRulesSource) {
	res, err := b.Build()
	if err != nil {
		panic(err)
	}
	return
}

func (b *Builder) Build() (res ndlvr.TopRulesSource, err error) {
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
