package Example

import "github.com/qlova/script/language"

func (implementation Implementation) And(a, b language.Bit) language.Bit {
	panic(implementation.Name()+".And() Unimplemented")
	return nil
}

func (implementation Implementation) Or(a, b language.Bit) language.Bit {
	panic(implementation.Name()+".Or() Unimplemented")
	return nil
}

func (implementation Implementation) Not(b language.Bit) language.Bit {
	panic(implementation.Name()+".Not() Unimplemented")
	return nil
}

func (implementation Implementation) Equals(a, b language.Type) language.Bit {
	panic(implementation.Name()+".Equals() Unimplemented")
	return nil
}

func (implementation Implementation) Smaller(a, b language.Type) language.Bit {
	panic(implementation.Name()+".Smaller() Unimplemented")
	return nil
}

func (implementation Implementation) Greater(a, b language.Type) language.Bit {
	panic(implementation.Name()+".Greater() Unimplemented")
	return nil
}

