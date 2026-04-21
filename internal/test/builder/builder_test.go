package builder_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/builder"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestBuilder(t *testing.T) {
	t.Parallel()

	t.Run("func Build()", func(t *testing.T) {
		t.Run("it returns a new message each call", func(t *testing.T) {
			t.Parallel()
			b := NewMessageBuilder()
			first, second := b.Build(), b.Build()
			if first == second {
				t.Fatal("Build() returned the same pointer on consecutive calls")
			}
		})
	})

	t.Run("func From()", func(t *testing.T) {
		t.Run("it clears stale values from the previous prototype", func(t *testing.T) {
			t.Parallel()
			b := NewMessageBuilder()
			b.WithFieldA(123)
			b.WithFieldB("abc")
			b.From(&Message{})
			if diff := cmp.Diff(
				&Message{},
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

			prototype := &Message{
				FieldA: 789,
				FieldB: "abc",
			}

			builder := NewMessageBuilder().
				From(prototype).
				WithFieldA(123).
				WithFieldC(456).
				WithFieldE([]int32{1, 2, 3}).
				WithFieldF(map[string]int32{"a": 1}).
				WithNested(
					NewMessage_NestedBuilder().
						WithField(789).
						Build(),
				)

			want := &Message{
				FieldA: 123,
				FieldB: "abc",
				Group: &Message_FieldC{
					FieldC: 456,
				},
				FieldE: []int32{1, 2, 3},
				FieldF: map[string]int32{"a": 1},
				Nested: &Message_Nested{
					Field: 789,
				},
			}

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
