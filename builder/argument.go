package builder

type ArgumentMap map[string]interface{}

func (am ArgumentMap) Put(key string, val interface{}) {
	am[key] = val
}
