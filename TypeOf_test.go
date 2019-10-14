package mirror_test

import (
	"reflect"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/mirror"
)

func TestTypeOf(t *testing.T) {
	typ := mirror.TypeOf(1)
	assert.NotNil(t, typ)
	assert.Equal(t, typ.Kind(), reflect.Int)
}
