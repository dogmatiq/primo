package builder_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/builder"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestBuilder_opaque(t *testing.T) {
	t.Parallel()

	t.Run("func Build()", func(t *testing.T) {
		t.Run("it returns a new message each call", func(t *testing.T) {
			t.Parallel()
			b := NewOpaqueMessageBuilder()
			first, second := b.Build(), b.Build()
			if first == second {
				t.Fatal("Build() returned the same pointer on consecutive calls")
			}
		})
	})

	t.Run("func From()", func(t *testing.T) {
		t.Run("it clears stale values from the previous prototype", func(t *testing.T) {
			t.Parallel()
			b := NewOpaqueMessageBuilder()
			b.WithFieldA(123)
			b.WithFieldB("abc")
			b.From(&OpaqueMessage{})
			if diff := cmp.Diff(
				&OpaqueMessage{},
				b.Build(),
				protocmp.Transform(),
			); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	})

	t.Run("when built from a prototype with field overrides", func(t *testing.T) {
		t.Run("it returns the expected message on every Build call", func(t *testing.T) {
			t.Parallel()

			prototype := &OpaqueMessage{}
			prototype.SetFieldA(789)
			prototype.SetFieldB("abc")

			builder := NewOpaqueMessageBuilder().
				From(prototype).
				WithFieldA(123).
				WithFieldC(456).
				WithFieldE(101).
				WithFieldF([]int32{1, 2, 3}).
				WithFieldG(map[string]int32{"a": 1}).
				WithNested(
					NewOpaqueMessage_NestedBuilder().
						WithField(789).
						Build(),
				)

			want := &OpaqueMessage{}
			want.SetFieldA(123)
			want.SetFieldB("abc")
			want.SetFieldC(456)
			want.SetFieldE(101)
			want.SetFieldF([]int32{1, 2, 3})
			want.SetFieldG(map[string]int32{"a": 1})
			nested := &OpaqueMessage_Nested{}
			nested.SetField(789)
			want.SetNested(nested)

			for i := range 2 {
				if diff := cmp.Diff(
					want,
					builder.Build(),
					protocmp.Transform(),
				); diff != "" {
					t.Fatalf("Build() call %d: unexpected result (-want +got):\n%s", i+1, diff)
				}
			}
		})
	})
}
