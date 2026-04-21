package builder_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

// builderOf is the interface that every generated builder type satisfies.
type builderOf[M proto.Message, B any] interface {
	WithFieldA(int32) B
	WithFieldB(string) B
	From(M) B
	Build() M
}

// newEmpty returns a new zero-value instance of M.
func newEmpty[M proto.Message]() M {
	var zero M
	return reflect.New(reflect.TypeOf(zero).Elem()).Interface().(M)
}

// testBuilderSuite runs the tests that apply to every generated builder.
func testBuilderSuite[M proto.Message, B builderOf[M, B]](t *testing.T, newBuilder func() B) {
	t.Helper()

	t.Run("func Build()", func(t *testing.T) {
		t.Run("it returns a new message each call", func(t *testing.T) {
			t.Parallel()
			b := newBuilder()
			first, second := b.Build(), b.Build()
			if any(first) == any(second) {
				t.Fatal("Build() returned the same pointer on consecutive calls")
			}
		})
	})

	t.Run("func From()", func(t *testing.T) {
		t.Run("it clears stale values from the previous prototype", func(t *testing.T) {
			t.Parallel()
			b := newBuilder()
			b.WithFieldA(123)
			b.WithFieldB("abc")
			b.From(newEmpty[M]())
			if diff := cmp.Diff(
				newEmpty[M](),
				b.Build(),
				protocmp.Transform(),
			); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	})
}
