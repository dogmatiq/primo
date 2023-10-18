package builder_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/builder"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestBuilder(t *testing.T) {
	t.Parallel()

	prototype := &Message{
		FieldA: 789,
		FieldB: "abc",
	}

	builder := NewMessageBuilder().
		From(prototype).
		WithFieldA(123).
		WithFieldC(456)

	want := &Message{
		FieldA: 123,
		FieldB: "abc",
		Group: &Message_FieldC{
			FieldC: 456,
		},
	}

	gotFirst := builder.Build()

	if diff := cmp.Diff(
		want,
		gotFirst,
		protocmp.Transform(),
	); diff != "" {
		t.Fatalf("unexpected result (+got, -want):\n%s", diff)
	}

	gotSecond := builder.Build()

	if gotSecond == gotFirst {
		t.Fatal("Build() returned the same message multiple times")
	}

	if diff := cmp.Diff(
		want,
		gotSecond,
		protocmp.Transform(),
	); diff != "" {
		t.Fatalf("unexpected result (+got, -want):\n%s", diff)
	}
}
