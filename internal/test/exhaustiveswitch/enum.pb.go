// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: github.com/dogmatiq/primo/internal/test/exhaustiveswitch/enum.proto

package exhaustiveswitch

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

type Color int32

const (
	Color_UNKNOWN Color = 0
	Color_RED     Color = 1
	Color_GREEN   Color = 2
	Color_BLUE    Color = 3
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "UNKNOWN",
		1: "RED",
		2: "GREEN",
		3: "BLUE",
	}
	Color_value = map[string]int32{
		"UNKNOWN": 0,
		"RED":     1,
		"GREEN":   2,
		"BLUE":    3,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescGZIP(), []int{0}
}

var File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto protoreflect.FileDescriptor

var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67,
	0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x68, 0x61, 0x75, 0x73,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x69, 0x6d, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x65, 0x78, 0x68, 0x61, 0x75, 0x73, 0x74, 0x69, 0x76, 0x65, 0x73, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x2a, 0x32, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x6d, 0x61, 0x74, 0x69, 0x71, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x65, 0x78, 0x68, 0x61, 0x75, 0x73, 0x74, 0x69, 0x76, 0x65, 0x73, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescOnce sync.Once
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescData = file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDesc
)

func file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescGZIP() []byte {
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescOnce.Do(func() {
		file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescData)
	})
	return file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDescData
}

var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_goTypes = []interface{}{
	(Color)(0), // 0: primo.test.exhaustiveswitch.Color
}
var file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_init() }
func file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_init() {
	if File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_goTypes,
		DependencyIndexes: file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_depIdxs,
		EnumInfos:         file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_enumTypes,
	}.Build()
	File_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto = out.File
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_rawDesc = nil
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_goTypes = nil
	file_github_com_dogmatiq_primo_internal_test_exhaustiveswitch_enum_proto_depIdxs = nil
}
