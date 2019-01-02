package typesystem

//HTuple represents a tuple of types
type HTuple struct {
	types []*IHType
}

//DefineTuple creates a new tuple definition
func DefineTuple(types []*IHType) *HTuple {
	return &HTuple{types: types}
}

//Is returns if every type of this tuple is another type
func (htuple HTuple) Is(other IHType) bool {
	oTpl, ok := other.(*HTuple)
	if !ok {
		if len(htuple.types) == 1 {
			return (*htuple.types[0]).Is(other)
		}
		return false
	}
	if len(htuple.types) != len(oTpl.types) {
		return false
	}
	for i, htype := range htuple.types {
		if !(*htype).Is(*oTpl.types[i]) {
			return false
		}
	}
	return true
}

//Pointer represents a pointer to an IHType
type Pointer struct {
	htype *IHType
}

//DefinePointer returns a pointer object pointing to a specific IHType
func DefinePointer(htype *IHType) Pointer {
	return Pointer{htype: htype}
}

//Is represents if this pointer is equal to another
func (ptr Pointer) Is(other IHType) bool {
	return ptr.htype == &other
}

//HArray represents an array of a given type
type HArray struct {
	htype *IHType
}

//DefineArray defines an array type
func DefineArray(htype *IHType) HArray {
	return HArray{htype: htype}
}

//Is returns if this array type is a subclass of another
func (harr HArray) Is(other IHType) bool {
	arr, ok := other.(HArray)
	if !ok { //if the other type isn't even an array, then false
		return false
	}
	return (*harr.htype).Is(*arr.htype)
}
