package script

import "reflect"

func GoTypeOf(value Value) reflect.Type {
	return reflect.TypeOf(GoValueOf(value))
}

func GoValueOf(value Value) interface{} {
	switch value.(type) {
	case Int:
		return int(0)
	case String:
		return string("")
	case Bool:
		return bool(false)
	default:
		return nil
	}
}
