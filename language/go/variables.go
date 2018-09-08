package Go

import (
	"fmt"
	"reflect"
	"strconv"
	"math/big"
)
import "github.com/qlova/script/language"

//Returns an Array of type T with 'length' length.
func (l *implementation) TypeStringOf(T reflect.Type) string {
	switch T.Kind() {		
		case reflect.Bool:
			return "bool"

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, 
		     reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		     reflect.Uint32, reflect.Uint64:
				
			return "*big.Int"
		
		case reflect.String:
			return "string"
			
		case reflect.Array:
			return "["+fmt.Sprint(T.Len())+"]"+l.TypeStringOf(T.Elem())

		default:
			panic("Error in "+Name+".TypeStringOf("+fmt.Sprint(T)+"): Unimplemented")
	}
	
	panic("Error in "+Name+".TypeStringOf("+fmt.Sprint(T)+"): Unimplemented")
	return ""
}


//Returns an Array of type T with 'length' length.
func (l *implementation) Literal(value interface{}) language.Type {
	
	if b, ok := value.(*big.Int); ok {
		return Number{Literal:b}
	}
	
	switch value.(type) {
		case rune:
			return Symbol(strconv.QuoteRune(value.(rune)));
	}
	
	var T = reflect.TypeOf(value)
	
	switch T.Kind() {
		case reflect.Bool:
			return Boolean(fmt.Sprint(value))

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, 
		     reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		     reflect.Uint32, reflect.Uint64:
				
			var result Number
			result.Literal = new(big.Int)
			result.Literal.SetString(fmt.Sprint(value), 10)
			return result
		
		case reflect.String:
			return String(strconv.Quote(value.(string)))
			
		default:
			panic("Error in "+Name+".Literal("+fmt.Sprint(T)+"): Unimplemented")
	}
	
	panic("Error in "+Name+".Literal(Type, Number): Unimplemented")
	return nil
}


func GetVariable(name string, T language.Type) language.Type {
	
	switch T.(type) {
		case language.Number:
			return Number{Expression: name}
			
		case language.String:
			return String(name)
		
		case language.Boolean:
			return Boolean(name)
			
		case language.Function:
			var ret = Function{}
			ret.Expression = name
			return ret
		
		case language.Array:
			var ret = Array{}
			ret.Expression = name
			return ret
			
		case language.Symbol, 
			language.Custom, language.Stream, language.List,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Metatype:
		
		var PanicName = "Error in "+Name+".GetVariable("+T.Name()+")"
		panic(PanicName+": Unimplented")
			
		default:
			var PanicName = "Error in "+Name+".GetVariable("+T.Name()+")"
			panic(PanicName+": Invalid Type")
	}
}

//Returns a statement that defines 'name' to be of type 'value', initialised to 'value'.
func (l *implementation) Define(name string, value language.Type) (language.Type, language.Statement) {

	var statement = "var "+name+" = "+l.GetExpression(value)+"\n"
	var variable = GetVariable(name, value)

	return variable, language.Statement(statement)
}

//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
func (l *implementation) Set(T language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Set("+T.Name()+", "+value.Name()+")"
	var result string
	
	switch T.(type) {
		case language.Boolean, language.Number:
			result = l.GetExpression(T) + " = " + l.GetExpression(value) + "\n"
		
		case language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return language.Statement(result)
}

//Returns the Type at 'index' of 'T'.
func (l *implementation) Index(T language.Type, index language.Type) language.Type {
	var PanicName = "Error in "+Name+".Index("+T.Name()+", "+index.Name()+")"
	var result language.Type
	
	switch T.(type) {
		case language.Array:
			
			l.AddHelper(NumberToInt64)
			
			var expression = l.GetExpression(T) + "["+l.GetInt(l.Mod(index.(Number), l.Length(T)))+"]"
			
			switch T.(Array).SubType().(type) {
				case language.Number:
					result = Number{Expression: expression}
				default:
					panic(PanicName+": Unimplented")
			}
			
		
		case language.Boolean, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return result
}

func (l *implementation) GetInt(T language.Type) string {
	Number := T.(Number)
	
	l.AddHelper(NumberToInt64)
	
	if Number.Literal == nil {
		return "int("+Number.Expression+".Int64())"
	} else {
		return Number.Literal.String()
	}
}

//Returns a statement that modifies type T at 'index' to be 'value'.
func (l *implementation) Modify(T language.Type, index language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Modify("+T.Name()+", "+index.Name()
	PanicName += ", "+value.Name()+")"
	
	switch T.(type) {
		case language.Array:
			return language.Statement( l.GetExpression(T) + "["+l.GetInt(index)+"] = "+l.GetExpression(value)+"\n" )
		
		case language.Boolean, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}
