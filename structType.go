package mirror

import (
	"reflect"
)

type structType struct {
	genericType
}

func newStructType(typ reflect.Type) *structType {
	return &structType{
		genericType: *newGenericType(typ),
	}
}
