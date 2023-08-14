package mutators_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutators"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func TestMutators_oneOf(t *testing.T) {
	m := &Message{}

	{
		v := &Value{}
		m.SetFieldA1(v)
		if m.GetFieldA1() != v {
			t.Fatal("SetFieldA1() did not set the field")
		}
	}

	{
		v := &Value{}
		m.SetFieldA2(v)
		if m.GetFieldA2() != v {
			t.Fatal("SetFieldA2() did not set the field")
		}
	}

	{
		v := "<string>"
		m.SetFieldA3(v)
		if m.GetFieldA3() != v {
			t.Fatal("SetFieldA3() did not set the field")
		}
	}

	{
		v := &timestamppb.Timestamp{}
		m.SetFieldB1(v)
		if m.GetFieldB1() != v {
			t.Fatal("SetFieldB1() did not set the field")
		}
	}

	{
		v := int32(123)
		m.SetFieldB2(v)
		if m.GetFieldB2() != v {
			t.Fatal("SetFieldB2() did not set the field")
		}
	}

	{
		v := Enumeration(123)
		m.SetFieldB3(v)
		if m.GetFieldB3() != v {
			t.Fatal("SetFieldB3() did not set the field")
		}
	}
}
