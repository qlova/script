package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Type(name string, registers []string, elements []language.Type) language.Statement {
	if len(registers) != len(elements) {
		panic(implementation.Name()+".Type() register elements length mismatch")
	}
	
	var definition string
		definition += "type "+name+" struct {\n"
		for i := range registers {
			definition += "\t"+registers[i]+" "+elements[i].Name()+"\n"
		}
		definition += "}\n"

	return language.Statement(definition)
}

func (implementation Implementation) Method(t string, name string, registers []string, arguments []language.Type, returns language.Type) language.Statement {
	
	var definition string
		definition += "func (this "+t+") "+name+"("
		for i := range registers {
			definition += registers[i]+" "+arguments[i].Name()
			if i < len(registers)-1 {
				definition += ","
			}
		}
		definition += ")"
		definition += returns.Name()
		definition += ") {\n"

	return language.Statement(definition)
}

func (implementation Implementation) This() language.Type {
	panic(implementation.Name()+".This() Unimplemented")
	return nil
}

func (implementation Implementation) New(name string) language.Type {
	return language.NewType{
		Custom: name,
		Expression: language.Statement(name+"{}"),
	}
}

func (implementation Implementation) Invoke(t language.Type, method string, arguments []language.Type) language.Type {
	panic(implementation.Name()+".Invoke() Unimplemented")
	return nil
}

func (implementation Implementation) Execute(t language.Type, method string, arguments []language.Type) language.Statement {
	panic(implementation.Name()+".Execute() Unimplemented")
	return language.Statement("")
}

func (implementation Implementation) EndMethod() language.Statement {
	return language.Statement("}\n")
}

