package reflect_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/reflect"
)

func TestTypeOf(t *testing.T) {
	typ := reflect.TypeOf(1)
	assert.NotNil(t, typ)
}
