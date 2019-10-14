package mirror

import (
	"reflect"
	"unsafe"
)

type structType struct {
	genericType
	fields     fieldNameMap
	jsonFields fieldNameMap
}

func newStructType(typ reflect.Type) *structType {
	fields, jsonFields := newFieldNameMaps(typ)

	return &structType{
		genericType: *newGenericType(typ),
		fields:      fields,
		jsonFields:  jsonFields,
	}
}

func (typ *structType) SetField(object interface{}, fieldName string, value interface{}) {
	field, exists := typ.fields[fieldName]

	if !exists {
		panic("Unknown field: " + fieldName)
	}

	typ.setField(object, field, value)
}

func (typ *structType) SetFieldByJSONTag(object interface{}, fieldName string, value interface{}) {
	field, exists := typ.jsonFields[fieldName]

	if !exists {
		panic("Unknown field: " + fieldName)
	}

	typ.setField(object, field, value)
}

func (typ *structType) setField(object interface{}, field *reflect.StructField, value interface{}) {
	objectPtr := unpack(object).data
	valuePtr := unpack(value).data
	fieldPtr := unsafe.Pointer(uintptr(objectPtr) + field.Offset)
	typedmemmove(unpack(field.Type).data, fieldPtr, valuePtr)
}
