package mirror

import (
	"reflect"
	"unsafe"
)

type genericType struct {
	reflect.Type
	typePtr unsafe.Pointer
}

func newGenericType(typ reflect.Type) *genericType {
	return &genericType{
		Type:    typ,
		typePtr: unpack(typ).data,
	}
}

func (typ *genericType) Set(object interface{}, value interface{}) {
	objectIFace := unpack(object)
	valueIFace := unpack(value)
	typedmemmove(typ.typePtr, objectIFace.data, valueIFace.data)
}

func (typ *genericType) Elem() Type {
	return upgradeType(typ.Type.Elem())
}
