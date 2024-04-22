package accessor_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/accessor"
)

func TestOneOfAccessor(t *testing.T) {
	t.Parallel()

	t.Run("when the one-of value is set", func(t *testing.T) {
		t.Run("it returns the value with ok set to true", func(t *testing.T) {
			testAccessor(
				t,
				(*OneOf).SetFieldA,
				(*OneOf).TryGetFieldA,
				123,
			)

			testAccessor(
				t,
				(*OneOf).SetFieldA,
				(*OneOf).TryGetFieldA,
				0,
			)

			testAccessor(
				t,
				(*OneOf).SetFieldB,
				(*OneOf).TryGetFieldB,
				456,
			)

			testAccessor(
				t,
				(*OneOf).SetFieldB,
				(*OneOf).TryGetFieldB,
				0,
			)

			testAccessor(
				t,
				(*OneOf).SetFieldC,
				(*OneOf).TryGetFieldC,
				"<value>",
			)

			testAccessor(
				t,
				(*OneOf).SetFieldC,
				(*OneOf).TryGetFieldC,
				"",
			)
		})
	})

	t.Run("when the one-of value is set to some other option", func(t *testing.T) {
		t.Run("it returns with ok set to false", func(t *testing.T) {
			m := &OneOf{}
			m.SetFieldB(100)

			if _, ok := m.TryGetFieldA(); ok {
				t.Fatalf("TryGetFieldA() returned ok as true, want false")
			}

			if _, ok := m.TryGetFieldC(); ok {
				t.Fatalf("TryGetFieldC() returned ok as true, want false")
			}
		})
	})
}
