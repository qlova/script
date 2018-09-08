package compiler

import "github.com/qlova/script"

type Variable struct {
	Type
	Defined, Protected, Modified, Embedded bool
	
	Index int
	List string
	
	DefinedAtLineNumber int
	DefinedAtLine string
}

func (c *Compiler) Define(name string, t script.Type) {
	c.SetVariable(name, c.Script.Define(name, t))
}

func (c *Compiler) Variable(name string) Type {
	return c.GetVariable(name).Type
}


func (c *Compiler) SetVariable(name string, t Type) {
	if _, ok := c.GetScope().Variables[name]; ok {
		c.RaiseError(Translatable{
			English: name+" already defined!",
		})
	}
	c.GetScope().Variables[name] = Variable{ 
		Type: t,
		Defined: true,
		DefinedAtLineNumber: c.Scanners[len(c.Scanners)-1].Line,
		DefinedAtLine: c.CurrentLines[len(c.Scanners)-1],
	}
}

func (c *Compiler) SetGlobal(name string, t Type) {
	if _, ok := c.GlobalScope.Variables[name]; ok {
		c.RaiseError(Translatable{
			English: name+" already defined!",
		})
	}
	c.GlobalScope.Variables[name] = Variable{ 
		Type: t,
		Defined: true,
		DefinedAtLineNumber: c.Scanners[len(c.Scanners)-1].Line,
		DefinedAtLine: c.CurrentLines[len(c.Scanners)-1],
	}
}


func (c *Compiler) UpdateVariable(name string, t Type) {
	for i:=len(c.Scope)-1; i>=0; i-- {
		if v, ok := c.Scope[i].Variables[name]; ok {
			v.Type = t
			c.Scope[i].Variables[name] = v
		}
	} 
}

func (c *Compiler) GetVariable(name string) Variable {
	for i:=len(c.Scope)-1; i>=0; i-- {
		if v, ok := c.Scope[i].Variables[name]; ok {
			return v
		}
	} 

	return Variable{}
}
