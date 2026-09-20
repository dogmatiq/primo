package formatting_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/formatting"
)

var (
	_ fmt.Formatter = (*Stringable)(nil)
	_ fmt.Formatter = (*Plain)(nil)
)

func TestFormat(t *testing.T) {
	t.Parallel()

	t.Run("uses AsString() for the %s and %q verbs", func(t *testing.T) {
		x := &Stringable{Value: "hello"}

		if got, want := fmt.Sprintf("%s", x), x.AsString(); got != want {
			t.Errorf("%%s: got %q, want %q", got, want)
		}
		if got, want := fmt.Sprintf("%q", x), strconv.Quote(x.AsString()); got != want {
			t.Errorf("%%q: got %q, want %q", got, want)
		}
	})

	t.Run("uses String() for other verbs even when AsString() is implemented", func(t *testing.T) {
		x := &Stringable{Value: "hello"}

		if got, want := fmt.Sprintf("%v", x), x.String(); got != want {
			t.Errorf("%%v: got %q, want %q", got, want)
		}
	})

	t.Run("delegates the Stringer verbs to String() when AsString() is absent", func(t *testing.T) {
		x := &Plain{Value: "hello"}

		for _, verb := range []string{"%v", "%+v", "%s", "%x", "%X", "%q"} {
			if got, want := fmt.Sprintf(verb, x), fmt.Sprintf(verb, x.String()); got != want {
				t.Errorf("%s: got %q, want %q", verb, got, want)
			}
		}
	})

	t.Run("falls back to default formatting for the other verbs", func(t *testing.T) {
		x := &Plain{Value: "hello"}

		// A method-less type with the same structure produces the output that
		// fmt would produce if Format() were absent.
		type realType = Plain
		type Plain realType

		// The format string is held in a variable so that it is not analyzed by
		// go vet, which rejects the %d verb applied to a struct.
		verb := "%d"
		if got, want := fmt.Sprintf(verb, x), fmt.Sprintf(verb, (*Plain)(x)); got != want {
			t.Errorf("%s: got %q, want %q", verb, got, want)
		}
	})

	t.Run("preserves the type name for the %#v verb", func(t *testing.T) {
		x := &Plain{Value: "hello"}

		got := fmt.Sprintf("%#v", x)
		if !strings.Contains(got, "formatting.Plain{") || !strings.Contains(got, `Value:"hello"`) {
			t.Errorf("%%#v: got %q, want it to contain the qualified type name and field", got)
		}
	})
}
