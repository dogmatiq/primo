package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

// TestOneOfSwitch_proto3optional verifies that a real oneof declared alongside
// proto3 optional fields still generates Switch_/Map_ functions.
func TestOneOfSwitch_proto3optional(t *testing.T) {
	t.Parallel()

	t.Run("it calls the function for the populated one-of option", func(t *testing.T) {
		t.Parallel()
		m := &Proto3OptionalMessage{}
		m.SetChoiceA("<value>")

		called := false
		Switch_Proto3OptionalMessage_Choice(
			m,
			func(v string) { called = true },
			func(v int32) { panic("unexpected ChoiceB case") },
			func() { panic("unexpected none case") },
		)
		if !called {
			t.Fatal("Switch_Proto3OptionalMessage_Choice did not call the ChoiceA case")
		}
	})

	t.Run("it calls the 'none' function if the one-of field is unset", func(t *testing.T) {
		t.Parallel()
		m := &Proto3OptionalMessage{}

		called := false
		Switch_Proto3OptionalMessage_Choice(
			m,
			func(v string) { panic("unexpected ChoiceA case") },
			func(v int32) { panic("unexpected ChoiceB case") },
			func() { called = true },
		)
		if !called {
			t.Fatal("Switch_Proto3OptionalMessage_Choice did not call none()")
		}
	})
}
