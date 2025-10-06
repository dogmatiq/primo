package marshaling

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates MarshalBinary() and UnmarshalBinary() methods for each
// message type.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		generateMarshalBinary(code, m)
		generateUnmarshalBinary(code, m)
	}

	return nil
}

// generateMarshalBinary generates a MarshalBinary() method for the message.
func generateMarshalBinary(code *jen.File, m *scope.Message) {
	code.Comment("MarshalBinary returns the binary representation of the message, equivalent to")
	code.Comment("calling proto.Marshal(x).")
	code.Comment("")
	code.Commentf("It allows [*%s] to implement [encoding.BinaryMarshaler].", m.GoTypeName)

	code.
		Func().
		Params(
			jen.Id("x").Op("*").Id(m.GoTypeName),
		).
		Id("MarshalBinary").
		Params().
		Params(
			jen.Index().Byte(),
			jen.Error(),
		).
		Block(
			jen.Return().
				Qual("google.golang.org/protobuf/proto", "Marshal").
				Call(jen.Id("x")),
		)
}

// generateUnmarshalBinary generates an UnmarshalBinary() method for the message.
func generateUnmarshalBinary(code *jen.File, m *scope.Message) {
	code.Comment("UnmarshalBinary populates x from its binary representation, equivalent to")
	code.Comment("calling proto.Unmarshal(data, x).")
	code.Comment("")
	code.Commentf("It allows [*%s] to implement [encoding.BinaryUnmarshaler].", m.GoTypeName)

	code.
		Func().
		Params(
			jen.Id("x").Op("*").Id(m.GoTypeName),
		).
		Id("UnmarshalBinary").
		Params(
			jen.Id("data").Index().Byte(),
		).
		Error().
		Block(
			jen.Return().
				Qual("google.golang.org/protobuf/proto", "Unmarshal").
				Call(
					jen.Id("data"),
					jen.Id("x"),
				),
		)
}
