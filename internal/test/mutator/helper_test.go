package mutator_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

type message interface {
	comparable
	proto.Message
}

// testMutator calls a mutator method and verifies that the corresponding
// accessor returns the value that was set.
func testMutator[M message, T comparable](
	t *testing.T,
	mutate func(M, T),
	access func(M) T,
	want T,
) {
	t.Helper()

	testMutatorFunc(
		t,
		mutate,
		access,
		want,
		func(a, b T) bool { return a == b },
	)
}

// testMutator calls a mutator method and verifies that the corresponding
// accessor returns the value that was set.
func testMutatorFunc[M message, T any](
	t *testing.T,
	mutate func(M, T),
	access func(M) T,
	want T,
	eq func(T, T) bool,
) {
	t.Helper()

	if reflect.ValueOf(want).IsZero() {
		panic("cannot test mutator with zero value")
	}

	m := reflect.New(
		reflect.TypeFor[M]().Elem(),
	).Interface().(M)

	mutate(m, want)

	if got := access(m); !eq(got, want) {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}
}
