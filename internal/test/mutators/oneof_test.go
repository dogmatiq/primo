package mutators_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutators"
)

func TestOneOf(t *testing.T) {
	t.Parallel()

	testMutator(
		t,
		(*OneOf).SetFieldA,
		(*OneOf).GetFieldA,
		123,
	)

	testMutator(
		t,
		(*OneOf).SetFieldB,
		(*OneOf).GetFieldB,
		456,
	)

	testMutator(
		t,
		(*OneOf).SetFieldC,
		(*OneOf).GetFieldC,
		"<value>",
	)
}
