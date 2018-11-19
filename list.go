package script

import "reflect"
import "fmt"
import "github.com/qlova/script/language"

type List struct {
	language.List
	EmbeddedScript
	
	Literal []Type
}

func (q Script) List(T ...Type) List {
	if len(T) == 0 {
		panic(fmt.Sprint("script.List: List cannot have zero elements!"))
	}
	
	var FirstType = T[0]
	
	//We are going to make a slice of the first type.
	var Converted = reflect.MakeSlice(
			reflect.SliceOf(reflect.TypeOf(convert(FirstType))), len(T), len(T))
	
	for i, v := range T {
		if !v.SameAs(FirstType) {
			panic(fmt.Sprint("script.List: Inconsistent List!"))
		}
		Converted.Index(i).Set(reflect.ValueOf(convert(v)))
	}
	
	return q.wrap(q.lang.Literal(Converted.Interface())).(List)
}
