package reflect

import (
	std "reflect"
)

// Type gives you access to information about the data type.
type Type interface {
	Kind() std.Kind
}
