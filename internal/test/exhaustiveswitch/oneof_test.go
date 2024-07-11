package exhaustiveswitch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/exhaustiveswitch"
)

func TestOneOfSwitch(t *testing.T) {
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

			called := false

			Switch_Record_Operation(
				rec,
				func(v int32) { panic("unexpected increment operation") },
				func(v int32) { called = true },
				func(m string) { panic("unexpected log operation") },
				func(*Record_NamingCollision) { panic("unexpected NamingCollision operation") },
				func() { panic("unexpected default case") },
			)

			if !called {
				t.Fatalf("expected case function to be called")
			}
		},
	)

	t.Run(
		"it calls the function associated with the default case if the one-of field is nil",
		func(t *testing.T) {
			t.Parallel()

			cases := []struct {
				name   string
				record *Record
			}{
				{
					"nil message",
					nil,
				},
				{
					"nil field",
					&Record{},
				},
			}

			for _, c := range cases {
				t.Run(
					c.name,
					func(t *testing.T) {
						t.Parallel()

						called := false

						Switch_Record_Operation(
							c.record,
							func(int32) { panic("unexpected increment operation") },
							func(int32) { panic("unexpected decrement operation") },
							func(string) { panic("unexpected log operation") },
							func(*Record_NamingCollision) { panic("unexpected NamingCollision operation") },
							func() { called = true },
						)

						if !called {
							t.Fatalf("expected case function to be called")
						}
					})
			}
		},
	)
}

func TestOneOfMap(t *testing.T) {
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
			got := Map_Record_Operation(
				rec,
				func(v int32) int32 { panic("unexpected increment operation") },
				func(v int32) int32 { return v - 1 },
				func(m string) int32 { panic("unexpected log operation") },
				func(*Record_NamingCollision) int32 { panic("unexpected NamingCollision operation") },
				func() int32 { panic("unexpected default case") },
			)

			if got != want {
				t.Fatalf("unexpected result: got %q, want %q", got, want)
			}
		},
	)

	t.Run(
		"it calls the function associated with the default case if the one-of field is nil",
		func(t *testing.T) {
			t.Parallel()

			cases := []struct {
				name   string
				record *Record
			}{
				{
					"nil message",
					nil,
				},
				{
					"nil field",
					&Record{},
				},
			}

			for _, c := range cases {
				t.Run(
					c.name,
					func(t *testing.T) {
						t.Parallel()

						want := int32(123)
						got := Map_Record_Operation(
							c.record,
							func(int32) int32 { panic("unexpected increment operation") },
							func(int32) int32 { panic("unexpected decrement operation") },
							func(string) int32 { panic("unexpected log operation") },
							func(*Record_NamingCollision) int32 { panic("unexpected NamingCollision operation") },
							func() int32 { return 123 },
						)

						if got != want {
							t.Fatalf("unexpected result: got %q, want %q", got, want)
						}
					})
			}
		},
	)
}
