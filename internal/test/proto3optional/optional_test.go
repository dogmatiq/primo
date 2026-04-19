package proto3optional_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/proto3optional"
)

// TestOptionalMutators verifies that proto3 optional fields generate plain
// Set/Get accessors — not oneof-style TryGet — and that the setter correctly
// takes a pointer to the value.
func TestOptionalMutators(t *testing.T) {
	t.Parallel()

	testOptionalMutator(t, (*Message).SetNum, (*Message).GetNum, int32(123))
	testOptionalMutator(t, (*Message).SetStr, (*Message).GetStr, "<value>")
	testOptionalMutator(t, (*Message).SetFlag, (*Message).GetFlag, true)
	testOptionalMutator(t, (*Message).SetStatus, (*Message).GetStatus, Status_STATUS_OK)

	t.Run("message field", func(t *testing.T) {
		child := &Message_Child{}
		m := &Message{}
		m.SetChild(child)
		if got := m.GetChild(); got != child {
			t.Fatal("SetChild did not set the field")
		}
	})
}

// TestRealOneofAccessors verifies that a real oneof declared alongside optional
// fields still generates TryGet accessor methods.
func TestRealOneofAccessors(t *testing.T) {
	t.Parallel()

	t.Run("when the oneof value is set", func(t *testing.T) {
		t.Run("it returns the value with ok set to true", func(t *testing.T) {
			m := &Message{}
			m.SetChoiceA("<value>")
			if v, ok := m.TryGetChoiceA(); !ok || v != "<value>" {
				t.Fatalf("TryGetChoiceA() = (%q, %v), want (%q, true)", v, ok, "<value>")
			}

			m = &Message{}
			m.SetChoiceB(123)
			if v, ok := m.TryGetChoiceB(); !ok || v != 123 {
				t.Fatalf("TryGetChoiceB() = (%v, %v), want (123, true)", v, ok)
			}
		})
	})

	t.Run("when a different oneof option is set", func(t *testing.T) {
		t.Run("it returns with ok set to false", func(t *testing.T) {
			m := &Message{}
			m.SetChoiceB(123)
			if _, ok := m.TryGetChoiceA(); ok {
				t.Fatal("TryGetChoiceA() returned ok as true, want false")
			}
		})
	})
}

// TestRealOneofSwitch verifies that a real oneof declared alongside optional
// fields still generates Switch_ and Map_ functions.
func TestRealOneofSwitch(t *testing.T) {
	t.Parallel()

	t.Run("it calls the function for the populated oneof option", func(t *testing.T) {
		m := &Message{}
		m.SetChoiceA("<value>")

		called := false
		Switch_Message_Choice(
			m,
			func(v string) { called = true },
			func(v int32) { panic("unexpected ChoiceB case") },
			func() { panic("unexpected none case") },
		)
		if !called {
			t.Fatal("Switch_Message_Choice did not call the ChoiceA case")
		}
	})

	t.Run("it calls none() when choice is unset", func(t *testing.T) {
		m := &Message{}

		called := false
		Switch_Message_Choice(
			m,
			func(v string) { panic("unexpected ChoiceA case") },
			func(v int32) { panic("unexpected ChoiceB case") },
			func() { called = true },
		)
		if !called {
			t.Fatal("Switch_Message_Choice did not call none()")
		}
	})
}

// TestBuilder verifies that the builder correctly handles optional fields,
// which use pointer-typed WithXxx methods.
func TestBuilder(t *testing.T) {
	t.Parallel()

	num := int32(123)
	str := "<value>"
	flag := true
	status := Status_STATUS_OK

	m := NewMessageBuilder().
		WithNum(&num).
		WithStr(&str).
		WithFlag(&flag).
		WithStatus(&status).
		Build()

	if got := m.GetNum(); got != num {
		t.Fatalf("Num: got %v, want %v", got, num)
	}
	if got := m.GetStr(); got != str {
		t.Fatalf("Str: got %q, want %q", got, str)
	}
	if got := m.GetFlag(); got != flag {
		t.Fatalf("Flag: got %v, want %v", got, flag)
	}
	if got := m.GetStatus(); got != status {
		t.Fatalf("Status: got %v, want %v", got, status)
	}
}
