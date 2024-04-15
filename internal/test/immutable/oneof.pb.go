// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: github.com/dogmatiq/primo/internal/test/immutable/oneof.proto

package immutable

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OneOf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Group:
	//
	//	*OneOf_FieldA
	//	*OneOf_FieldB
	//	*OneOf_FieldC
	Group isOneOf_Group `protobuf_oneof:"group"`
}

func (x *OneOf) Reset() {
	*x = OneOf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOf) ProtoMessage() {}

func (x *OneOf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOf.ProtoReflect.Descriptor instead.
func (*OneOf) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *OneOf) GetGroup() isOneOf_Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (x *OneOf) GetFieldA() int32 {
	if x, ok := x.GetGroup().(*OneOf_FieldA); ok {
		return x.FieldA
	}
	return 0
}

func (x *OneOf) GetFieldB() int32 {
	if x, ok := x.GetGroup().(*OneOf_FieldB); ok {
		return x.FieldB
	}
	return 0
}

func (x *OneOf) GetFieldC() string {
	if x, ok := x.GetGroup().(*OneOf_FieldC); ok {
		return x.FieldC
	}
	return ""
}

type isOneOf_Group interface {
	isOneOf_Group()
}

type OneOf_FieldA struct {
	FieldA int32 `protobuf:"varint,1,opt,name=field_a,json=fieldA,proto3,oneof"` // note: two fields of the same type
}

type OneOf_FieldB struct {
	FieldB int32 `protobuf:"varint,2,opt,name=field_b,json=fieldB,proto3,oneof"`
}

type OneOf_FieldC struct {
	FieldC string `protobuf:"bytes,3,opt,name=field_c,json=fieldC,proto3,oneof"`
}

func (*OneOf_FieldA) isOneOf_Group() {}

func (*OneOf_FieldB) isOneOf_Group() {}

func (*OneOf_FieldC) isOneOf_Group() {}

var File_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6d, 0x6d, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x61, 0x0a, 0x05, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x12, 0x19,
	0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x12, 0x19, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x42, 0x12, 0x19, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x42,
	0x07, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f,
	0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_goTypes = []interface{}{
	(*OneOf)(nil), // 0: primo.test.immutable.OneOf
}
var file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOf); i {
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
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OneOf_FieldA)(nil),
		(*OneOf_FieldB)(nil),
		(*OneOf_FieldC)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_immutable_oneof_proto_depIdxs = nil
}
