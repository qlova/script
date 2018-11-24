package script

import "github.com/qlova/script/language"

type String struct {
	language.String
	EmbeddedScript

	Literal *string
}
func (String) SameAs(i interface{}) bool { _, ok := i.(String); return ok }

func (String) Joinable() {}

//Return the raw reference to this string.
func (s String) Raw() string {
	if s.Literal != nil {
		return `"`+*s.Literal+`"`
	} else {
		return s.String.Raw()
	}
}

//Converts a Go string to a language.String.
func (q Script) String(s ...string) String {
	str := ""
	if len(s) > 0 {
		str = s[0]
	}
	return String{Literal: &str, EmbeddedScript: EmbeddedScript{ q: q }}
}

type Joinable interface {
	Type
	Joinable()
}

func (q Script) Join(a, b Joinable) Type {
	return q.wrap(q.lang.Join(convert(a), convert(b)))
}
