package generator

import (
	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/types/descriptorpb"
)

type messageScope struct {
	*fileScope
	MessageDescriptor *descriptorpb.DescriptorProto
}

func (s *messageScope) Generate(code *jen.File) error {
	messageName := camelCase(s.MessageDescriptor.GetName())
	oneOfs := s.MessageDescriptor.GetOneofDecl()

	for _, fd := range s.MessageDescriptor.GetField() {
		if fd.OneofIndex == nil {
			continue
		}

		fieldName := camelCase(oneOfs[fd.GetOneofIndex()].GetName())
		optionName := camelCase(fd.GetName())
		methodName := "Set" + optionName
		discriminatorName := messageName + "_" + optionName

		code.
			Commentf(
				"%s sets x.%s to a [%s] value.",
				methodName,
				fieldName,
				discriminatorName,
			)

		code.
			Func().
			Params(
				jen.
					Id("x").
					Op("*").
					Id(messageName),
			).
			Id(methodName).
			Params(
				jen.
					Id("v").
					Add(s.fieldGoType(fd)),
			).
			Block(
				jen.
					Id("x").
					Dot(fieldName).
					Op("=").
					Op("&").
					Id(discriminatorName).
					Values(
						jen.Id(optionName).Op(":").Id("v"),
					),
			)
	}

	return nil
}

func (s *messageScope) fieldGoType(fd *descriptorpb.FieldDescriptorProto) jen.Code {
	switch fd.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return jen.Int32()
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return jen.Uint32()
	case descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return jen.Int64()
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return jen.Uint64()
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return jen.Float32()
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return jen.Float64()
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return jen.Bool()
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return jen.String()
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return jen.Index().Byte()
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP,
		descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		return jen.Add(s.Types[fd.GetTypeName()])
	default:
		panic("unknown field type")
	}
}
