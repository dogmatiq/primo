package constructor_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/constructor"
	"google.golang.org/protobuf/proto"
)

func TestConstructor(t *testing.T) {
	t.Parallel()

	got := NewMessage()
	want := &Message{}

	if !proto.Equal(got, want) {
		t.Fatalf(
			"constructor did not return the expected message: got: %v, want: %v",
			got,
			want,
		)
	}
}
