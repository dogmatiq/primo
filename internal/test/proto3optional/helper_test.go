package proto3optional_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

// testOptionalMutator calls a mutator that takes a *T value and verifies that
// the accessor returns the corresponding T value.
//
// This covers the proto3 optional field pattern, where the underlying Go struct
// field is *T, the setter takes *T, and the protoc-gen-go accessor returns T.
func testOptionalMutator[M proto.Message, T comparable](
	t *testing.T,
	mutate func(M, *T),
	access func(M) T,
	want T,
) {
	t.Helper()

	var m M
	m = reflect.New(
		reflect.TypeOf(m).Elem(),
	).Interface().(M)

	mutate(m, &want)

	if got := access(m); got != want {
		t.Fatalf(
			"mutator did not set the field: got: %v, want: %v",
			got,
			want,
		)
	}
}
