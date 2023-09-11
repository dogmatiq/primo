package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

func TestEnum(t *testing.T) {
	t.Parallel()

	t.Run(
		"it calls the function associated with the enum value",
		func(t *testing.T) {
			t.Parallel()

			want := "<green>"
			got := Switch_Color(
				Color_GREEN,
				func(Color_UNKNOWN_Case) string {
					panic("dispatched to unexpected function: UNKNOWN")
				},
				func(Color_RED_Case) string {
					panic("dispatched to unexpected function: RED")
				},
				func(Color_GREEN_Case) string {
					return "<green>"
				},
				func(Color_BLUE_Case) string {
					panic("dispatched to unexpected function: BLUE")
				},
			)

			if got != want {
				t.Fatalf("unexpected result: got %q, want %q", got, want)
			}
		},
	)
}
