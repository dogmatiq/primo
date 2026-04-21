package mutator_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutator"
)

// TestMutator_proto3optional verifies that proto3 optional fields generate
// plain Set/Get accessors — not oneof-style TryGet — and that the setter
// correctly takes a pointer to the value.
func TestMutator_proto3optional(t *testing.T) {
	t.Parallel()

	testOptionalMutator(t, (*Proto3OptionalMessage).SetNum, (*Proto3OptionalMessage).GetNum, int32(123))
	testOptionalMutator(t, (*Proto3OptionalMessage).SetStr, (*Proto3OptionalMessage).GetStr, "<value>")
	testOptionalMutator(t, (*Proto3OptionalMessage).SetFlag, (*Proto3OptionalMessage).GetFlag, true)
	testOptionalMutator(t, (*Proto3OptionalMessage).SetStatus, (*Proto3OptionalMessage).GetStatus, Proto3OptionalStatus_PROTO3_OPTIONAL_STATUS_OK)

	t.Run("message field", func(t *testing.T) {
		child := &Proto3OptionalMessage_Proto3OptionalChild{}
		m := &Proto3OptionalMessage{}
		m.SetChild(child)
		if got := m.GetChild(); got != child {
			t.Fatal("SetChild did not set the field")
		}
	})
}
