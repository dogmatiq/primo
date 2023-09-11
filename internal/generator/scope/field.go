package scope

import (
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
