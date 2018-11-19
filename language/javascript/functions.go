package Javascript

import "github.com/qlova/script/language"

//Returns a Stream connected to 'function' called with 'arguments' starting on another thread, netowork, coroutine or process.
func (l *implementation) Thread(functiom language.Function, arguments []language.Type) language.Stream {
	panic("Error in "+Name+".Thread(Function, []Type): Unimplemented")
	return nil
}

type Function struct {
	language.FunctionType

	External bool
}

//Returns a Function that refers to external 'name'.
func (l *implementation) External(name language.String) (language.Function) {
	l.Import("os/exec")
	l.AddHelper(ErrorHelper)
	
	l.AddHelper(`func External(command string) func(args ...string) []byte {
	return func(args ...string) []byte {
		var cmd = exec.Command(command, args...)
		result, err := cmd.CombinedOutput()
		if err != nil {
			errors = append(errors, Error{Line:line, File:filename, Code: 1, Message: err.Error()})
		}
		return result
	}
}

`)
	
	var f Function
	f.External = true
	f.Expression = "External("+l.GetExpression(name)+")"
	f.Rets = String("")
	
	return f
}

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *implementation) Function(name string, names []string, arguments []language.Type, returns language.Type) (language.Function, language.Statement) {
	var ArgsString = ""
	
	var f Function
	f.Expression = name
	if len(arguments) == 0 {
		f.Args = nil 
	} else {
		f.Args = make([]language.Type, len(arguments))
		for i := range arguments {
			f.Args[i] = GetVariable(names[i], arguments[i])
			ArgsString += names[i]+" "+l.GoTypeOf(arguments[i])
			if i < len(arguments)-1 {
				ArgsString += ","
			}
		}
	}
	f.Rets = returns
	
	return f, language.Statement("func "+name+"("+ArgsString+") "+l.GoTypeOf(returns)+" {\n")
	
	panic("Error in "+Name+".Function("+name+", []Type, []Type): Unimplemented")
	return nil, ""
}

//Returns a Statement that begins the function 'name' with 'arguments' and 'returns'.
func (l *implementation) UpdateFunction(f language.Function, name string, names []string, arguments []language.Type, returns language.Type) (language.Function, language.Statement) {
	return l.Function(name, names, arguments, returns)
}

//Returns the resulting Type from calling 'function' with 'arguments'
func (l *implementation) Call(function language.Function, arguments []language.Type) language.Type {
	
	if len(arguments) == 0 {
		return GetVariable(l.GetExpression(function)+"()", function.Returns())
	}
	
	var result = l.GetExpression(function) + "("
	for i := range arguments {
		result += l.GetExpression(arguments[i])
		
		if i < len(arguments) - 1 {
			result +=  ","
		}
	}
	result += ")"
	
	if function.(Function).External {
		return GetVariable("string("+result+")", String(""))
	}

	return GetVariable(result, function.Returns())
}

//Returns a Statement that closes the last function.
func (l *implementation) EndFunction() language.Statement {
	return language.Statement("}\n\n")
}
		
//Returns a Statement that calls 'function' with 'arguments'.
func (l *implementation) Run(function language.Function, arguments []language.Type) language.Statement {

	if function.(Function).External {
		l.Import("os")
		
		if len(arguments) == 0 {
			return language.Statement("os.Stdout.Write("+l.GetExpression(function)+"())\n")
		}
			
		var result = l.GetExpression(function) + "("
		for i := range arguments {
			result += l.GetExpression(arguments[i])
			
			if i < len(arguments) - 1 {
				result +=  ","
			}
		}
		result += ")"
		
		return language.Statement("os.Stdout.Write("+result+")\n")
	}
	
	return language.Statement(l.GetExpression(l.Call(function, arguments))+"\n")
}

//Returns a Statement that returns T from the current function.
func (l *implementation) Return(T language.Type) language.Statement {
	return language.Statement("return "+l.GetExpression(T)+"\n")
}
