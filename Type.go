package mirror

import "reflect"

// Type gives you access to information about the data type.
type Type interface {
	Kind() reflect.Kind
}
