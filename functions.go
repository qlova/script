package script

import "bytes"
import "github.com/qlova/script/language"

type Func struct {
	script Script
	internal language.Function

	arguments []Type
	returns Type
}

//Return a new String type with the value s.
func (q Script) Return(value Type) {
	q.indent()
	q.write(q.script.lang.Return(value.LanguageType()))
}

//Return a new Function based on the contents of the provided function.
// Arguments are defined inside the function, with Type.Value().Arg()
func (q Script) Func(f func(), names ...string) Func {
	var name string
	if len(names) > 0 {
		name = names[0]
	}
	
	q.push()
	f()
	var context = q.pop()
	//println(context.body.String())
	
	var Converted = make([]language.Type, len(context.arguments))
	
	for i := range context.arguments {
		Converted[i] = context.arguments[i].LanguageType()
	}
	
	var returns language.Type
	if context.returns != nil {
		returns = context.returns.LanguageType()
	}
	
	statement, function := q.script.lang.Function(name, context.registers, Converted, returns)
	
	//If we are creating a global scope function.
	if name != "" {
		q.head.WriteString(string(statement))
		q.head.Write(context.body.Bytes())
		q.head.WriteString(string(q.script.lang.EndFunction()))
		
		return Func{
			internal: function,
			script: q,
			arguments: context.arguments,
			returns: context.returns,
		}
		
	} else {
		var expression bytes.Buffer
		expression.WriteString(string(statement))
		expression.Write(context.body.Bytes())
		expression.WriteString(string(q.script.lang.EndFunction()))
		return Func{
			internal: function.Register(string(expression.Bytes())).(language.Function),
			script: q,
			arguments: context.arguments,
			returns: context.returns,
		}
	}
	
}

func (f Func) LanguageType() language.Type {
	return f.internal
}

func (f Func) Value() Value {
	return Value{
		script: f.script,
		internal: f.LanguageType(),
	}
}

func (v Value) IsFunc() bool {
	_, ok := v.LanguageType().(language.Function)
	return ok
}

//Get this value as a string or cast to a string.
func (v Value) Func() Func {
	if f, ok := v.internal.(language.Function); ok {
		return Func{
			script: v.script,
			internal: f,

			arguments: v.arguments,
		}
	}
	
	panic("Cannot cast to Function")
	return Func{}
}


func (f Func) Arguments() []Type {
	return f.arguments
}

func (f Func) Call(arguments ...Type) Value {
	
	var Converted = make([]language.Type, len(arguments))
	
	for i := range arguments {
		Converted[i] = arguments[i].LanguageType()
	}
	
	f.script.indent()
	f.script.lang.Run(f.LanguageType().(language.Function), Converted)
	
	return Value{}
}

func (f Func) Run(arguments ...Type) Value {
	
	var Converted = make([]language.Type, len(arguments))
	
	for i := range arguments {
		Converted[i] = arguments[i].LanguageType()
	}
	
	f.script.indent()
	f.script.write(f.script.lang.Run(f.LanguageType().(language.Function), Converted))
	
	return Value{}
}
