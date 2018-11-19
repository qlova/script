package Go

import "math/big"
import "github.com/qlova/script/language"

type Number struct {
	Expression string

	Literal *big.Int
}

func (Number) Name() string { return "number" }
func (Number) SameAs(i interface{}) bool {
	_, ok := i.(Number)
	return ok
}
func (Number) Number() {}

//Returns a Number that the Go style literal represents (01 1 0x1).
func (l *implementation) LiteralNumber(literal *big.Int) language.Number {
	return Number{Literal: literal}
}

//Returns a Number that is the sum of 'a' and 'b'.
func (l *implementation) Add(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Add(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func (a Number) Add(b Number) Number {
	if ca, cb := a.Large == nil, b.Large == nil; ca || cb {
		if ca && cb {
			if (a.Small > 0 && b.Small > (1<<63 - 1) - a.Small) || (a.Small < 0 && b.Small < (-1 << 63) - a.Small) {
				return Number{Large: new(big.Int).Add(big.NewInt(a.Small), big.NewInt(b.Small))}
			}
			return Number{Small:a.Small+b.Small}
		} else if !ca {
			return Number{Large: new(big.Int).Add(a.Large, big.NewInt(b.Small))}
		} else if !cb {
			return Number{Large: new(big.Int).Add(big.NewInt(a.Small), b.Large)}
		}
	}
	return Number{Large: new(big.Int).Add(a.Large, b.Large)}
}
`)
	
	return Number{Expression: l.GetExpression(A)+".Add("+l.GetExpression(B)+")"}
}

//Returns a Number that is the difference of 'a' and 'b'.
func (l *implementation) Sub(a, b language.Number) language.Number {
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Sub(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func Sub(a, b *big.Int) *big.Int {
	var z = big.NewInt(0)
	z.Sub(a, b)
	return &z
}
`)
	
	return Number{Expression: "Sub("+l.GetExpression(A)+","+l.GetExpression(B)+")"}
}

//Returns a Number that is the product of 'a' and 'b'.
func (l *implementation) Mul(a, b language.Number) language.Number {
	
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Mul(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func (a Number) Mul(b Number) Number {
	if ca, cb := a.Large == nil, b.Large == nil; ca || cb {
		if ca && cb {
			var z = Number{Small:a.Small*b.Small}
			if (a.Small != 0 && z.Small / a.Small != b.Small) {
				return Number{Large:new(big.Int).Mul(big.NewInt(a.Small), big.NewInt(b.Small))}
			}
			return z
		} else if !ca {
			return Number{Large: new(big.Int).Mul(a.Large, big.NewInt(b.Small))}
		} else if !cb {
			return Number{Large: new(big.Int).Mul(big.NewInt(a.Small), b.Large)}
		}
	}
	return Number{Large: new(big.Int).Mul(a.Large, b.Large)}
}
`)
	
	return Number{Expression: l.GetExpression(A)+".Mul("+l.GetExpression(B)+")"}
}

//Returns a Number that is the quotient of 'a' and 'b'.
func (l *implementation) Div(a, b language.Number) language.Number {
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Div(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func Mul(a, b *big.Int) *big.Int {
	var z = big.NewInt(0)
	z.Div(a, b)
	return &z
}
`)
	
	return Number{Expression: "Div("+l.GetExpression(A)+","+l.GetExpression(B)+")"}
}

//Returns a Number that is 'a' taken to the power of 'b'.
func (l *implementation) Pow(a, b language.Number) language.Number {
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Exp(A.Literal, B.Literal, nil)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func Pow(a, b *big.Int) *big.Int {
	var z = big.NewInt(0)
	z.Exp(a, b, nil)
	return &z
}
`)
	
	return Number{Expression: "Pow("+l.GetExpression(A)+","+l.GetExpression(B)+")"}
}

//Returns a Number that is modulos of 'a' and 'b'.
func (l *implementation) Mod(a, b language.Number) language.Number {
	A, B := a.(Number), b.(Number)
	if A.Literal != nil && B.Literal != nil {
		var z big.Int
		z.Mod(A.Literal, B.Literal)
		return Number{Literal: &z}
	}
	
	l.Import("math/big")
	l.AddHelper(`func Mul(a, b *big.Int) *big.Int {
	var z = big.NewInt(0)
	z.Mod(a, b)
	return &z
}
`)
	
	return Number{Expression: "Mod("+l.GetExpression(A)+","+l.GetExpression(B)+")"}
}
