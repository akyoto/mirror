package mirror_test

import (
	"reflect"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/mirror"
)

func TestTypeOf(t *testing.T) {
	// Run these statements twice to test the type caching
	for i := 0; i < 2; i++ {
		assert.Equal(t, mirror.TypeOf(true).Kind(), reflect.Bool)
		assert.Equal(t, mirror.TypeOf(1).Kind(), reflect.Int)
		assert.Equal(t, mirror.TypeOf("Hello").Kind(), reflect.String)
		assert.Equal(t, mirror.TypeOf(struct{}{}).Kind(), reflect.Struct)
	}
}
