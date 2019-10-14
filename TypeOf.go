package mirror

import (
	"reflect"
	"sync"
)

var typeCache sync.Map

// TypeOf returns the (possibly cached) type information for the object.
func TypeOf(obj interface{}) Type {
	key := uintptr(unpack(obj).typ)
	typeObj, found := typeCache.Load(key)

	if found {
		return typeObj.(Type)
	}

	typ := wrapType(reflect.TypeOf(obj))
	typeCache.Store(key, typ)
	return typ
}

func upgradeType(typ reflect.Type) Type {
	if typ == nil {
		return nil
	}

	key := uintptr(unpack(typ).data)
	typeObj, found := typeCache.Load(key)

	if found {
		return typeObj.(Type)
	}

	newTyp := wrapType(typ)
	typeCache.Store(key, newTyp)
	return newTyp
}

func wrapType(typ reflect.Type) Type {
	switch typ.Kind() {
	case reflect.Struct:
		return newStructType(typ)
	case reflect.Ptr:
		return newPointerType(typ)
	default:
		return newGenericType(typ)
	}
}
