// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: github.com/dogmatiq/primo/internal/test/mutator/wellknown.proto

package mutator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	apipb "google.golang.org/protobuf/types/known/apipb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	sourcecontextpb "google.golang.org/protobuf/types/known/sourcecontextpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	typepb "google.golang.org/protobuf/types/known/typepb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WellKnown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Any              *anypb.Any                     `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Api              *apipb.Api                     `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	BoolValue        *wrapperspb.BoolValue          `protobuf:"bytes,3,opt,name=BoolValue,proto3" json:"BoolValue,omitempty"`
	BytesValue       *wrapperspb.BytesValue         `protobuf:"bytes,4,opt,name=BytesValue,proto3" json:"BytesValue,omitempty"`
	DoubleValue      *wrapperspb.DoubleValue        `protobuf:"bytes,5,opt,name=DoubleValue,proto3" json:"DoubleValue,omitempty"`
	Duration         *durationpb.Duration           `protobuf:"bytes,6,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Empty            *emptypb.Empty                 `protobuf:"bytes,7,opt,name=Empty,proto3" json:"Empty,omitempty"`
	Enum             *typepb.Enum                   `protobuf:"bytes,8,opt,name=Enum,proto3" json:"Enum,omitempty"`
	EnumValue        *typepb.EnumValue              `protobuf:"bytes,9,opt,name=EnumValue,proto3" json:"EnumValue,omitempty"`
	Field            *typepb.Field                  `protobuf:"bytes,10,opt,name=Field,proto3" json:"Field,omitempty"`
	FieldCardinality typepb.Field_Cardinality       `protobuf:"varint,11,opt,name=FieldCardinality,proto3,enum=google.protobuf.Field_Cardinality" json:"FieldCardinality,omitempty"`
	FieldKind        typepb.Field_Kind              `protobuf:"varint,12,opt,name=FieldKind,proto3,enum=google.protobuf.Field_Kind" json:"FieldKind,omitempty"`
	FieldMask        *fieldmaskpb.FieldMask         `protobuf:"bytes,13,opt,name=FieldMask,proto3" json:"FieldMask,omitempty"`
	FloatValue       *wrapperspb.FloatValue         `protobuf:"bytes,14,opt,name=FloatValue,proto3" json:"FloatValue,omitempty"`
	Int32Value       *wrapperspb.Int32Value         `protobuf:"bytes,15,opt,name=Int32Value,proto3" json:"Int32Value,omitempty"`
	Int64Value       *wrapperspb.Int64Value         `protobuf:"bytes,16,opt,name=Int64Value,proto3" json:"Int64Value,omitempty"`
	ListValue        *structpb.ListValue            `protobuf:"bytes,17,opt,name=ListValue,proto3" json:"ListValue,omitempty"`
	Method           *apipb.Method                  `protobuf:"bytes,18,opt,name=Method,proto3" json:"Method,omitempty"`
	Mixin            *apipb.Mixin                   `protobuf:"bytes,19,opt,name=Mixin,proto3" json:"Mixin,omitempty"`
	NullValue        structpb.NullValue             `protobuf:"varint,20,opt,name=NullValue,proto3,enum=google.protobuf.NullValue" json:"NullValue,omitempty"`
	Option           *typepb.Option                 `protobuf:"bytes,21,opt,name=Option,proto3" json:"Option,omitempty"`
	SourceContext    *sourcecontextpb.SourceContext `protobuf:"bytes,22,opt,name=SourceContext,proto3" json:"SourceContext,omitempty"`
	StringValue      *wrapperspb.StringValue        `protobuf:"bytes,23,opt,name=StringValue,proto3" json:"StringValue,omitempty"`
	Struct           *structpb.Struct               `protobuf:"bytes,24,opt,name=Struct,proto3" json:"Struct,omitempty"`
	Syntax           typepb.Syntax                  `protobuf:"varint,25,opt,name=Syntax,proto3,enum=google.protobuf.Syntax" json:"Syntax,omitempty"`
	Timestamp        *timestamppb.Timestamp         `protobuf:"bytes,26,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Type             *typepb.Type                   `protobuf:"bytes,27,opt,name=Type,proto3" json:"Type,omitempty"`
	UInt32Value      *wrapperspb.UInt32Value        `protobuf:"bytes,28,opt,name=UInt32Value,proto3" json:"UInt32Value,omitempty"`
	UInt64Value      *wrapperspb.UInt64Value        `protobuf:"bytes,29,opt,name=UInt64Value,proto3" json:"UInt64Value,omitempty"`
	Value            *structpb.Value                `protobuf:"bytes,30,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *WellKnown) Reset() {
	*x = WellKnown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WellKnown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WellKnown) ProtoMessage() {}

func (x *WellKnown) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WellKnown.ProtoReflect.Descriptor instead.
func (*WellKnown) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescGZIP(), []int{0}
}

func (x *WellKnown) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *WellKnown) GetApi() *apipb.Api {
	if x != nil {
		return x.Api
	}
	return nil
}

func (x *WellKnown) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *WellKnown) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *WellKnown) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *WellKnown) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *WellKnown) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *WellKnown) GetEnum() *typepb.Enum {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *WellKnown) GetEnumValue() *typepb.EnumValue {
	if x != nil {
		return x.EnumValue
	}
	return nil
}

func (x *WellKnown) GetField() *typepb.Field {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *WellKnown) GetFieldCardinality() typepb.Field_Cardinality {
	if x != nil {
		return x.FieldCardinality
	}
	return typepb.Field_Cardinality(0)
}

func (x *WellKnown) GetFieldKind() typepb.Field_Kind {
	if x != nil {
		return x.FieldKind
	}
	return typepb.Field_Kind(0)
}

func (x *WellKnown) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *WellKnown) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *WellKnown) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *WellKnown) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *WellKnown) GetListValue() *structpb.ListValue {
	if x != nil {
		return x.ListValue
	}
	return nil
}

func (x *WellKnown) GetMethod() *apipb.Method {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *WellKnown) GetMixin() *apipb.Mixin {
	if x != nil {
		return x.Mixin
	}
	return nil
}

func (x *WellKnown) GetNullValue() structpb.NullValue {
	if x != nil {
		return x.NullValue
	}
	return structpb.NullValue(0)
}

func (x *WellKnown) GetOption() *typepb.Option {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *WellKnown) GetSourceContext() *sourcecontextpb.SourceContext {
	if x != nil {
		return x.SourceContext
	}
	return nil
}

func (x *WellKnown) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *WellKnown) GetStruct() *structpb.Struct {
	if x != nil {
		return x.Struct
	}
	return nil
}

func (x *WellKnown) GetSyntax() typepb.Syntax {
	if x != nil {
		return x.Syntax
	}
	return typepb.Syntax(0)
}

func (x *WellKnown) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WellKnown) GetType() *typepb.Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *WellKnown) GetUInt32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.UInt32Value
	}
	return nil
}

func (x *WellKnown) GetUInt64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.UInt64Value
	}
	return nil
}

func (x *WellKnown) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x0d, 0x0a, 0x09, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x26, 0x0a, 0x03,
	0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x03, 0x61, 0x6e, 0x79, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x38, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4e, 0x0a, 0x10, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x09, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x3b, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a,
	0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x52, 0x05, 0x4d, 0x69, 0x78, 0x69, 0x6e,
	0x12, 0x38, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0d, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52, 0x06, 0x53, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x29, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_goTypes = []interface{}{
	(*WellKnown)(nil),                     // 0: primo.test.mutators.WellKnown
	(*anypb.Any)(nil),                     // 1: google.protobuf.Any
	(*apipb.Api)(nil),                     // 2: google.protobuf.Api
	(*wrapperspb.BoolValue)(nil),          // 3: google.protobuf.BoolValue
	(*wrapperspb.BytesValue)(nil),         // 4: google.protobuf.BytesValue
	(*wrapperspb.DoubleValue)(nil),        // 5: google.protobuf.DoubleValue
	(*durationpb.Duration)(nil),           // 6: google.protobuf.Duration
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
	(*typepb.Enum)(nil),                   // 8: google.protobuf.Enum
	(*typepb.EnumValue)(nil),              // 9: google.protobuf.EnumValue
	(*typepb.Field)(nil),                  // 10: google.protobuf.Field
	(typepb.Field_Cardinality)(0),         // 11: google.protobuf.Field.Cardinality
	(typepb.Field_Kind)(0),                // 12: google.protobuf.Field.Kind
	(*fieldmaskpb.FieldMask)(nil),         // 13: google.protobuf.FieldMask
	(*wrapperspb.FloatValue)(nil),         // 14: google.protobuf.FloatValue
	(*wrapperspb.Int32Value)(nil),         // 15: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),         // 16: google.protobuf.Int64Value
	(*structpb.ListValue)(nil),            // 17: google.protobuf.ListValue
	(*apipb.Method)(nil),                  // 18: google.protobuf.Method
	(*apipb.Mixin)(nil),                   // 19: google.protobuf.Mixin
	(structpb.NullValue)(0),               // 20: google.protobuf.NullValue
	(*typepb.Option)(nil),                 // 21: google.protobuf.Option
	(*sourcecontextpb.SourceContext)(nil), // 22: google.protobuf.SourceContext
	(*wrapperspb.StringValue)(nil),        // 23: google.protobuf.StringValue
	(*structpb.Struct)(nil),               // 24: google.protobuf.Struct
	(typepb.Syntax)(0),                    // 25: google.protobuf.Syntax
	(*timestamppb.Timestamp)(nil),         // 26: google.protobuf.Timestamp
	(*typepb.Type)(nil),                   // 27: google.protobuf.Type
	(*wrapperspb.UInt32Value)(nil),        // 28: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil),        // 29: google.protobuf.UInt64Value
	(*structpb.Value)(nil),                // 30: google.protobuf.Value
}
var file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_depIdxs = []int32{
	1,  // 0: primo.test.mutators.WellKnown.any:type_name -> google.protobuf.Any
	2,  // 1: primo.test.mutators.WellKnown.api:type_name -> google.protobuf.Api
	3,  // 2: primo.test.mutators.WellKnown.BoolValue:type_name -> google.protobuf.BoolValue
	4,  // 3: primo.test.mutators.WellKnown.BytesValue:type_name -> google.protobuf.BytesValue
	5,  // 4: primo.test.mutators.WellKnown.DoubleValue:type_name -> google.protobuf.DoubleValue
	6,  // 5: primo.test.mutators.WellKnown.Duration:type_name -> google.protobuf.Duration
	7,  // 6: primo.test.mutators.WellKnown.Empty:type_name -> google.protobuf.Empty
	8,  // 7: primo.test.mutators.WellKnown.Enum:type_name -> google.protobuf.Enum
	9,  // 8: primo.test.mutators.WellKnown.EnumValue:type_name -> google.protobuf.EnumValue
	10, // 9: primo.test.mutators.WellKnown.Field:type_name -> google.protobuf.Field
	11, // 10: primo.test.mutators.WellKnown.FieldCardinality:type_name -> google.protobuf.Field.Cardinality
	12, // 11: primo.test.mutators.WellKnown.FieldKind:type_name -> google.protobuf.Field.Kind
	13, // 12: primo.test.mutators.WellKnown.FieldMask:type_name -> google.protobuf.FieldMask
	14, // 13: primo.test.mutators.WellKnown.FloatValue:type_name -> google.protobuf.FloatValue
	15, // 14: primo.test.mutators.WellKnown.Int32Value:type_name -> google.protobuf.Int32Value
	16, // 15: primo.test.mutators.WellKnown.Int64Value:type_name -> google.protobuf.Int64Value
	17, // 16: primo.test.mutators.WellKnown.ListValue:type_name -> google.protobuf.ListValue
	18, // 17: primo.test.mutators.WellKnown.Method:type_name -> google.protobuf.Method
	19, // 18: primo.test.mutators.WellKnown.Mixin:type_name -> google.protobuf.Mixin
	20, // 19: primo.test.mutators.WellKnown.NullValue:type_name -> google.protobuf.NullValue
	21, // 20: primo.test.mutators.WellKnown.Option:type_name -> google.protobuf.Option
	22, // 21: primo.test.mutators.WellKnown.SourceContext:type_name -> google.protobuf.SourceContext
	23, // 22: primo.test.mutators.WellKnown.StringValue:type_name -> google.protobuf.StringValue
	24, // 23: primo.test.mutators.WellKnown.Struct:type_name -> google.protobuf.Struct
	25, // 24: primo.test.mutators.WellKnown.Syntax:type_name -> google.protobuf.Syntax
	26, // 25: primo.test.mutators.WellKnown.Timestamp:type_name -> google.protobuf.Timestamp
	27, // 26: primo.test.mutators.WellKnown.Type:type_name -> google.protobuf.Type
	28, // 27: primo.test.mutators.WellKnown.UInt32Value:type_name -> google.protobuf.UInt32Value
	29, // 28: primo.test.mutators.WellKnown.UInt64Value:type_name -> google.protobuf.UInt64Value
	30, // 29: primo.test.mutators.WellKnown.Value:type_name -> google.protobuf.Value
	30, // [30:30] is the sub-list for method output_type
	30, // [30:30] is the sub-list for method input_type
	30, // [30:30] is the sub-list for extension type_name
	30, // [30:30] is the sub-list for extension extendee
	0,  // [0:30] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WellKnown); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_wellknown_proto_depIdxs = nil
}
