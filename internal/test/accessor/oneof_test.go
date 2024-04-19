package accessor_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/accessor"
)

func TestOneOfAccessor(t *testing.T) {
	t.Parallel()

	t.Run("it returns the set value and ok as true to signify the presence of the value", func(t *testing.T) {
		testAccessor(
			t,
			(*OneOf).SetFieldA,
			(*OneOf).TryFieldA,
			123,
		)

		testAccessor(
			t,
			(*OneOf).SetFieldA,
			(*OneOf).TryFieldA,
			0,
		)

		testAccessor(
			t,
			(*OneOf).SetFieldB,
			(*OneOf).TryFieldB,
			456,
		)

		testAccessor(
			t,
			(*OneOf).SetFieldB,
			(*OneOf).TryFieldB,
			0,
		)

		testAccessor(
			t,
			(*OneOf).SetFieldC,
			(*OneOf).TryFieldC,
			"<value>",
		)

		testAccessor(
			t,
			(*OneOf).SetFieldC,
			(*OneOf).TryFieldC,
			"",
		)
	})

	t.Run("it ok as false to signify the absence of the value", func(t *testing.T) {
		m := &OneOf{}

		if _, ok := m.TryFieldA(); ok {
			t.Fatalf("TryFieldA() returned ok as true, want false")
		}

		if _, ok := m.TryFieldB(); ok {
			t.Fatalf("TryFieldB() returned ok as true, want false")
		}

		if _, ok := m.TryFieldC(); ok {
			t.Fatalf("TryFieldC() returned ok as true, want false")
		}
	})
}
