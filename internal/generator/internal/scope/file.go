package scope

import (
	"github.com/dogmatiq/primo/internal/generator/internal/identifier"
	"google.golang.org/protobuf/types/descriptorpb"
)

// File encapsulates information about a single .proto file.
type File struct {
	Request       *Request
	Descriptor    *descriptorpb.FileDescriptorProto
	GoPackagePath string
	GoPackageName string

	services []*Service
	messages []*Message
	enums    []*Enum
}

// Services returns the set of services defined in this file.
func (f *File) Services() []*Service {
	if f.services != nil {
		return f.services
	}

	for _, d := range f.Descriptor.GetService() {
		s := &Service{
			File:       f,
			Descriptor: d,
			Name:       d.GetName(),
		}

		for _, md := range d.GetMethod() {
			clientName := ""
			if md.GetClientStreaming() || md.GetServerStreaming() {
				clientName = identifier.Exported(s.Name) + "_" + identifier.Exported(md.GetName()) + "Client"
			}

			s.Methods = append(
				s.Methods,
				&Method{
					Service:                   s,
					Descriptor:                md,
					Name:                      md.GetName(),
					ClientStreaming:           md.GetClientStreaming(),
					ServerStreaming:           md.GetServerStreaming(),
					StreamingClientGoTypeName: clientName,
				},
			)
		}

		f.services = append(f.services, s)
	}

	return f.services
}

// Messages returns the set of messages defined in this file.
func (f *File) Messages() []*Message {
	if f.messages != nil {
		return f.messages
	}

	var collect func(*Message, []*descriptorpb.DescriptorProto)
	collect = func(parent *Message, messages []*descriptorpb.DescriptorProto) {
		for _, d := range messages {
			if d.GetOptions().GetMapEntry() {
				continue
			}

			m := &Message{
				File:       f,
				Descriptor: d,
				GoTypeName: identifier.Exported(d.GetName()),
			}

			if parent != nil {
				m.GoTypeName = parent.GoTypeName + "_" + m.GoTypeName
			}

			f.messages = append(f.messages, m)
			collect(m, d.GetNestedType())
		}
	}

	collect(nil, f.Descriptor.GetMessageType())

	return f.messages
}

// Enums returns the set of enum types defined in this file.
func (f *File) Enums() []*Enum {
	if f.enums == nil {
		for _, d := range f.Descriptor.GetEnumType() {
			n := identifier.Exported(d.GetName())

			f.enums = append(
				f.enums,
				&Enum{
					File:             f,
					Descriptor:       d,
					GoTypeName:       n,
					GoConstantPrefix: n + "_",
				},
			)
		}

		for _, m := range f.Messages() {
			for _, d := range m.Descriptor.GetEnumType() {
				f.enums = append(
					f.enums,
					&Enum{
						File:             f,
						Descriptor:       d,
						GoTypeName:       m.GoTypeName + "_" + identifier.Exported(d.GetName()),
						GoConstantPrefix: m.GoTypeName + "_",
					},
				)
			}
		}
	}
	return f.enums
}
