package builder_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/builder"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

// TestBuilder_proto3optional verifies that proto3 optional fields use
// pointer-typed WithXxx methods in the generated builder.
func TestBuilder_proto3optional(t *testing.T) {
	t.Parallel()

	t.Run("func Build()", func(t *testing.T) {
		t.Run("it returns a new message each call", func(t *testing.T) {
			t.Parallel()
			b := NewProto3OptionalMessageBuilder()
			first, second := b.Build(), b.Build()
			if first == second {
				t.Fatal("Build() returned the same pointer on consecutive calls")
			}
		})
	})

	t.Run("func From()", func(t *testing.T) {
		t.Run("it clears stale values from the previous prototype", func(t *testing.T) {
			t.Parallel()
			num := int32(123)
			str := "<value>"
			b := NewProto3OptionalMessageBuilder()
			b.WithNum(&num)
			b.WithStr(&str)
			b.From(&Proto3OptionalMessage{})
			if diff := cmp.Diff(
				&Proto3OptionalMessage{},
				b.Build(),
				protocmp.Transform(),
			); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	})

	t.Run("it sets optional fields via pointer-typed With methods", func(t *testing.T) {
		t.Parallel()

		num := int32(123)
		str := "<value>"
		flag := true
		status := Proto3OptionalStatus_PROTO3_OPTIONAL_STATUS_OK

		m := NewProto3OptionalMessageBuilder().
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
	})
}
