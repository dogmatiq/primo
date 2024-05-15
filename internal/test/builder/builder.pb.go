// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: github.com/dogmatiq/primo/internal/test/builder/builder.proto

package builder

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA int32  `protobuf:"varint,1,opt,name=field_a,json=fieldA,proto3" json:"field_a,omitempty"`
	FieldB string `protobuf:"bytes,2,opt,name=field_b,json=fieldB,proto3" json:"field_b,omitempty"`
	// Types that are assignable to Group:
	//
	//	*Message_FieldC
	//	*Message_FieldD
	Group  isMessage_Group `protobuf_oneof:"group"`
	Nested *Message_Nested `protobuf:"bytes,5,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetFieldA() int32 {
	if x != nil {
		return x.FieldA
	}
	return 0
}

func (x *Message) GetFieldB() string {
	if x != nil {
		return x.FieldB
	}
	return ""
}

func (m *Message) GetGroup() isMessage_Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (x *Message) GetFieldC() int32 {
	if x, ok := x.GetGroup().(*Message_FieldC); ok {
		return x.FieldC
	}
	return 0
}

func (x *Message) GetFieldD() int32 {
	if x, ok := x.GetGroup().(*Message_FieldD); ok {
		return x.FieldD
	}
	return 0
}

func (x *Message) GetNested() *Message_Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

type isMessage_Group interface {
	isMessage_Group()
}

type Message_FieldC struct {
	FieldC int32 `protobuf:"varint,3,opt,name=field_c,json=fieldC,proto3,oneof"`
}

type Message_FieldD struct {
	FieldD int32 `protobuf:"varint,4,opt,name=field_d,json=fieldD,proto3,oneof"`
}

func (*Message_FieldC) isMessage_Group() {}

func (*Message_FieldD) isMessage_Group() {}

type Message_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"varint,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *Message_Nested) Reset() {
	*x = Message_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Nested) ProtoMessage() {}

func (x *Message_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Nested.ProtoReflect.Descriptor instead.
func (*Message_Nested) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_Nested) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_github_com_dogmatiq_primo_internal_test_builder_builder_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0xd6, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x12, 0x19, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x12, 0x19, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x12, 0x3a, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x1a, 0x1e, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61,
	0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: primo.test.builder.Message
	(*Message_Nested)(nil), // 1: primo.test.builder.Message.Nested
}
var file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_depIdxs = []int32{
	1, // 0: primo.test.builder.Message.nested:type_name -> primo.test.builder.Message.Nested
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_builder_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Nested); i {
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
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_FieldC)(nil),
		(*Message_FieldD)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_builder_builder_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_builder_builder_proto_depIdxs = nil
}
