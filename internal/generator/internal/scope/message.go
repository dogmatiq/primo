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
			goFieldName := identifier.Exported(d.GetName())
			g := &OneOfGroup{
				Message:    m,
				Descriptor: d,
			}
			g.GoIdentifiers.Base = goFieldName
			if m.File.IsOpaqueAPI() {
				g.GoIdentifiers.WhichMethod = "Which" + goFieldName
			} else {
				g.GoIdentifiers.ExportedField = goFieldName
				g.GoIdentifiers.GetMethod = "Get" + goFieldName
			}
			m.groups = append(m.groups, g)
		}

		for _, d := range m.Descriptor.GetField() {
			goFieldName := identifier.Exported(d.GetName())

			if fieldNameCollidesWithMethod(d) {
				goFieldName += "_"
			}

			f := &Field{
				Message:    m,
				Descriptor: d,
			}
			f.GoIdentifiers.Base = goFieldName

			// Proto3 optional fields have a synthetic oneof_decl entry in the
			// descriptor, but they are not real oneofs. Skip the oneof branch
			// for these fields so that they are treated as plain optional
			// fields.
			if d.OneofIndex != nil && !d.GetProto3Optional() {
				group := m.groups[d.GetOneofIndex()]

				typeName := m.GoTypeName + "_" + goFieldName
				if m.File.Request.hasGoType(m.File.GoPackagePath, typeName) {
					typeName += "_"
				}

				f.OneOfOption = &OneOfOption{
					Group: group,
					Field: f,
				}
				f.OneOfOption.GoIdentifiers.Base = goFieldName
				if m.File.IsOpaqueAPI() {
					f.OneOfOption.GoIdentifiers.HasMethod = "Has" + goFieldName
					f.OneOfOption.GoIdentifiers.GetMethod = "Get" + goFieldName
					f.OneOfOption.GoIdentifiers.CaseConstant = m.GoTypeName + "_" + goFieldName + "_case"
				} else {
					f.OneOfOption.GoIdentifiers.DiscriminatorType = typeName
				}

				group.Options = append(group.Options, f.OneOfOption)

				if m.File.IsOpaqueAPI() {
					f.GoIdentifiers.HasMethod = "Has" + goFieldName
					f.GoIdentifiers.GetMethod = "Get" + goFieldName
					f.GoIdentifiers.SetMethod = "Set" + goFieldName
				} else {
					f.GoIdentifiers.ExportedField = group.GoIdentifiers.ExportedField
				}
			} else {
				if m.File.IsOpaqueAPI() {
					f.GoIdentifiers.GetMethod = "Get" + goFieldName
					f.GoIdentifiers.SetMethod = "Set" + goFieldName
					if f.hasExplicitPresence() {
						f.GoIdentifiers.HasMethod = "Has" + goFieldName
					}
				} else {
					f.GoIdentifiers.ExportedField = goFieldName
				}
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
