package mutators_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/test/mutators"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/apipb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	sourcecontextpb "google.golang.org/protobuf/types/known/sourcecontextpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	typepb "google.golang.org/protobuf/types/known/typepb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestWellKnown(t *testing.T) {
	testMutator(
		t,
		(*WellKnown).SetAny,
		(*WellKnown).GetAny,
		&anypb.Any{},
	)

	testMutator(
		t,
		(*WellKnown).SetApi,
		(*WellKnown).GetApi,
		&apipb.Api{},
	)

	testMutator(
		t,
		(*WellKnown).SetBoolValue,
		(*WellKnown).GetBoolValue,
		&wrapperspb.BoolValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetBytesValue,
		(*WellKnown).GetBytesValue,
		&wrapperspb.BytesValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetDoubleValue,
		(*WellKnown).GetDoubleValue,
		&wrapperspb.DoubleValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetDuration,
		(*WellKnown).GetDuration,
		&durationpb.Duration{},
	)

	testMutator(
		t,
		(*WellKnown).SetEmpty,
		(*WellKnown).GetEmpty,
		&emptypb.Empty{},
	)

	testMutator(
		t,
		(*WellKnown).SetEnum,
		(*WellKnown).GetEnum,
		&typepb.Enum{},
	)

	testMutator(
		t,
		(*WellKnown).SetEnumValue,
		(*WellKnown).GetEnumValue,
		&typepb.EnumValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetField,
		(*WellKnown).GetField,
		&typepb.Field{},
	)

	testMutator(
		t,
		(*WellKnown).SetFieldCardinality,
		(*WellKnown).GetFieldCardinality,
		typepb.Field_CARDINALITY_REPEATED,
	)

	testMutator(
		t,
		(*WellKnown).SetFieldKind,
		(*WellKnown).GetFieldKind,
		typepb.Field_TYPE_DOUBLE,
	)

	testMutator(
		t,
		(*WellKnown).SetFieldMask,
		(*WellKnown).GetFieldMask,
		&fieldmaskpb.FieldMask{},
	)

	testMutator(
		t,
		(*WellKnown).SetFloatValue,
		(*WellKnown).GetFloatValue,
		&wrapperspb.FloatValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetInt32Value,
		(*WellKnown).GetInt32Value,
		&wrapperspb.Int32Value{},
	)

	testMutator(
		t,
		(*WellKnown).SetInt64Value,
		(*WellKnown).GetInt64Value,
		&wrapperspb.Int64Value{},
	)

	testMutator(
		t,
		(*WellKnown).SetListValue,
		(*WellKnown).GetListValue,
		&structpb.ListValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetMethod,
		(*WellKnown).GetMethod,
		&apipb.Method{},
	)

	testMutator(
		t,
		(*WellKnown).SetMixin,
		(*WellKnown).GetMixin,
		&apipb.Mixin{},
	)

	testMutator(
		t,
		(*WellKnown).SetNullValue,
		(*WellKnown).GetNullValue,
		// This is not a valid NullValue, but we need to test with something
		// that is not the Go zero-value.
		structpb.NullValue(1),
	)

	testMutator(
		t,
		(*WellKnown).SetOption,
		(*WellKnown).GetOption,
		&typepb.Option{},
	)

	testMutator(
		t,
		(*WellKnown).SetSourceContext,
		(*WellKnown).GetSourceContext,
		&sourcecontextpb.SourceContext{},
	)

	testMutator(
		t,
		(*WellKnown).SetStringValue,
		(*WellKnown).GetStringValue,
		&wrapperspb.StringValue{},
	)

	testMutator(
		t,
		(*WellKnown).SetStruct,
		(*WellKnown).GetStruct,
		&structpb.Struct{},
	)

	testMutator(
		t,
		(*WellKnown).SetSyntax,
		(*WellKnown).GetSyntax,
		typepb.Syntax_SYNTAX_PROTO3,
	)

	testMutator(
		t,
		(*WellKnown).SetTimestamp,
		(*WellKnown).GetTimestamp,
		&timestamppb.Timestamp{},
	)

	testMutator(
		t,
		(*WellKnown).SetType,
		(*WellKnown).GetType,
		&typepb.Type{},
	)

	testMutator(
		t,
		(*WellKnown).SetUInt32Value,
		(*WellKnown).GetUInt32Value,
		&wrapperspb.UInt32Value{},
	)

	testMutator(
		t,
		(*WellKnown).SetUInt64Value,
		(*WellKnown).GetUInt64Value,
		&wrapperspb.UInt64Value{},
	)

	testMutator(
		t,
		(*WellKnown).SetValue,
		(*WellKnown).GetValue,
		&structpb.Value{},
	)
}
