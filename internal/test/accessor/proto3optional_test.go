package accessor_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/accessor"
)

// TestOneOfAccessor_proto3optional verifies that a real oneof declared
// alongside proto3 optional fields still generates TryGetXxx accessor methods.
func TestOneOfAccessor_proto3optional(t *testing.T) {
	t.Parallel()

	t.Run("when the one-of value is set", func(t *testing.T) {
		t.Run("it returns the value with ok set to true", func(t *testing.T) {
			m := &Proto3OptionalMessage{}
			m.SetChoiceA("<value>")
			if v, ok := m.TryGetChoiceA(); !ok || v != "<value>" {
				t.Fatalf("TryGetChoiceA() = (%q, %v), want (%q, true)", v, ok, "<value>")
			}

			m = &Proto3OptionalMessage{}
			m.SetChoiceB(123)
			if v, ok := m.TryGetChoiceB(); !ok || v != 123 {
				t.Fatalf("TryGetChoiceB() = (%v, %v), want (123, true)", v, ok)
			}
		})
	})

	t.Run("when the one-of value is set to some other option", func(t *testing.T) {
		t.Run("it returns with ok set to false", func(t *testing.T) {
			m := &Proto3OptionalMessage{}
			m.SetChoiceB(123)
			if _, ok := m.TryGetChoiceA(); ok {
				t.Fatal("TryGetChoiceA() returned ok as true, want false")
			}
		})
	})
}
