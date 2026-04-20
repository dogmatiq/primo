package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

// TestEnumSwitch_opaque verifies that Switch_/Map_ functions are generated
// for opaque API enums.
func TestEnumSwitch_opaque(t *testing.T) {
	t.Parallel()

	t.Run("it calls the function associated with the enum value", func(t *testing.T) {
		t.Parallel()

		called := false
		Switch_OpaqueDirection(
			OpaqueDirection_OPAQUE_DIRECTION_LEFT,
			func(OpaqueDirection_OPAQUE_DIRECTION_UNSPECIFIED_Case) {
				panic("unexpected UNSPECIFIED case")
			},
			func(OpaqueDirection_OPAQUE_DIRECTION_LEFT_Case) { called = true },
			func(OpaqueDirection_OPAQUE_DIRECTION_RIGHT_Case) {
				panic("unexpected RIGHT case")
			},
		)
		if !called {
			t.Fatal("Switch_OpaqueDirection did not call the LEFT case")
		}
	})
}

// TestOneOfSwitch_opaque verifies that Switch_/Map_ functions for opaque API
// oneofs use WhichXxx() and case constants rather than a type switch.
func TestOneOfSwitch_opaque(t *testing.T) {
	t.Parallel()

	t.Run("it calls the function for the populated oneof option", func(t *testing.T) {
		t.Parallel()

		m := &OpaqueRecord{}
		m.SetDecrement(123)

		called := false
		Switch_OpaqueRecord_Operation(
			m,
			func(v int32) { panic("unexpected increment case") },
			func(v int32) { called = true },
			func(v string) { panic("unexpected log case") },
			func(*OpaqueRecord_NamingCollision) { panic("unexpected naming_collision case") },
			func() { panic("unexpected none case") },
		)
		if !called {
			t.Fatal("Switch_OpaqueRecord_Operation did not call the decrement case")
		}
	})

	t.Run("it calls the naming_collision case for a message-typed option", func(t *testing.T) {
		t.Parallel()

		m := &OpaqueRecord{}
		m.SetNamingCollision(&OpaqueRecord_NamingCollision{})

		called := false
		Switch_OpaqueRecord_Operation(
			m,
			func(v int32) { panic("unexpected increment case") },
			func(v int32) { panic("unexpected decrement case") },
			func(v string) { panic("unexpected log case") },
			func(*OpaqueRecord_NamingCollision) { called = true },
			func() { panic("unexpected none case") },
		)
		if !called {
			t.Fatal("Switch_OpaqueRecord_Operation did not call the naming_collision case")
		}
	})

	t.Run("it calls none() when the oneof is unset", func(t *testing.T) {
		t.Parallel()

		m := &OpaqueRecord{}

		called := false
		Switch_OpaqueRecord_Operation(
			m,
			func(v int32) { panic("unexpected increment case") },
			func(v int32) { panic("unexpected decrement case") },
			func(v string) { panic("unexpected log case") },
			func(*OpaqueRecord_NamingCollision) { panic("unexpected naming_collision case") },
			func() { called = true },
		)
		if !called {
			t.Fatal("Switch_OpaqueRecord_Operation did not call none()")
		}
	})

	t.Run("Map_ returns the result of the matching case function", func(t *testing.T) {
		t.Parallel()

		m := &OpaqueRecord{}
		m.SetIncrement(42)

		got := Map_OpaqueRecord_Operation(
			m,
			func(v int32) int { return int(v) * 2 },
			func(v int32) int { panic("unexpected decrement case") },
			func(v string) int { panic("unexpected log case") },
			func(*OpaqueRecord_NamingCollision) int { panic("unexpected naming_collision case") },
			func() int { panic("unexpected none case") },
		)
		if got != 84 {
			t.Fatalf("Map_OpaqueRecord_Operation = %d, want 84", got)
		}
	})
}
