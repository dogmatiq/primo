package scope

import (
	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Service encapsulates information about a service defined in a .proto file.
type Service struct {
	File       *File
	Descriptor *descriptorpb.ServiceDescriptorProto
	Name       string
	Methods    []*Method
}

// Method encapsulates information about an RPC method within in a [Service].
type Method struct {
	Service                   *Service
	Descriptor                *descriptorpb.MethodDescriptorProto
	Name                      string
	ClientStreaming           bool
	ServerStreaming           bool
	StreamingClientGoTypeName string
}

// InputGoType returns the code snippet that refers to the type of the method's
// input message.
func (m *Method) InputGoType() jen.Code {
	return m.Service.File.Request.typeExpr(m.Descriptor.GetInputType())
}

// OutputGoType returns the code snippet that refers to the type of the method's
// output message.
func (m *Method) OutputGoType() jen.Code {
	return m.Service.File.Request.typeExpr(m.Descriptor.GetOutputType())
}
