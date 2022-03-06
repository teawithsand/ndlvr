package builder

// Builder to prettify process of creating LIVR rules.
type Builder struct {
	entries map[string][]Rule
}

type Rule struct {
	Name     string
	Argument interface{}
}

func (b *Builder) AddRule(field string, rule Rule) {
	if b.entries == nil {
		b.entries = map[string][]Rule{}
	}
	b.entries[field] = append(b.entries[field], rule)
}

func (b *Builder) Build() (res map[string]interface{}, err error) {
	return
}
