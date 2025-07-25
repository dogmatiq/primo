// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: github.com/dogmatiq/primo/internal/test/grpcstub/service.proto

package grpcstub

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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescGZIP(), []int{1}
}

var File_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto protoreflect.FileDescriptor

const file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDesc = "" +
	"\n" +
	">github.com/dogmatiq/primo/internal/test/grpcstub/service.proto\x12\x13primo.test.grpcstub\"\t\n" +
	"\aRequest\"\n" +
	"\n" +
	"\bResponse2\xcd\x02\n" +
	"\aService\x12F\n" +
	"\x05Unary\x12\x1c.primo.test.grpcstub.Request\x1a\x1d.primo.test.grpcstub.Response\"\x00\x12O\n" +
	"\fServerStream\x12\x1c.primo.test.grpcstub.Request\x1a\x1d.primo.test.grpcstub.Response\"\x000\x01\x12O\n" +
	"\fClientStream\x12\x1c.primo.test.grpcstub.Request\x1a\x1d.primo.test.grpcstub.Response\"\x00(\x01\x12X\n" +
	"\x13BidirectionalStream\x12\x1c.primo.test.grpcstub.Request\x1a\x1d.primo.test.grpcstub.Response\"\x00(\x010\x01B2Z0github.com/dogmatiq/primo/internal/test/grpcstubb\x06proto3"

var (
	file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescData []byte
)

func file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDesc), len(file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDesc)))
	})
	return file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_goTypes = []any{
	(*Request)(nil),  // 0: primo.test.grpcstub.Request
	(*Response)(nil), // 1: primo.test.grpcstub.Response
}
var file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_depIdxs = []int32{
	0, // 0: primo.test.grpcstub.Service.Unary:input_type -> primo.test.grpcstub.Request
	0, // 1: primo.test.grpcstub.Service.ServerStream:input_type -> primo.test.grpcstub.Request
	0, // 2: primo.test.grpcstub.Service.ClientStream:input_type -> primo.test.grpcstub.Request
	0, // 3: primo.test.grpcstub.Service.BidirectionalStream:input_type -> primo.test.grpcstub.Request
	1, // 4: primo.test.grpcstub.Service.Unary:output_type -> primo.test.grpcstub.Response
	1, // 5: primo.test.grpcstub.Service.ServerStream:output_type -> primo.test.grpcstub.Response
	1, // 6: primo.test.grpcstub.Service.ClientStream:output_type -> primo.test.grpcstub.Response
	1, // 7: primo.test.grpcstub.Service.BidirectionalStream:output_type -> primo.test.grpcstub.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDesc), len(file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_depIdxs,
		MessageInfos:      file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_msgTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_grpcstub_service_proto_depIdxs = nil
}
