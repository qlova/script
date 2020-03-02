package script

type Value interface {
	T() Type
	AnyValue
}

type AnyValue interface {
	ValueFromCtx(AnyCtx) Value
}

type Values []Value

//Runtime retrieved the runtime values.
func (values Values) Runtime() (result []interface{}) {
	for _, value := range values {
		result = append(result, value.T().Get())
	}
	return
}
