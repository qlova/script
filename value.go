package script

type Value interface {
	T() Type
}

type Values []Value

//Runtime retrieved the runtime values.
func (values Values) Runtime() (result []interface{}) {
	for _, value := range values {
		result = append(result, value.T().Get())
	}
	return
}
