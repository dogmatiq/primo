package mutators_test

import (
	"maps"
	"slices"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutators"
)

func TestCollection(t *testing.T) {
	testMutatorFunc(
		t,
		(*Collection).SetRepeated,
		(*Collection).GetRepeated,
		[]int32{1, 2, 3},
		func(a, b []int32) bool {
			return slices.Equal(a, b)
		},
	)

	testMutatorFunc(
		t,
		(*Collection).SetMap,
		(*Collection).GetMap,
		map[string]int32{
			"one": 1,
			"two": 2,
		},
		func(a, b map[string]int32) bool {
			return maps.Equal(a, b)
		},
	)
}
