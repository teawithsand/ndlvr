package ndlvr

// RulesSource is just source of rules.
// For each key it returns rules.
//
// Usually, it's JSON parsed map but in future go tags probably will be supported.
type RulesSource interface {
	// Note: each key may be yielded only once.
	GetRules(recv func(fieldName string, rawRule interface{}) (err error)) (err error)
}

// RulesMap defines type to deserialize from JSON when parsing is required.
type RulesMap map[string]interface{}

func (r RulesMap) GetRules(recv func(key string, value interface{}) (err error)) (err error) {
	for k, v := range r {
		err = recv(k, v)
		if err != nil {
			return
		}
	}

	return
}
