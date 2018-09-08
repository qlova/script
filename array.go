package script

import "fmt"
import "reflect"
import "github.com/qlova/script/language"

type Array struct {
	language.Array
	EmbeddedScript
	
	Literal []Type
}

func (a Array) SubType() Type {
	return a.q.wrap(a.Array.SubType())
}

func (q *Script) Array(T interface{}, length ...Number) Array {
	if len(length) > 0 {
		return q.wrap(q.lang.Array(convert(T.(Type)), convert(length[0]).(language.Number))).(Array)
	}
	
	if reflect.TypeOf(T).Kind() != reflect.Slice {
		panic(fmt.Sprint("Type: ", reflect.TypeOf(T), ", not allowed in call to script.Array()"))
	}
	
	
	Slice := reflect.ValueOf(T)
	Length := Slice.Len()
	
	//Reconstruct.
	var elements = make([]language.Type, Length)

	for i := 0; i < Length; i++ {
		elements[i] = convert(Slice.Index(i).Interface().(Type)).(language.Type)
	}

	return q.wrap(q.lang.Fill(q.lang.Array(elements[0], convert(q.Number(Length)).(language.Number)), elements)).(Array)
}

func (q *Script) Length(T Type) Number {
	return q.wrap(q.lang.Length(convert(T))).(Number)
}

func (q *Script) Index(T Type, index Number) Type {
	return q.wrap(q.lang.Index(convert(T), convert(index)))
}

func (q *Script) Modify(T Type, index Number, value Type) {
	q.indent()
	q.write(q.lang.Modify(convert(T), convert(index), convert(value)))
}
