package Go

import "strconv"
import "github.com/qlova/script/language"

func (implementation Implementation) ArrayOf(t language.Type, length int) language.Array {	
	return Array{
		Length: length,
		Subtype: t,
		Expression: language.Statement(`[`+strconv.Itoa(length)+`]`+t.Name()+"{}"),
	}
}

func (implementation Implementation) Join(a, b language.Type) language.Type {
	
	switch a.(type) {
		case String:
			switch b.(type) {
				case String:
					return String{
						Expression: "("+a.Raw()+"+"+b.Raw()+")",
					}
			}
	}
	
	panic(implementation.Name()+".Join() Unimplemented")
	return nil
}

func (implementation Implementation) Length(t language.Type) language.Integer {
	return Integer{
		Expression: `(len(`+t.Raw()+`))`,
	}
}

func (implementation Implementation) Push(value language.Type, list language.Type) language.Statement {
	return list.Raw()+" = append("+list.Raw()+", "+value.Raw()+")"
}

func (implementation Implementation) Pop(list language.Type) language.Type {
	panic(implementation.Name()+".Pop() Unimplemented")
	return nil
}

func (implementation Implementation) TableOf(t language.Type) language.Table {
	return Table{
		Subtype: t,
		Expression: language.Statement(`map[string]`+t.Name()),
	}
}

func (implementation Implementation) ListOf(t language.Type) language.List {
	return List{
		Subtype: t,
		Expression: language.Statement(`[]`+t.Name()),
	}
}

