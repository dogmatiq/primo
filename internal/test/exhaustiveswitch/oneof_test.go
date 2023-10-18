package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

func TestOneOf(t *testing.T) {
	t.Parallel()

	t.Run(
		"it calls the function associated with the populated one-of option",
		func(t *testing.T) {
			t.Parallel()

			rec := &Record{
				Operation: &Record_Decrement{
					Decrement: 123,
				},
			}

			want := int32(122)
			got := Switch_Record_Operation(
				rec,
				func(v int32) int32 { panic("unexpected increment operation") },
				func(v int32) int32 { return v - 1 },
				func(m string) int32 { panic("unexpected log operation") },
				func(*Record_NamingCollision) int32 { panic("unexpected NamingCollision operation") },
			)

			if got != want {
				t.Fatalf("unexpected result: got %q, want %q", got, want)
			}
		},
	)

	t.Run(
		"it panics if the field is nil",
		func(t *testing.T) {
			t.Parallel()

			rec := &Record{}

			defer func() {
				got := recover()
				want := "Switch_Record_Operation: x.Operation is nil"

				if got != want {
					t.Fatalf(
						"unexpected panic message: got %q, want %q",
						got,
						want,
					)
				}
			}()

			Switch_Record_Operation(
				rec,
				func(int32) error { panic("unexpected increment operation") },
				func(int32) error { panic("unexpected decrement operation") },
				func(string) error { panic("unexpected log operation") },
				func(*Record_NamingCollision) error { panic("unexpected NamingCollision operation") },
			)
		},
	)
}
