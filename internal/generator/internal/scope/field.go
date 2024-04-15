package scope

import (
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Field encapsulates information about a single field within a message.
type Field struct {
	Message     *Message
	Descriptor  *descriptorpb.FieldDescriptorProto
	GoFieldName string
	OneOfOption *OneOfOption
}

// GoType returns the code snippet that refers to the type of the field.
func (f *Field) GoType() jen.Code {
	return f.Message.File.Request.typeExprForField(f.Descriptor)
}

// HasReferenceSemantics returns true if the field's type behaves as though it
// is a "reference" type (e.g., pointer, interface, slice, etc).
func (f *Field) HasReferenceSemantics() bool {
	switch f.Kind() {
	case reflect.Slice, reflect.Map, reflect.Pointer, reflect.Interface:
		return true
	default:
		return false
	}
}

// Kind returns the [reflect.Kind] of the field's type.
func (f *Field) Kind() reflect.Kind {
	if f.Message.File.Request.isMap(f.Descriptor) {
		return reflect.Map
	}

	if f.Descriptor.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return reflect.Slice
	}

	switch f.Descriptor.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return reflect.Int32
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return reflect.Uint32
	case descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return reflect.Int64
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return reflect.Uint64
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return reflect.Float32
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return reflect.Float64
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return reflect.Bool
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return reflect.String
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return reflect.Slice
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		return reflect.Pointer
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		return reflect.Int32
	default:
		return reflect.Invalid
	}
}

// fieldNameCollidesWithMethod returns true if the given field name should be
// suffixed with an underscore because it would otherwise collide with one of
// the automatically generated methods.
func fieldNameCollidesWithMethod(d *descriptorpb.FieldDescriptorProto) bool {
	switch strings.ToLower(d.GetName()) {
	case "string":
		return true
	default:
		return false
	}
}
