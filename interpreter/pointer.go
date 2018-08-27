package Interpreter

import "github.com/qlova/script/language"


//Returns a Pointer type based of value 'T'.
func (l *implementation) PointerTo(value language.Type) language.Pointer {
	panic("Error in "+Name+".PointerTo(Type): Unimplemented")
	return nil
}
		
//Returns the refernce of the Pointer 'pointer'.
func (l *implementation) Dereference(pointer language.Pointer) language.Type {
	panic("Error in "+Name+".Dereference(Pointer): Unimplemented")
	return nil
}

//Returns a Dynamic type based of value 'T'.
func (l *implementation) ToDynamic(value language.Type) language.Dynamic {
	panic("Error in "+Name+".ToDynamic(Type): Unimplemented")
	return nil
}
		
//Returns a Type cast from value 'T'.
func (l *implementation) DynamicTo(value language.Type) language.Type {
	panic("Error in "+Name+".DynamicTo(Type): Unimplemented")
	return nil
}
		
//Returns Dynamic's type as a Metatype.
func (l *implementation) DynamicMetatype(value language.Dynamic) language.Metatype {
	panic("Error in "+Name+".DynamicMetatype(Dynamic): Unimplemented")
	return nil
}
