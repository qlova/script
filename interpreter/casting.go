package Interpreter

import "math/big"
import "github.com/qlova/script/language"


func (l *implementation) NumberToString(number Number) String {
	if number.Literal != nil {
		return String{IsLiteral: true, Literal: number.Literal.String()}
	}
	
	Block := l.loadBlock()

	var Address = Block.CreateString()
	var s = String{}
	s.Address = Address
	s.BlockPointer = Block 
	
	Block.AddInstruction(func() {
		Block.SetString(Address, number.BlockPointer.GetNumber(number.Address).String())
	})
	
	return s
}

func (l *implementation) SymbolToString(symbol Symbol) String {
	if symbol.IsLiteral {
		return String{IsLiteral: true, Literal: string(symbol.Literal)}
	}
	
	Block := l.loadBlock()

	var Address = Block.CreateString()
	var s = String{}
	s.Address = Address
	s.BlockPointer = Block 
	
	Block.AddInstruction(func() {
		Block.SetString(Address, string(symbol.BlockPointer.GetSymbol(symbol.Address)))
	})
	
	return s
}

//Returns Type cast to String.
func (l *implementation) ToString(T language.Type) language.String {
	
	switch T.(type) {
		case String:
			return T.(String)
			
		case Number:
			return l.NumberToString(T.(Number))
			
		case Symbol:
			return l.SymbolToString(T.(Symbol))
			
	}
	
	panic("Error in "+Name+".ToString("+T.Name()+"): Unimplemented")
	return nil
}

func (l *implementation) SymbolToNumber(symbol Symbol) Number {
	if symbol.IsLiteral {
		return Number{Literal: big.NewInt(int64(symbol.Literal))}
	}
	
	Block := l.loadBlock()

	var Address = Block.CreateNumber()

	var s = Number{}
	s.Address = Address
	s.BlockPointer = Block 
	
	Block.AddInstruction(func() {
		Block.SetNumber(Address, big.NewInt(int64(symbol.BlockPointer.GetSymbol(symbol.Address))))
	})
	
	return s
}
		
//Returns Type cast to Number.
func (l *implementation) ToNumber(T language.Type) language.Number {
	
	switch T.(type) {
		case Number:
			return T.(Number)
			
		case Symbol:
			return l.SymbolToNumber(T.(Symbol))
			
	}
	
	panic("Error in "+Name+".ToNumber("+T.Name()+"): Unimplemented")
	return nil
}

func (l *implementation) NumberToSymbol(number Number) Symbol {
	if number.Literal != nil {
		return Symbol{IsLiteral: true, Literal: rune(number.Literal.Int64())}
	}
	
	Block := l.loadBlock()

	var Address = Block.CreateSymbol()

	var s = Symbol{}
	s.Address = Address
	s.BlockPointer = Block 
	
	Block.AddInstruction(func() {
		Block.SetSymbol(Address, rune(number.BlockPointer.GetNumber(number.Address).Int64() ))
	})
	
	return s
}
		
//Returns Type cast to Number.
func (l *implementation) ToSymbol(T language.Type) language.Symbol {
	
	switch T.(type) {
		case Symbol:
			return T.(Symbol)
			
		case Number:
			return l.NumberToSymbol(T.(Number))
			
	}
	
	panic("Error in "+Name+".ToSymbol("+T.Name()+"): Unimplemented")
	return nil
}
