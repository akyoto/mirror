package mirror

import (
	"reflect"
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
