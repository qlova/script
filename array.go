package script

import "fmt"
import "math/big"
import "reflect"
import "github.com/qlova/script/language"

type Array struct {
	language.Array
	EmbeddedScript
	
	Literal []Type
	msg string //debug?
}

func (a Array) SubType() Type {
	return a.q.wrap(a.Array.SubType())
}

func (q Script) Array(T interface{}, length ...int) Array {
	if len(length) > 0 {
		return q.wrap(q.lang.Array(convert(T.(Type)), length[0])).(Array)
	}
	
	if reflect.TypeOf(T).Kind() != reflect.Slice {
		panic(fmt.Sprint("Type: ", reflect.TypeOf(T), ", not allowed in call to script.Array()"))
	}
	
	
	Slice := reflect.ValueOf(T)
	Length := Slice.Len()
	
	var a Array
	a.EmbeddedScript.q = q
	a.Literal = make([]Type, Length)
	a.msg = "Hello"

	for i := 0; i < Length; i++ {
		a.Literal[i] = Slice.Index(i).Interface().(Type)
	}
	
	return a
}

func (q Script) Length(T Type) Number {
	
	if q.Optimise {
		if array, ok := T.(Array); ok {
			return Number{EmbeddedScript: EmbeddedScript{q:q}, Literal: big.NewInt(int64(array.Length()))}
		}
	}
	
	return q.wrap(q.lang.Length(convert(T))).(Number)
}

func (q Script) Index(T Type, index Number) Type {
	return q.wrap(q.lang.Index(convert(T), convert(index)))
}

func (q Script) Modify(T Type, index Number, value Type) {
	q.indent()
	q.write(q.lang.Modify(convert(T), convert(index), convert(value)))
}
