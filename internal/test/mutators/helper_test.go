package mutators_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

// testMutator calls a mutator method and verifies that the corresponding
// accessor returns the value that was set.
func testMutator[M proto.Message, T comparable](
	t *testing.T,
	mutator func(M, T),
	accessor func(M) T,
	want T,
) {
	t.Helper()

	testMutatorFunc(
		t,
		mutator,
		accessor,
		want,
		func(a, b T) bool { return a == b },
	)
}

// testMutator calls a mutator method and verifies that the corresponding
// accessor returns the value that was set.
func testMutatorFunc[M proto.Message, T any](
	t *testing.T,
	mutator func(M, T),
	accessor func(M) T,
	want T,
	eq func(T, T) bool,
) {
	t.Helper()

	if reflect.ValueOf(want).IsZero() {
		panic("cannot test mutator with zero value")
	}

	var m M
	m = reflect.New(
		reflect.TypeOf(m).Elem(),
	).Interface().(M)

	mutator(m, want)

	got := accessor(m)

	if !eq(got, want) {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}
}
