package builder

type Rule struct {
	Name     string
	Argument interface{}
}

// Renders rule into LIVR language using simplest notation possible.
func (r *Rule) Render() interface{} {
	if r.Argument == nil {
		return r.Name
	}
	return map[string]interface{}{
		r.Name: r.Argument,
	}
}
