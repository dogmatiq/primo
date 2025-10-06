package marshaling_test

import (
	"encoding"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/marshaling"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var (
	_ encoding.BinaryMarshaler   = (*Message)(nil)
	_ encoding.BinaryUnmarshaler = (*Message)(nil)
)

func TestMarshaling(t *testing.T) {
	t.Parallel()

	t.Run("MarshalBinary()", func(t *testing.T) {
		want := &Message{
			Value: "<value>",
		}

		data, err := want.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}

		got := &Message{}
		if err := proto.Unmarshal(data, got); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(
			want,
			got,
			protocmp.Transform(),
		); diff != "" {
			t.Fatalf("unexpected result (+got, -want):\n%s", diff)
		}
	})

	t.Run("UnmarshalBinary()", func(t *testing.T) {
		want := &Message{
			Value: "<value>",
		}

		data, err := proto.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}

		got := &Message{}
		if err := got.UnmarshalBinary(data); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(
			want,
			got,
			protocmp.Transform(),
		); diff != "" {
			t.Fatalf("unexpected result (+got, -want):\n%s", diff)
		}
	})
}
