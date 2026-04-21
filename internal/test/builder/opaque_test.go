package builder_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/builder"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestBuilder_opaque(t *testing.T) {
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
