package accessor_test

import (
	reflect "reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

// oneOfAccessors is the interface that both open-API and opaque-API generated
// types with the same one-of field layout must implement.
type oneOfAccessors interface {
	proto.Message
	SetFieldA(int32)
	TryGetFieldA() (int32, bool)
	SetFieldB(int32)
	TryGetFieldB() (int32, bool)
	SetFieldC(string)
	TryGetFieldC() (string, bool)
}

// testOneOfAccessorSuite runs the full accessor test suite against any type
// that implements oneOfAccessors.
func testOneOfAccessorSuite[M oneOfAccessors](t *testing.T) {
	t.Helper()

	t.Run("when the one-of value is set", func(t *testing.T) {
		t.Run("it returns the value with ok set to true", func(t *testing.T) {
			testAccessor(t, func(m M, v int32) { m.SetFieldA(v) }, func(m M) (int32, bool) { return m.TryGetFieldA() }, 123)
			testAccessor(t, func(m M, v int32) { m.SetFieldA(v) }, func(m M) (int32, bool) { return m.TryGetFieldA() }, 0)
			testAccessor(t, func(m M, v int32) { m.SetFieldB(v) }, func(m M) (int32, bool) { return m.TryGetFieldB() }, 456)
			testAccessor(t, func(m M, v int32) { m.SetFieldB(v) }, func(m M) (int32, bool) { return m.TryGetFieldB() }, 0)
			testAccessor(t, func(m M, v string) { m.SetFieldC(v) }, func(m M) (string, bool) { return m.TryGetFieldC() }, "<value>")
			testAccessor(t, func(m M, v string) { m.SetFieldC(v) }, func(m M) (string, bool) { return m.TryGetFieldC() }, "")
		})
	})

	t.Run("when the one-of value is set to some other option", func(t *testing.T) {
		t.Run("it returns with ok set to false", func(t *testing.T) {
			var zero M
			m := reflect.New(reflect.TypeOf(zero).Elem()).Interface().(M)
			m.SetFieldB(456)

			if _, ok := m.TryGetFieldA(); ok {
				t.Fatal("TryGetFieldA() returned ok as true, want false")
			}

			if _, ok := m.TryGetFieldC(); ok {
				t.Fatal("TryGetFieldC() returned ok as true, want false")
			}
		})
	})
}

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
