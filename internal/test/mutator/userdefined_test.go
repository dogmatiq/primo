package mutator_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutator"
)

func TestUserDefined(t *testing.T) {
	t.Parallel()

	testMutator(
		t,
		(*UserDefined).SetMessage,
		(*UserDefined).GetMessage,
		&UserDefinedMessage{},
	)

	testMutator(
		t,
		(*UserDefined).SetEnum,
		(*UserDefined).GetEnum,
		UserDefinedEnum_VALUE,
	)

	testMutator(
		t,
		(*UserDefined).SetNested,
		(*UserDefined).GetNested,
		&UserDefined_Nested{},
	)
}
