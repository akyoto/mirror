package mirror_test

import (
	"reflect"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/mirror"
)

type A struct {
	B int    `json:"b"`
	C string `json:"c"`
	D *int   `json:"d"`
}

func TestKind(t *testing.T) {
	// Run these statements twice to test the type caching
	for i := 0; i < 2; i++ {
		assert.Equal(t, mirror.TypeOf(true).Kind(), reflect.Bool)
		assert.Equal(t, mirror.TypeOf(1).Kind(), reflect.Int)
		assert.Equal(t, mirror.TypeOf("Hello").Kind(), reflect.String)
		assert.Equal(t, mirror.TypeOf(struct{}{}).Kind(), reflect.Struct)
	}
}

func TestSet(t *testing.T) {
	i := 0
	typ := mirror.TypeOf(i)
	assert.Equal(t, i, 0)
	typ.Set(&i, 1)
	assert.Equal(t, i, 1)
}

func TestSetField(t *testing.T) {
	a := A{}
	typ := mirror.TypeOf(a).(mirror.StructType)
	assert.Equal(t, a.B, 0)
	typ.SetField(&a, "B", 1)
	assert.Equal(t, a.B, 1)
}

func TestSetJSONField(t *testing.T) {
	a := A{}
	typ := mirror.TypeOf(a).(mirror.StructType)
	assert.Equal(t, a.B, 0)
	typ.SetFieldByJSONTag(&a, "b", 1)
	assert.Equal(t, a.B, 1)
}

func BenchmarkSetField(b *testing.B) {
	a := A{}
	typ := mirror.TypeOf(a).(mirror.StructType)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		typ.SetField(&a, "B", 1)
	}
}

func BenchmarkSetFieldByJSONTag(b *testing.B) {
	a := A{}
	typ := mirror.TypeOf(a).(mirror.StructType)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		typ.SetFieldByJSONTag(&a, "b", 1)
	}
}
