package Null

import . "github.com/qlova/script/language"


//Returns a Pointer type based of value 'T'.
func (l *language) PointerTo(value Type) Pointer {
	panic("Error in "+Name+".PointerTo(Type): Unimplemented")
	return nil
}
		
//Returns the refernce of the Pointer 'pointer'.
func (l *language) Dereference(pointer Pointer) Type {
	panic("Error in "+Name+".Dereference(Pointer): Unimplemented")
	return nil
}

//Returns a Dynamic type based of value 'T'.
func (l *language) ToDynamic(value Type) Dynamic {
	panic("Error in "+Name+".ToDynamic(Type): Unimplemented")
	return nil
}
		
//Returns a Type cast from value 'T'.
func (l *language) DynamicTo(value Type) Type {
	panic("Error in "+Name+".DynamicTo(Type): Unimplemented")
	return nil
}
		
//Returns Dynamic's type as a Metatype.
func (l *language) DynamicMetatype(value Dynamic) Metatype {
	panic("Error in "+Name+".DynamicMetatype(Dynamic): Unimplemented")
	return nil
}
