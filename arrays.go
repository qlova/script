package script

import "github.com/qlova/script/language"

//A String is an ordered sequence of symbols, often used to denote words, sentences and text.
type Array struct {
	script Script
	internal language.Array
	
	subtype Type
	length int
}

func (q Script) Array(elements ...Type) Array {
	var Converted = make([]language.Type, len(elements))
	for i := range elements {
		Converted[i] = elements[i].LanguageType()
	}
	return Array {
		script: q,
		internal: q.lang.Array(Converted...),
	}
}

func (a Array) Value() Value {	
	return Value{
		script: a.script,
		internal: a.LanguageType(),
		
		subtype: a.subtype,
		length: a.length,
	}
}

//Return this Integer as a variable (optionally named).
func (a Array) Var(name ...string) Array {
	return a.Value().Var(name...).Array()
}

//Cast a String to a language.Type ready to be passed to the method of a Language.
func (a Array) LanguageType() language.Type {
	return a.internal
}

//Get this value as a string or cast to a string.
func (v Value) Array() Array {
	if a, ok := v.internal.(language.Array); ok {
		return Array{
			script: v.script,
			internal: a,
			
			subtype: v.subtype,
			length: v.length,
		}
	}
	
	panic("Cannot cast to Array")
	return Array{}
}


//Create a new Int Array with of the specified length.
func (i Int) Array(length int) Array {
	return Array{
		script: i.script,
		internal: i.script.lang.ArrayOf(i.LanguageType(), length),
		length: length,
		subtype: i,
	}
}

//Return the value at index of the Array.
func (a Array) Index(index Int) Value {
	return Value{
		script: a.script,
		internal: a.script.lang.Index(a.LanguageType(), index.LanguageType()),
	}	
}

//Return the value at index of the Array.
func (a Array) Modify(index Int, value Type) {
	a.script.indent()
	a.script.write(a.script.lang.Modify(a.LanguageType(), index.LanguageType(), value.LanguageType()))
}

//Return the value at index of the Array.
func (a Array) Subtype() Type {
	return a.subtype
}
