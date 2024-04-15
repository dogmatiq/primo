package scope

import (
	"github.com/dogmatiq/primo/internal/generator/internal/identifier"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Message encapsulates information about a single message type.
type Message struct {
	File       *File
	Descriptor *descriptorpb.DescriptorProto
	GoTypeName string

	groups []*OneOfGroup
	fields []*Field
}

// OneOfGroups returns the set of one-of groups within the message.
func (m *Message) OneOfGroups() []*OneOfGroup {
	m.Fields()
	return m.groups
}

// Fields returns the set of fields within the message.
func (m *Message) Fields() []*Field {
	if m.fields == nil {
		for _, d := range m.Descriptor.GetOneofDecl() {
			m.groups = append(
				m.groups,
				&OneOfGroup{
					Message:     m,
					Descriptor:  d,
					GoFieldName: identifier.Exported(d.GetName()),
				},
			)
		}

		for _, d := range m.Descriptor.GetField() {
			f := &Field{
				Message:     m,
				Descriptor:  d,
				GoFieldName: identifier.Exported(d.GetName()),
			}

			if fieldNameCollidesWithMethod(d) {
				f.GoFieldName += "_"
			}

			if d.OneofIndex != nil {
				group := m.groups[d.GetOneofIndex()]

				fieldName := f.GoFieldName
				f.GoFieldName = identifier.Exported(group.Descriptor.GetName())

				typeName := m.GoTypeName + "_" + fieldName
				if m.File.Request.hasGoType(m.File.GoPackagePath, typeName) {
					typeName += "_"
				}

				f.OneOfOption = &OneOfOption{
					Group:                  group,
					Field:                  f,
					Position:               len(group.Options),
					DiscriminatorTypeName:  typeName,
					DiscriminatorFieldName: fieldName,
				}

				group.Options = append(group.Options, f.OneOfOption)
			}

			m.fields = append(m.fields, f)
		}
	}

	return m.fields
}
