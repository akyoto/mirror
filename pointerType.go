package mirror

import "reflect"

type pointerType struct {
	genericType
}

func newPointerType(typ reflect.Type) *pointerType {
	return &pointerType{
		genericType: *newGenericType(typ),
	}
}
