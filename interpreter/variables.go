package Interpreter

import (
	"reflect"
	"fmt"
	"math/big"
)
import "github.com/qlova/script/language"

import "github.com/qlova/script/interpreter/internal"

//Returns an Array of type T with 'length' length.
func (l *implementation) Literal(value interface{}) language.Type {
	block := l.loadBlock()

	if b, ok := value.(*big.Int); ok {
		var Address = block.CreateNumber()
		
		var literal = *b
		block.AddInstruction(func() {
			block.SetNumber(Address, &literal)
		})
		
		var result Number
		result.Address = Address
		result.BlockPointer = block

		return result
	}
	
	switch v := value.(type) {
		case rune:
			return Symbol{ IsLiteral: true, Literal: value.(rune) };
			
		case string:
			
			var Address = block.CreateString()
			
			var literal = v
			block.AddInstruction(func() {
				block.SetString(Address, literal)
			})
			
			var result String
			result.Address = Address
			result.BlockPointer = block

			return result
			
		case []Number:
			var Address = block.CreateValue()
			var SliceType = reflect.SliceOf(reflect.TypeOf(big.NewInt(0)))
			var length = len(v)
			block.AddInstruction(func() {
				var Slice = reflect.MakeSlice(SliceType, length, length)
				
				for i := 0; i < length; i++ {
					BlockPointer := v[i].BlockPointer
					Address := v[i].Address
					Slice.Index(i).Set(reflect.ValueOf(BlockPointer.GetNumber(Address)))
				}
				
				block.SetValue(Address, Slice)
			})
			
			var result List
			result.Subtype = Number{}
			result.Address = Address
			result.BlockPointer = block
			return result
	}
	
	var T = reflect.TypeOf(value)
	
	switch T.Kind() {
		
		/*case reflect.Bool:
			return Boolean(fmt.Sprint(value))*/

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, 
		     reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		     reflect.Uint32, reflect.Uint64:
				
			var result Number
			result.Literal = new(big.Int)
			result.Literal.SetString(fmt.Sprint(value), 10)
			return result
		
		default:
			panic("Error in "+Name+".Literal("+fmt.Sprint(T)+"): Unimplemented")
	}
	
	panic("Error in "+Name+".Literal(): Unimplemented")
	return nil
}

//Returns a statement that defines 'name' to be of type 'T' with optional 'value'.
func (l *implementation) Define(name string, value language.Type) (language.Type, language.Statement) {
	var PanicName = "Error in "+Name+".Define("+name+", "+value.Name()+")"
	
	block := l.loadBlock()
	
	switch value.(type) {
		case language.Function:
			var Address = block.CreateFunction()
			f := value.(Function)
			
			if f.IsLiteral {
				literal := f.BlockPointer
				block.AddInstruction(func() {
					block.SetFunction(Address, literal)
				})
			} else {
				var BlockPointer = f.BlockPointer
				var ValueAddress = f.Address
				block.AddInstruction(func() {
					block.SetFunction(Address, BlockPointer.GetFunction(ValueAddress))
				})
			}
	
			n := Function{}
			n.BlockPointer = block
			n.Address = Address
			
			return n, ""
			
		case language.String:
			var Address = block.CreateString()
			
			str := value.(String)
			
			if str.IsLiteral {
				literal := str.Literal
				block.AddInstruction(func() {
					block.SetString(Address, literal)
				})
			} else {
				var BlockPointer = str.BlockPointer
				var ValueAddress = str.Address
				block.AddInstruction(func() {
					block.SetString(Address, BlockPointer.GetString(ValueAddress))
				})
			}
			
			n := String{}
			n.BlockPointer = block
			n.Address = Address
			
			return n, ""
		
		case language.Number:
			var Address = block.CreateNumber()
			
			number := value.(Number)
			
			if number.Literal != nil {
				literal := *number.Literal
				block.AddInstruction(func() {
					block.SetNumber(Address, &literal)
				})
			} else {
				var BlockPointer = number.BlockPointer
				var ValueAddress = number.Address
				block.AddInstruction(func() {
					block.SetNumber(Address, BlockPointer.GetNumber(ValueAddress))
				})
			}
			
			n := Number{}
			n.BlockPointer = block
			n.Address = Address
			
			return n, ""
		
		case language.Boolean:
			var Address = block.CreateBoolean()
			
			boolean := value.(Boolean)
			
			if boolean.Literal != nil {
				b := *boolean.Literal
				block.AddInstruction(func() {
					block.SetBoolean(Address, b)
				})
			} else {
				var BlockPointer = boolean.BlockPointer
				var ValueAddress = boolean.Address
				block.AddInstruction(func() {
					block.SetBoolean(Address, BlockPointer.GetBoolean(ValueAddress))
				})
			}
			
			n := Boolean{}
			n.BlockPointer = block
			n.Address = Address
			
			return n, ""
			
		case language.Array:
			array := value.(Array)
			
			var Address = block.CreateArray(array)
			
			a := Array{}
			a.BlockPointer = block
			a.Address = Address
			a.Size = array.Size
			a.Subtype = array.Subtype
			
			return a, ""
			
		case language.List:
			list := value.(List)
			
			var Address = block.CreateValue()
			
			a := List{}
			a.BlockPointer = block
			a.Address = Address
			a.Size = list.Size
			a.Subtype = list.Subtype
			
			return a, ""
			
		case language.Symbol,
			language.Custom, language.Stream,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic,language.Metatype:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return nil, ""
}

//Returns a Statement that sets the type 'T' variable 'name' to be set to 'value'.
func (l *implementation) Set(T language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Set("+T.Name()+", "+value.Name()+")"
	
	if !T.SameAs(value) {
		panic(PanicName+": type mismatch!")
	}
	
	switch T.(type) {
		case language.Number:
			
			var Address = T.(Number).Address
			var Block = T.(Number).BlockPointer
			
			number := value.(Number)
			
			if number.Literal != nil {
				b := number.Literal
				Block.AddInstruction(func() {
					Block.SetNumber(Address, b)
				})
			} else {
				var BlockPointer = number.BlockPointer
				var ValueAddress = number.Address
				Block.AddInstruction(func() {
					Block.SetNumber(Address, BlockPointer.GetNumber(ValueAddress))
				})
			}
		
		case language.Boolean:
			
			var Address = T.(Boolean).Address
			var Block = T.(Boolean).BlockPointer
			
			boolean := value.(Boolean)
			
			if boolean.Literal != nil {
				b := *boolean.Literal
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, b)
				})
			} else {
				var BlockPointer = boolean.BlockPointer
				var ValueAddress = boolean.Address
				Block.AddInstruction(func() {
					Block.SetBoolean(Address, BlockPointer.GetBoolean(ValueAddress))
				})
			}
		
		case language.Symbol, language.String, 
			language.Custom, language.Stream, language.List, language.Array, 
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return ""
}

//Returns the Type at 'index' of 'T'.
func (l *implementation) Index(T language.Type, index language.Type) language.Type {
	var PanicName = "Error in "+Name+".Index("+T.Name()+", "+index.Name()+")"
	
	switch T.(type) {
		case language.Array:
			block := l.loadBlock()

			var Length = T.(Array).Length()

			switch T.(Array).SubType().(type) {
				case language.Number:
					
					var Address = block.CreateNumber()
					
					var BlockPointer = T.(Array).BlockPointer
					var Pointer = T.(Array).Address.Address()
					
					if index.(Number).Literal == nil {
						
						var IndexBlock = index.(Number).BlockPointer
						var Index = index.(Number).Address
						
						block.AddInstruction(func() {
							block.SetNumber(Address, BlockPointer.GetNumber(internal.NumberAddress(
								Pointer + (int(IndexBlock.GetNumber(Index).Int64()) % Length),
							)))
						})
					} else {
						var Index = Pointer + (int(index.(Number).Literal.Int64()) % Length)
						block.AddInstruction(func() {
							block.SetNumber(Address, BlockPointer.GetNumber(internal.NumberAddress(Index)))
						})
					}
					
					var result Number
					result.BlockPointer = block
					result.Address = Address
					return result
					
				default:
					panic("Error in Block.Index("+T.Name()+", Number): Unsupported subtype for Array: ")
			}
		
		case language.Boolean, language.Number, language.Symbol, language.String, 
			language.Custom, language.Stream, language.List,
			language.Table, language.Error, language.Float, language.Pointer, 
			language.Dynamic, language.Function, language.Metatype, language.FunctionType:
		
		panic(PanicName+": Unimplented")
			
		default:
			panic(PanicName+": Invalid Type")
	}
	
	return nil
}

//Returns a statement that modifies type T at 'index' to be 'value'.
func (l *implementation) Modify(T language.Type, index language.Type, value language.Type) language.Statement {
	var PanicName = "Error in "+Name+".Modify("+T.Name()+", "+index.Name()
	PanicName += ", "+value.Name()+")"
	
	switch T.(type) {
		
		case language.Array:
			block := l.loadBlock()

			var Length = T.(Array).Length()

			switch T.(Array).SubType().(type) {
				case language.Number:

					var BlockPointer = T.(Array).BlockPointer
					var Pointer = int(T.(Array).Address)
					
					if index.(Number).Literal == nil {
						var IndexBlock = index.(Number).BlockPointer
						var Index = index.(Number).Address
						
						if value.(Number).Literal == nil {
							var ValueBlock = value.(Number).BlockPointer
							var Value =  value.(Number).Address
							block.AddInstruction(func() {
								BlockPointer.SetNumber(
									internal.NumberAddress(Pointer + (int(IndexBlock.GetNumber(Index).Int64()) % Length)), 
									ValueBlock.GetNumber(Value),
								)
							})
						} else {
							var literal = *value.(Number).Literal
							block.AddInstruction(func() {
								BlockPointer.SetNumber(internal.NumberAddress(Pointer + (int(IndexBlock.GetNumber(Index).Int64()) % Length)), &literal)
							})
						}
					} else {
						
						var Index = Pointer + int(index.(Number).Literal.Int64())
						if value.(Number).Literal == nil {
							var ValueBlock = value.(Number).BlockPointer
							var Value =  value.(Number).Address
							block.AddInstruction(func() {
								BlockPointer.SetNumber(internal.NumberAddress(Index), ValueBlock.GetNumber(Value))
							})
						} else {
							var literal = *value.(Number).Literal
							block.AddInstruction(func() {
								BlockPointer.SetNumber(internal.NumberAddress(Index), &literal)
							})
						}
						
					}
					return ""
					
				default:
					panic("Error in Block.Index("+T.Name()+", Number): Unsupported subtype for Array: ")
			}
		
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
 
func (l *implementation) GetVariable(T language.Type) language.Type {
	
}
