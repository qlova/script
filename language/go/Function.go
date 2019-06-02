package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Function(name string, registers []string, arguments []language.Type, returns language.Type) (language.Statement, language.Function) {
	
	if name != "" {
		
		var definition = "func "+name+"("
		
		for i := range registers {
			definition += registers[i]+" "+arguments[i].Name()
			if i < len(registers)-1 {
				definition += ","
			}	
		}
		
		var returndefinition string
		if returns != nil {
			returndefinition = returns.Name()
		}

		
		return language.Statement(definition+") "+returndefinition+" {\n"), Function{
			Expression: language.Statement(name),
			Subtype: returns,
		}
	}
	
	panic(implementation.Name()+".Function() Unimplemented")
	return language.Statement(""), nil
}

func (implementation Implementation) EndFunction() language.Statement {
	return language.Statement("}\n\n")
}

func (implementation Implementation) Call(f language.Function, arguments []language.Type) language.Type {
	
	var call = implementation.ExpressionOf(f)+"("
	
	for i := range arguments {
		call += arguments[i].Raw()
		if i < len(arguments)-1 {
			call += ","
		}
	}
	
	call += ")"

	return f.(Function).Subtype.Register(string(call))
}

func (implementation Implementation) Run(f language.Function, arguments []language.Type) language.Statement {
	var call = implementation.ExpressionOf(f)+"("
	
	for i := range arguments {
		call += arguments[i].Raw()
		if i < len(arguments)-1 {
			call += ","
		}
	}
	
	return call+")\n"
}

func (implementation Implementation) Return(value language.Type) language.Statement {
	if value == nil {
		return language.Statement("return\n")
	}
	return language.Statement("return "+value.Raw()+"\n")
}

