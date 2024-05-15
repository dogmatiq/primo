package scope

import (
	"github.com/dogmatiq/primo/internal/generator/internal/identifier"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Enum encapsulates information about a single enum type.
type Enum struct {
	File             *File
	Descriptor       *descriptorpb.EnumDescriptorProto
	GoTypeName       string
	GoConstantPrefix string

	members []*EnumMember
}

// Members returns members of the enum.
func (e *Enum) Members() []*EnumMember {
	if e.members == nil {
		canonical := map[int32]*EnumMember{}

		for _, d := range e.Descriptor.GetValue() {
			m := &EnumMember{
				Enum:           e,
				Descriptor:     d,
				GoConstantName: e.GoConstantPrefix + identifier.Exported(d.GetName()),
			}

			if c, ok := canonical[d.GetNumber()]; ok {
				m.AliasFor = c
			} else {
				canonical[d.GetNumber()] = m
			}

			e.members = append(e.members, m)
		}
	}

	return e.members
}

// EnumMember encapsulates information about a single member of an enum.
type EnumMember struct {
	Enum           *Enum
	AliasFor       *EnumMember
	Descriptor     *descriptorpb.EnumValueDescriptorProto
	GoConstantName string
}
