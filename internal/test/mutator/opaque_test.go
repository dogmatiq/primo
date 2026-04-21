package mutator_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutator"
)

// TestOpaqueMutatorSkipped verifies that the mutator generator is skipped for
// opaque API messages. The assertion is structural: if the generator emitted
// SetXxx methods, they would conflict with the identically-named SetXxx methods
// that protoc-gen-go itself generates for opaque API messages, causing a
// compile error. Passing this test confirms no such code was generated.
//
// The SetXxx mutators provided by protoc-gen-go itself remain available.
func TestOpaqueMutatorSkipped(t *testing.T) {
	t.Parallel()

	m := &OpaqueMessage{}
	m.SetNum(42)
	m.SetStr("<value>")

	if got := m.GetNum(); got != 42 {
		t.Fatalf("GetNum() = %v, want 42", got)
	}
	if got := m.GetStr(); got != "<value>" {
		t.Fatalf("GetStr() = %q, want %q", got, "<value>")
	}
}
