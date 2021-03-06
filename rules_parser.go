package ndlvr

type TopLevelRuleReceiver = func(rd TopLevelRuleData) (err error)
type RuleReceiver = func(rd RuleData) (err error)

type TopLevelRuleData struct {
	FieldName string
	RuleData
}

type RuleData struct {
	ValidationName     string
	ValidationArgument interface{}
}

// Parser, which is used to parse LIVR rules into immediate representation, which can be used
// to construct Validations and Validators.
// For now it has no options.
type Parser struct{}

func (p *Parser) ParseTopLevelEntry(rawRule interface{}, recv RuleReceiver) (err error) {
	switch rule := rawRule.(type) {
	case map[string]interface{}:
		for k, v := range rule {
			err = recv(RuleData{
				ValidationName:     k,
				ValidationArgument: v,
			})
			if err != nil {
				return
			}
		}
	case []interface{}:
		for _, e := range rule {
			err = p.ParseInnerEntry(e, recv)
			if err != nil {
				return
			}
		}
	case string:
		err = recv(RuleData{
			ValidationName: rule,
		})
		if err != nil {
			return
		}
	default:
		err = &RuleParseError{
			Rule: rawRule,
		}
		return
	}
	return
}

// Just like ParseTopLevelRules, but disallows list.
// This way nested lists in rules are not allowed.
//
// Think of:
// ```
// ...
// "asdf": [ ["required"] ] // Not allowed
// "fdsa": [ "required", { "list_of_objects" : { ... }}] // Ok
// ...
// ```
func (p *Parser) ParseInnerEntry(rawRule interface{}, recv RuleReceiver) (err error) {
	switch rule := rawRule.(type) {
	case map[string]interface{}:
		for k, v := range rule {
			err = recv(RuleData{
				ValidationName:     k,
				ValidationArgument: v,
			})
			if err != nil {
				return
			}
		}
	case string:
		err = recv(RuleData{
			ValidationName: rule,
		})
		if err != nil {
			return
		}
	default:
		err = &RuleParseError{
			Rule: rawRule,
		}
	}
	return
}
