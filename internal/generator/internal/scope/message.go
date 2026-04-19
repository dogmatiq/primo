package scope

import (
	"slices"

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

			// Proto3 optional fields have a synthetic oneof_decl entry in the
			// descriptor, but they are not real oneofs. Skip the oneof branch
			// for these fields so that they are treated as plain optional
			// fields.
			if d.OneofIndex != nil && !d.GetProto3Optional() {
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
					DiscriminatorTypeName:  typeName,
					DiscriminatorFieldName: fieldName,
				}

				group.Options = append(group.Options, f.OneOfOption)
			}

			m.fields = append(m.fields, f)
		}

		// Remove synthetic oneof groups (those created for proto3 optional
		// fields) which have no options after field processing.
		m.groups = slices.DeleteFunc(m.groups, func(g *OneOfGroup) bool {
			return len(g.Options) == 0
		})
	}

	return m.fields
}
