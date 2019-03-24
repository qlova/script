package script

import "github.com/qlova/script/language"

//A String is an ordered sequence of symbols, often used to denote words, sentences and text.
type List struct {
	script Script
	internal language.List
	
	subtype Type
}

func (q Script) List(elements ...Type) List {
	var Converted = make([]language.Type, len(elements))
	for i := range elements {
		Converted[i] = elements[i].LanguageType()
	}
	return List {
		script: q,
		internal: q.lang.List(Converted...),
	}
}

func (l List) Value() Value {	
	return Value{
		script: l.script,
		internal: l.LanguageType(),
		
		subtype: l.subtype,
	}
}

//Return this Integer as a variable (optionally named).
func (l List) Var(name ...string) List {
	return l.Value().Var(name...).List()
}

//Cast a String to a language.Type ready to be passed to the method of a Language.
func (l List) LanguageType() language.Type {
	return l.internal
}

//Get this value as a string or cast to a string.
func (v Value) List() List {
	if l, ok := v.internal.(language.List); ok {
		return List{
			script: v.script,
			internal: l,
			
			subtype: v.subtype,
		}
	}
	
	panic("Cannot cast to Array")
	return List{}
}

//Return the value at index of the Array.
func (l List) Index(index Int) Value {
	return Value{
		script: l.script,
		internal: l.script.lang.Index(l.LanguageType(), index.LanguageType()),
	}	
}

//Return the value at index of the Array.
func (l List) Modify(index Int, value Type) {
	l.script.indent()
	l.script.write(l.script.lang.Modify(l.LanguageType(), index.LanguageType(), value.LanguageType()))
}

//Return the value at index of the Array.
func (l List) Subtype() Type {
	return l.subtype
}

func (l List) ForEach(f func(), names ...string) {
	l.script.foreach(l, f, names...)
}
