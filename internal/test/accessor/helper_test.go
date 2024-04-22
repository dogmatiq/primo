package accessor_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

// testAccessor calls a mutator method to set the value and verifies that the
// corresponding accessor returns the expected value and an ok value of true.
func testAccessor[M proto.Message, T comparable](
	t *testing.T,
	mutator func(M, T),
	accessor func(M) (T, bool),
	want T,
) {
	t.Helper()

	testAccessorFunc(
		t,
		mutator,
		accessor,
		want,
		func(a, b T) bool { return a == b },
	)
}

// testAccessorFunc calls a mutator method to set the value and verifies that the
// corresponding accessor returns the expected value and an ok value of true.
func testAccessorFunc[M proto.Message, T any](
	t *testing.T,
	mutate func(M, T),
	access func(M) (T, bool),
	want T,
	eq func(T, T) bool,
) {
	t.Helper()

	var m M
	m = reflect.New(
		reflect.TypeOf(m).Elem(),
	).Interface().(M)

	mutate(m, want)

	got, gotOK := access(m)
	if !gotOK {
		t.Fatal("accessor did not return with ok set to true")
	}

	if !eq(got, want) {
		t.Fatalf(
			"accessor did not return the expected value: got: %v, want: %v",
			got,
			want,
		)
	}
}
