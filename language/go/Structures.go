package Go

import "github.com/qlova/script/language"

func (implementation Implementation) Index(structure, index language.Type) language.Type {
	
	switch t := structure.(type) {
		case Array:
			return t.Subtype.Register(string(implementation.ExpressionOf(structure)+"["+implementation.ExpressionOf(index)+" % len("+implementation.ExpressionOf(structure)+")]"))
	}
	
	panic(implementation.Name()+".Index("+structure.Name()+", "+index.Name()+") Unimplemented")
	return nil
}

func (implementation Implementation) Modify(structure, index, value language.Type) language.Statement {
	
	switch structure.(type) {
		case Array:
			return implementation.ExpressionOf(structure)+"["+implementation.ExpressionOf(index)+" % len("+implementation.ExpressionOf(structure)+")] = "+implementation.ExpressionOf(value)+"\n"
	}
	
	panic(implementation.Name()+".Modify("+structure.Name()+", "+index.Name()+", "+value.Name()+") Unimplemented")
	return language.Statement("")
}

