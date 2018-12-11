package script

import "github.com/qlova/script/language"
import "github.com/qlova/script/go"

type string struct {
	literal *Go.String

	EmbeddedScript
	language.String
}

func String(s Go.String) string {
	return string{literal: &s}
}

func (s string) Raw() Go.String {
	if s.literal != nil {
		return `"`+*s.literal+`"`
	} else {
		return s.String.Raw()
	}
}

/*func (String) SameAs(i interface{}) bool { _, ok := i.(string); return ok }

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
*/