// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: github.com/dogmatiq/primo/internal/test/exhaustiveswitch/oneof.proto

package exhaustiveswitch

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Record struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Operation:
	//
	//	*Record_Increment
	//	*Record_Decrement
	//	*Record_Log
	//	*Record_NamingCollision_
	Operation     isRecord_Operation `protobuf_oneof:"Operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetOperation() isRecord_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *Record) GetIncrement() int32 {
	if x != nil {
		if x, ok := x.Operation.(*Record_Increment); ok {
			return x.Increment
		}
	}
	return 0
}

func (x *Record) GetDecrement() int32 {
	if x != nil {
		if x, ok := x.Operation.(*Record_Decrement); ok {
			return x.Decrement
		}
	}
	return 0
}

func (x *Record) GetLog() string {
	if x != nil {
		if x, ok := x.Operation.(*Record_Log); ok {
			return x.Log
		}
	}
	return ""
}

func (x *Record) GetNamingCollision() *Record_NamingCollision {
	if x != nil {
		if x, ok := x.Operation.(*Record_NamingCollision_); ok {
			return x.NamingCollision
		}
	}
	return nil
}

type isRecord_Operation interface {
	isRecord_Operation()
}

type Record_Increment struct {
	Increment int32 `protobuf:"varint,1,opt,name=increment,proto3,oneof"`
}

type Record_Decrement struct {
	Decrement int32 `protobuf:"varint,2,opt,name=decrement,proto3,oneof"`
}

type Record_Log struct {
	Log string `protobuf:"bytes,3,opt,name=log,proto3,oneof"`
}

type Record_NamingCollision_ struct {
	NamingCollision *Record_NamingCollision `protobuf:"bytes,4,opt,name=naming_collision,json=namingCollision,proto3,oneof"`
}

func (*Record_Increment) isRecord_Operation() {}

func (*Record_Decrement) isRecord_Operation() {}

func (*Record_Log) isRecord_Operation() {}

func (*Record_NamingCollision_) isRecord_Operation() {}

type Record_NamingCollision struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record_NamingCollision) Reset() {
	*x = Record_NamingCollision{}
	mi := &file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record_NamingCollision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_NamingCollision) ProtoMessage() {}

func (x *Record_NamingCollision) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record_NamingCollision.ProtoReflect.Descriptor instead.
func (*Record_NamingCollision) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescGZIP(), []int{0, 0}
}

var File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto protoreflect.FileDescriptor

const file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDesc = "" +
	"\n" +
	"Dgithub.com/dogmatiq/primo/internal/test/exhaustiveswitch/oneof.proto\x12\x1bprimo.test.exhaustiveswitch\"\xde\x01\n" +
	"\x06Record\x12\x1e\n" +
	"\tincrement\x18\x01 \x01(\x05H\x00R\tincrement\x12\x1e\n" +
	"\tdecrement\x18\x02 \x01(\x05H\x00R\tdecrement\x12\x12\n" +
	"\x03log\x18\x03 \x01(\tH\x00R\x03log\x12`\n" +
	"\x10naming_collision\x18\x04 \x01(\v23.primo.test.exhaustiveswitch.Record.NamingCollisionH\x00R\x0fnamingCollision\x1a\x11\n" +
	"\x0fNamingCollisionB\v\n" +
	"\tOperationB:Z8github.com/dogmatiq/primo/internal/test/exhaustiveswitchb\x06proto3"

var (
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescData []byte
)

func file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDesc), len(file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDesc)))
	})
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_goTypes = []any{
	(*Record)(nil),                 // 0: primo.test.exhaustiveswitch.Record
	(*Record_NamingCollision)(nil), // 1: primo.test.exhaustiveswitch.Record.NamingCollision
}
var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_depIdxs = []int32{
	1, // 0: primo.test.exhaustiveswitch.Record.naming_collision:type_name -> primo.test.exhaustiveswitch.Record.NamingCollision
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto != nil {
		return
	}
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes[0].OneofWrappers = []any{
		(*Record_Increment)(nil),
		(*Record_Decrement)(nil),
		(*Record_Log)(nil),
		(*Record_NamingCollision_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDesc), len(file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_oneof_proto_depIdxs = nil
}
