package mirror

import (
	"reflect"
	"strings"
)

// fieldNameMap maps a field name to its field.
type fieldNameMap = map[string]*reflect.StructField

// newFieldNameMaps returns field names mapped to their field for normal fields and JSON tags.
func newFieldNameMaps(structType reflect.Type) (fieldNameMap, fieldNameMap) {
	fields := fieldNameMap{}
	jsonFields := fieldNameMap{}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Save the normal field name
		fields[field.Name] = &field

		// Save the json tag if available
		jsonName := field.Tag.Get("json")

		if jsonName != "" {
			comma := strings.Index(jsonName, ",")

			if comma != -1 {
				jsonName = jsonName[:comma]
			}

			jsonFields[jsonName] = &field
		} else {
			jsonFields[field.Name] = &field
		}
	}

	return fields, jsonFields
}
