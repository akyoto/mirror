package mirror

import "unsafe"

// iface represents Go's internal interface structure.
// An interface object always includes a type and a data pointer.
type iface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

// unpack lets you access the internal interface data for the object.
func unpack(obj interface{}) *iface {
	return (*iface)(unsafe.Pointer(&obj))
}
