package scope

import (
	"github.com/dogmatiq/primo/internal/generator/identifier"
	"google.golang.org/protobuf/types/descriptorpb"
)

// File encapsulates information about a single .proto file.
type File struct {
	Request       *Request
	Descriptor    *descriptorpb.FileDescriptorProto
	GoPackagePath string
	GoPackageName string

	messages []*Message
}

// Messages returns the set of messages defined in this file.
func (f *File) Messages() []*Message {
	if f.messages == nil {
		for _, d := range f.Descriptor.GetMessageType() {
			f.messages = append(
				f.messages,
				&Message{
					File:       f,
					Descriptor: d,
					GoTypeName: identifier.Exported(d.GetName()),
				},
			)
		}
	}
	return f.messages
}
