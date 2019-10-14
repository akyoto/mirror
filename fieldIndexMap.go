package mirror

import (
	"reflect"
	"strings"
)

// fieldIndexMap maps a field name to its position in the struct (index).
type fieldIndexMap = map[string]int

// newJSONFieldMap returns a map of field names mapped to their index.
func newJSONFieldMap(structType reflect.Type) fieldIndexMap {
	fields := fieldIndexMap{}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Use the json tag if available
		jsonName := field.Tag.Get("json")

		if jsonName != "" {
			comma := strings.Index(jsonName, ",")

			if comma != -1 {
				jsonName = jsonName[:comma]
			}

			fields[jsonName] = i
			continue
		}

		// Otherwise just use the normal field name
		fields[field.Name] = i
	}

	return fields
}
