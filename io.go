package script

import "github.com/qlova/script/language"

func (q Script) Print(values ...Type) {
	
	var converted = make([]language.Type, len(values))
	for i := range values {
		converted[i] = values[i].LanguageType()
	}
	
	q.indent()
	q.write(q.lang.Print(converted...))
}

func (q Script) Write(values ...Type) {
	
	var converted = make([]language.Type, len(values))
	for i := range values {
		converted[i] = values[i].LanguageType()
	}
	
	q.indent()
	q.write(q.lang.Write(nil, converted...))
}

func (q Script) Read(mode Type) Value {
	return q.ValueFromLanguageType(q.lang.Read(nil, mode.LanguageType()))
}
