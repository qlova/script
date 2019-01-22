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
