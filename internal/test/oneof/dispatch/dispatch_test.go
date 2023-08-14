package dispatch_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/oneof/dispatch"
)

func TestOneOfDispatch(t *testing.T) {
	m := &Message{}

	want := &Value{}
	m.SetFieldA2(want)

	called := false

	m.DispatchGroup(
		func(*Value) {
			t.Fatal("called the function for the FieldA1 case, expected FieldA2")
		},
		func(got *Value) {
			if got != want {
				t.Fatal("unexpected value")
			}
			called = true
		},
		func(string) {
			t.Fatal("called the function for the FieldA3 case, expected FieldA2")
		},
		func() {
			t.Fatal("called the function for the default case, expected FieldA2")
		},
	)

	if !called {
		t.Fatal("did not call any function")
	}
}