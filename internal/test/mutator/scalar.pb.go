// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: github.com/dogmatiq/primo/internal/test/mutator/scalar.proto

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

type Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32    int32   `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64    int64   `protobuf:"varint,2,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32   uint32  `protobuf:"varint,3,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64   uint64  `protobuf:"varint,4,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Sint32   int32   `protobuf:"zigzag32,5,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64   int64   `protobuf:"zigzag64,6,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Fixed32  uint32  `protobuf:"fixed32,7,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64  uint64  `protobuf:"fixed64,8,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32 int32   `protobuf:"fixed32,9,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64 int64   `protobuf:"fixed64,10,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Float    float32 `protobuf:"fixed32,11,opt,name=float,proto3" json:"float,omitempty"`
	Double   float64 `protobuf:"fixed64,12,opt,name=double,proto3" json:"double,omitempty"`
	Bool     bool    `protobuf:"varint,13,opt,name=bool,proto3" json:"bool,omitempty"`
	String_  string  `protobuf:"bytes,14,opt,name=string,proto3" json:"string,omitempty"`
	Bytes    []byte  `protobuf:"bytes,15,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Scalar) Reset() {
	*x = Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar) ProtoMessage() {}

func (x *Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar.ProtoReflect.Descriptor instead.
func (*Scalar) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescGZIP(), []int{0}
}

func (x *Scalar) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Scalar) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Scalar) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Scalar) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Scalar) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Scalar) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Scalar) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Scalar) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Scalar) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Scalar) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Scalar) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Scalar) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Scalar) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Scalar) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Scalar) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x75, 0x74, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0xf0, 0x02, 0x0a, 0x06, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f,
	0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_goTypes = []any{
	(*Scalar)(nil), // 0: primo.test.mutators.Scalar
}
var file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Scalar); i {
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
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_mutator_scalar_proto_depIdxs = nil
}
