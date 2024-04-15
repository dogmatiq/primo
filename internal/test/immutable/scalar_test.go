package immutable_test

// import (
// 	"bytes"
// 	"testing"

// 	. "github.com/dogmatiq/primo/internal/test/mutator"
// )

// func TestScalar(t *testing.T) {
// 	t.Parallel()

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetInt32,
// 		(ImmutableScalar).GetInt32,
// 		-123,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetInt64,
// 		ImmutableScalar.GetInt64,
// 		-456,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetUint32,
// 		ImmutableScalar.GetUint32,
// 		123,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetUint64,
// 		ImmutableScalar.GetUint64,
// 		456,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetSint32,
// 		ImmutableScalar.GetSint32,
// 		-123,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetSint64,
// 		ImmutableScalar.GetSint64,
// 		-456,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetFixed32,
// 		ImmutableScalar.GetFixed32,
// 		123,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetFixed64,
// 		ImmutableScalar.GetFixed64,
// 		456,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetSfixed32,
// 		ImmutableScalar.GetSfixed32,
// 		-123,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetSfixed64,
// 		ImmutableScalar.GetSfixed64,
// 		-456,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetFloat,
// 		ImmutableScalar.GetFloat,
// 		3.14,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetDouble,
// 		ImmutableScalar.GetDouble,
// 		3.14159,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetBool,
// 		ImmutableScalar.GetBool,
// 		true,
// 	)

// 	testInterfaces[*Scalar](
// 		t,
// 		MutableScalar.SetString_,
// 		ImmutableScalar.GetString_,
// 		"<value>",
// 	)

// 	testInterfacesFunc[*Scalar](
// 		t,
// 		MutableScalar.SetBytes,
// 		ImmutableScalar.GetBytes,
// 		[]byte("<value>"),
// 		bytes.Equal,
// 	)
// }
