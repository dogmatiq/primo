package scope

import "google.golang.org/protobuf/types/descriptorpb"

// OneOfGroup describes a one-of group.
type OneOfGroup struct {
	Message     *Message
	Descriptor  *descriptorpb.OneofDescriptorProto
	GoFieldName string
	Options     []*OneOfOption
}

// OneOfOption describes a field that is part of a one-of.
type OneOfOption struct {
	Group                  *OneOfGroup
	Field                  *Field
	DiscriminatorTypeName  string
	DiscriminatorFieldName string
}
