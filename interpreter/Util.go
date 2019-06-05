package interpreter

import "reflect"
import "strconv"
import "github.com/qlova/script/language"
import "github.com/qlova/script/interpreter/dynamic"

func (implementation Implementation) Copy(t language.Type) language.Type {

	if _, ok := t.(List); !ok {
		panic("Unimplemented: Copy(): non-list")
	}

	var register = implementation.ReserveRegister()
	var element = implementation.RegisterOf(t)

	implementation.AddInstruction(func(thread *dynamic.Thread) {
		var element = thread.Get(element)
		var length = reflect.ValueOf(element).Len()

		var destination = reflect.MakeSlice(reflect.TypeOf(element), length, length)

		reflect.Copy(destination, reflect.ValueOf(element))

		thread.Set(register, destination.Interface())
	})

	return t.Register(strconv.Itoa(register))
}
