package reflect

import "sync"

var typeCache sync.Map

// TypeOf returns the (possibly cached) type information for the object.
func TypeOf(obj interface{}) Type {
	key := uintptr(unpack(obj).typ)
	typeObj, found := typeCache.Load(key)

	if found {
		return typeObj.(Type)
	}

	return nil
}
