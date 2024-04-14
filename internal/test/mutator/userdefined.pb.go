// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: github.com/dogmatiq/primo/internal/test/mutator/userdefined.proto

package mutator

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

type UserDefinedEnum int32

const (
	UserDefinedEnum_UNKNOWN UserDefinedEnum = 0
	UserDefinedEnum_VALUE   UserDefinedEnum = 1
)

// Enum value maps for UserDefinedEnum.
var (
	UserDefinedEnum_name = map[int32]string{
		0: "UNKNOWN",
		1: "VALUE",
	}
	UserDefinedEnum_value = map[string]int32{
		"UNKNOWN": 0,
		"VALUE":   1,
	}
)

func (x UserDefinedEnum) Enum() *UserDefinedEnum {
	p := new(UserDefinedEnum)
	*p = x
	return p
}

func (x UserDefinedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserDefinedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_enumTypes[0].Descriptor()
}

func (UserDefinedEnum) Type() protoreflect.EnumType {
	return &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_enumTypes[0]
}

func (x UserDefinedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserDefinedEnum.Descriptor instead.
func (UserDefinedEnum) EnumDescriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescGZIP(), []int{0}
}

type UserDefinedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDefinedMessage) Reset() {
	*x = UserDefinedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDefinedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDefinedMessage) ProtoMessage() {}

func (x *UserDefinedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDefinedMessage.ProtoReflect.Descriptor instead.
func (*UserDefinedMessage) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescGZIP(), []int{0}
}

type UserDefined struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *UserDefinedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Enum    UserDefinedEnum     `protobuf:"varint,2,opt,name=enum,proto3,enum=primo.test.mutators.UserDefinedEnum" json:"enum,omitempty"`
	Nested  *UserDefined_Nested `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *UserDefined) Reset() {
	*x = UserDefined{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDefined) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDefined) ProtoMessage() {}

func (x *UserDefined) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDefined.ProtoReflect.Descriptor instead.
func (*UserDefined) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescGZIP(), []int{1}
}

func (x *UserDefined) GetMessage() *UserDefinedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *UserDefined) GetEnum() UserDefinedEnum {
	if x != nil {
		return x.Enum
	}
	return UserDefinedEnum_UNKNOWN
}

func (x *UserDefined) GetNested() *UserDefined_Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

type UserDefined_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDefined_Nested) Reset() {
	*x = UserDefined_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDefined_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDefined_Nested) ProtoMessage() {}

func (x *UserDefined_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDefined_Nested.ProtoReflect.Descriptor instead.
func (*UserDefined_Nested) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescGZIP(), []int{1, 0}
}

var File_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd5,
	0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x41,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x38, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x3f, 0x0a, 0x06, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72,
	0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x08, 0x0a, 0x06,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2a, 0x29, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_goTypes = []interface{}{
	(UserDefinedEnum)(0),       // 0: primo.test.mutators.UserDefinedEnum
	(*UserDefinedMessage)(nil), // 1: primo.test.mutators.UserDefinedMessage
	(*UserDefined)(nil),        // 2: primo.test.mutators.UserDefined
	(*UserDefined_Nested)(nil), // 3: primo.test.mutators.UserDefined.Nested
}
var file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_depIdxs = []int32{
	1, // 0: primo.test.mutators.UserDefined.message:type_name -> primo.test.mutators.UserDefinedMessage
	0, // 1: primo.test.mutators.UserDefined.enum:type_name -> primo.test.mutators.UserDefinedEnum
	3, // 2: primo.test.mutators.UserDefined.nested:type_name -> primo.test.mutators.UserDefined.Nested
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDefinedMessage); i {
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
		file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDefined); i {
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
		file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDefined_Nested); i {
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
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_depIdxs,
		EnumInfos:         file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_enumTypes,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_userdefined_proto_depIdxs = nil
}
