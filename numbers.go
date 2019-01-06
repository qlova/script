package script

import "github.com/qlova/script/go"
import "github.com/qlova/script/language"
import "math"

type int struct {
	internal language.Integer
	script Script
	
	literal *Go.Int

	symbol *Go.String
}

func (*int) wrap(T interface{}) int {
	if i, ok := T.(language.Integer); ok {
		return int{internal: i}
	}
	panic("Invalid wrap!")
	return int{}
}

func (i int) convert(q Script) language.Type {
	if i.symbol != nil {
		return q.lang.Get(*i.symbol, i.internal)
	} else if i.literal != nil {
		return q.lang.Integer(*i.literal)
	} else {
		return i.internal
	}
}

func (i int) Int() int {
	return i
}

func (i int) String() string {
	panic("Cannot cast Int to String")
	return string{}
}


func Int(s ...Go.Int) int {
	var result int
	if len(s) > 0 {
		result.literal = &(s[0])
	} else {
		result.literal = new(Go.Int)
	}
	return result
}

func (q Script) Int(i ...int) int {
	
	var unique = Unique()
	var value = Int()
	if len(i) > 0 {
		value = i[0]
	}
	
	q.indent()
	q.write(q.lang.Register(unique, value.convert(q)))
	
	return int{
		internal: value.convert(q).(language.Integer),
		symbol: &unique,
		script: q,
	}
}

func (i int) Add(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var sum = (*i.literal + *b.literal)
			return int{ literal: &sum }
		}
	}
	return new(int).wrap(q.lang.Add(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}

func (i int) Mul(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var product = (*i.literal * *b.literal)
			return int{ literal: &product }
		}
	}
	return new(int).wrap(q.lang.Mul(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}

func (i int) Sub(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var difference = (*i.literal - *b.literal)
			return int{ literal: &difference }
		}
	}
	return new(int).wrap(q.lang.Sub(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}

func (i int) Div(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var quotient = (*i.literal / *b.literal)
			return int{ literal: &quotient }
		}
	}
	return new(int).wrap(q.lang.Div(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}

func (i int) Mod(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var modulus = (*i.literal % *b.literal)
			return int{ literal: &modulus }
		}
	}
	return new(int).wrap(q.lang.Sub(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}

func (i int) Pow(b int) int {
	var q = i.script
	if q.script == nil {
		q = b.script
		if q.script == nil {
			var exponent = Go.Int(math.Pow(float64(*i.literal), float64(*b.literal)))
			return int{ literal: &exponent }
		}
	}
	return new(int).wrap(q.lang.Sub(i.convert(q).(language.Number), b.convert(q).(language.Number)))
}
