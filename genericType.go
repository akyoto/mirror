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

func (typ *genericType) Set(obj interface{}, val interface{}) {
	objIFace := unpack(obj)
	valIFace := unpack(val)
	typedmemmove(typ.typePtr, objIFace.data, valIFace.data)
}
