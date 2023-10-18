package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

func TestEnumSwitch(t *testing.T) {
	t.Parallel()

	t.Run(
		"it calls the function associated with the enum value",
		func(t *testing.T) {
			t.Parallel()

			called := false

			Switch_Direction(
				Direction_LEFT,
				func(Direction_UNKNOWN_Case) { panic("unexpected UNKNOWN direction") },
				func(Direction_SINISTER_Case) { called = true },
				func(Direction_DEXTER_Case) { panic("unexpected LEFT direction") },
			)

			if !called {
				t.Fatalf("expected case function to be called")
			}
		},
	)
}
func TestEnumMap(t *testing.T) {
	t.Parallel()

	t.Run(
		"it calls the function associated with the enum value",
		func(t *testing.T) {
			t.Parallel()

			want := "<value>"
			got := Map_Direction(
				Direction_LEFT,
				func(Direction_UNKNOWN_Case) string { panic("unexpected UNKNOWN direction") },
				func(Direction_SINISTER_Case) string { return "<value>" },
				func(Direction_DEXTER_Case) string { panic("unexpected LEFT direction") },
			)

			if got != want {
				t.Fatalf("unexpected result: got %q, want %q", got, want)
			}
		},
	)
}
