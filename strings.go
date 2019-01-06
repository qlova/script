package script

import "github.com/qlova/script/go"
import "github.com/qlova/script/language"

type string struct {
	internal language.String
	script Script
	
	literal *Go.String
	
	symbol *Go.String
}

func (s string) convert(q Script) language.Type {
	if s.literal != nil {
		return q.lang.String(*s.literal)
	} else {
		return s.internal
	}
}

func (s string) String() string {
	return s
}

func (s string) Int() int {
	panic("Cannot cast String to Int")
	return int{}
}


func String(s ...Go.String) string {
	var result string
	if len(s) > 0 {
		result.literal = &(s[0])
	} else {
		result.literal = new(Go.String)
	}
	return result
}
