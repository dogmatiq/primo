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
	mutator func(M, T) M,
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
func testMutatorFunc[M message, T any](
	t *testing.T,
	mutator func(M, T) M,
	accessor func(M) T,
	want T,
	eq func(T, T) bool,
) {
	t.Helper()

	if reflect.ValueOf(want).IsZero() {
		panic("cannot test mutator with zero value")
	}

	var nil_ M

	t.Log("calling mutator on a nil message")
	m := mutator(nil_, want)
	if m == nil_ {
		t.Fatal("mutator returned a nil message")
	}

	t.Log("checking that the mutator set the field")
	if got := accessor(m); !eq(got, want) {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}

	t.Log("calling mutator on a non-nil message")
	var zero T
	want = zero
	r := mutator(m, want)
	if r != m {
		t.Fatal("mutator did not return the mutated message")
	}

	if got := accessor(m); !eq(got, want) {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}
}
