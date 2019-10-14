package mirror

import "reflect"

// Type gives you access to information about the data type.
type Type interface {
	Kind() reflect.Kind
	Set(object interface{}, value interface{})
}

// StructType gives you access to structs.
type StructType interface {
	Type
	SetField(object interface{}, fieldName string, value interface{})
}
