package script

import "github.com/qlova/script/language"
import "math/big"

//An Int is a numeric integer value consisting of a magnitude and a sign.
type Int struct {
	script Script
	internal language.Integer

	literal *big.Int
}

//Wrap a language.Type to an Integer.
func (Int) FromLanguageType(T language.Type) Int {
	if internal, ok := T.(language.Integer); ok {
		return Int{
			internal: internal,
		}
	}
	panic("Invalid wrap!")
	return Int{}
}

//Return this Integer as a variable (optionally named).
func (i Int) Var(name ...string) Int {
	var register string
	if len(name) > 0 {
		register = name[0]
	} else {
		register = Unique()
	}
	
	i.script.indent()
	statement, variable := i.script.lang.Register(register, i.LanguageType())
	i.script.write(statement)
	
	return Int{
		script: i.script,
		internal: variable.(language.Integer),
	}
}

//Cast an Int to a language.Type ready to be passed to the method of a Language.
func (i Int) LanguageType() language.Type {
	if i.literal != nil {
		return i.script.lang.Integer(int(i.literal.Int64()))
	} else {
		return i.internal
	}
}

//Return a new String type with the value s.
func (q Script) Int(i ...int) Int {
	var literal = big.NewInt(0)

	if len(i) > 0 {
		literal = big.NewInt(int64(i[0]))
	}

	return Int{
		script: q,
		literal: literal,
	}
}

func (a Int) Add(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var sum = big.NewInt(0).Add(a.literal, b.literal)
		return Int{
			script: a.script,
			literal: sum,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Add(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}

func (a Int) Sub(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var difference = big.NewInt(0).Sub(a.literal, b.literal)
		return Int{
			script: a.script,
			literal: difference,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Sub(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}

func (a Int) Mul(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var product = big.NewInt(0).Mul(a.literal, b.literal)
		return Int{
			script: a.script,
			literal: product,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Mul(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}

func (a Int) Div(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var quotient = big.NewInt(0).Div(a.literal, b.literal)
		return Int{
			script: a.script,
			literal: quotient,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Div(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}

func (a Int) Mod(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var modulus = big.NewInt(0).Mod(a.literal, b.literal)
		return Int{
			script: a.script,
			literal: modulus,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Mod(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}

func (a Int) Pow(b Int) Int {
	if a.literal != nil && b.literal != nil {
		var exponent = big.NewInt(0).Exp(a.literal, b.literal, nil)
		return Int{
			script: a.script,
			literal: exponent,
		}
	}
	return Int{}.FromLanguageType(a.script.lang.Pow(a.LanguageType().(language.Number), b.LanguageType().(language.Number)))
}
