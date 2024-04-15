package immutable_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

type message interface {
	comparable
	proto.Message
}

// testInterfaces calls a mutator method and verifies that the corresponding
// accessors return the value that was set.
func testInterfaces[M message, S, G any, T comparable](
	t *testing.T,
	mutate func(S, T),
	access func(G) T,
	want T,
) {
	t.Helper()

	testInterfacesFunc[M](
		t,
		mutate,
		access,
		want,
		func(a, b T) bool { return a == b },
	)
}

// testMutator calls a mutator method and verifies that the corresponding
// accessors return the value that was set.
func testInterfacesFunc[M message, S, G any, T any](
	t *testing.T,
	mutate func(S, T),
	access func(G) T,
	want T,
	eq func(T, T) bool,
) {
	t.Helper()

	if reflect.ValueOf(want).IsZero() {
		panic("cannot test interfaces with zero value")
	}

	m := reflect.New(
		reflect.TypeFor[M]().Elem(),
	).Interface()

	mutate(m.(S), want)

	if got := access(m.(G)); !eq(got, want) {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}
}
