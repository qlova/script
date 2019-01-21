package Go

import "reflect"
import "github.com/qlova/script/language"

func (implementation Implementation) Add(a, b language.Number) language.Number {
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"+"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

func (implementation Implementation) Sub(a, b language.Number) language.Number {
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"-"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

func (implementation Implementation) Mul(a, b language.Number) language.Number {
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"*"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

func (implementation Implementation) Div(a, b language.Number) language.Number {
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"/"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

func (implementation Implementation) Pow(a, b language.Number) language.Number {
	
	/*
func pow(base, exp int) int {
    var result = 1;
    for {
        if (exp & 1 > 0) {
            result *= base
	}
        exp >>= 1
        if (exp == 0) {
            break
	}
        base *= base
    }

    return result
}
*/
	
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"^"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

func (implementation Implementation) Mod(a, b language.Number) language.Number {
	var result language.NewType
	result.Expression = implementation.ExpressionOf(a)+"%"+implementation.ExpressionOf(b)
	return reflect.ValueOf(result).Convert(reflect.TypeOf(a)).Interface().(language.Number)
}

