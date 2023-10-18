package mutator_test

import (
	"bytes"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutator"
)

func TestScalar(t *testing.T) {
	t.Parallel()

	testMutator(
		t,
		(*Scalar).SetInt32,
		(*Scalar).GetInt32,
		-123,
	)

	testMutator(
		t,
		(*Scalar).SetInt64,
		(*Scalar).GetInt64,
		-456,
	)

	testMutator(
		t,
		(*Scalar).SetUint32,
		(*Scalar).GetUint32,
		123,
	)

	testMutator(
		t,
		(*Scalar).SetUint64,
		(*Scalar).GetUint64,
		456,
	)

	testMutator(
		t,
		(*Scalar).SetSint32,
		(*Scalar).GetSint32,
		-123,
	)

	testMutator(
		t,
		(*Scalar).SetSint64,
		(*Scalar).GetSint64,
		-456,
	)

	testMutator(
		t,
		(*Scalar).SetFixed32,
		(*Scalar).GetFixed32,
		123,
	)

	testMutator(
		t,
		(*Scalar).SetFixed64,
		(*Scalar).GetFixed64,
		456,
	)

	testMutator(
		t,
		(*Scalar).SetSfixed32,
		(*Scalar).GetSfixed32,
		-123,
	)

	testMutator(
		t,
		(*Scalar).SetSfixed64,
		(*Scalar).GetSfixed64,
		-456,
	)

	testMutator(
		t,
		(*Scalar).SetFloat,
		(*Scalar).GetFloat,
		3.14,
	)

	testMutator(
		t,
		(*Scalar).SetDouble,
		(*Scalar).GetDouble,
		3.14159,
	)

	testMutator(
		t,
		(*Scalar).SetBool,
		(*Scalar).GetBool,
		true,
	)

	testMutator(
		t,
		(*Scalar).SetString_,
		(*Scalar).GetString_,
		"<value>",
	)

	testMutatorFunc(
		t,
		(*Scalar).SetBytes,
		(*Scalar).GetBytes,
		[]byte("<value>"),
		bytes.Equal,
	)
}
