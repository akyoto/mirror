package mirror

import (
	"reflect"
	"unsafe"
)

type structType struct {
	genericType
	jsonIndices fieldIndexMap
}

func newStructType(typ reflect.Type) *structType {
	return &structType{
		genericType: *newGenericType(typ),
		jsonIndices: newJSONFieldMap(typ),
	}
}

func (typ *structType) SetField(object interface{}, fieldName string, value interface{}) {
	field, exists := typ.FieldByName(fieldName)

	if !exists {
		panic("Unknown field: " + fieldName)
	}

	objectPtr := unpack(object).data
	valuePtr := unpack(value).data
	offset := field.Offset
	fieldPtr := unsafe.Pointer(uintptr(objectPtr) + offset)
	typedmemmove(unpack(field.Type).data, fieldPtr, valuePtr)
}
