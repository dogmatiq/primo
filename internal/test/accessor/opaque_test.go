package accessor_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/accessor"
)

func TestOneOfAccessor_opaque(t *testing.T) {
	t.Parallel()
	testOneOfAccessorSuite[*OpaqueOneOf](t)
}
