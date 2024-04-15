package immutable

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateImmutableWrapper(code *jen.File, m *scope.Message) {
	typeName := "Immutable" + m.GoTypeName

	code.
		Commentf(
			"%s is an immutable wrapper for [%s] messages.",
			typeName,
			m.GoTypeName,
		)

	code.
		Type().
		Id(typeName).
		Struct(
			jen.
				Id("m").
				Op("*").
				Id(m.GoTypeName),
		)

	generateFreezeMethod(code, typeName, m)
	generateThawMethod(code, typeName, m)
	generateMutateCloneMethod(code, typeName, m)

	for _, f := range m.Fields() {
		generateAccessor(code, typeName, f)
	}
}

func generateFreezeMethod(code *jen.File, typeName string, m *scope.Message) {
	code.
		Commentf("Freeze returns a clone of x that cannot be modified.")

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(m.GoTypeName),
		).
		Id("Freeze").
		Params().
		Params(
			jen.Id(typeName),
		).
		Block(
			jen.Return(
				jen.
					Id(typeName).
					Values(
						jen.
							Qual("google.golang.org/protobuf/proto", "Clone").
							Call(
								jen.Id("x"),
							).
							Op(".").
							Call(
								jen.
									Op("*").
									Id(m.GoTypeName),
							),
					),
			),
		).
		Line()
}

func generateThawMethod(code *jen.File, typeName string, m *scope.Message) {
	code.
		Commentf("Thaw sets the fields of x to a deep-clone of src.")

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(m.GoTypeName),
		).
		Id("Thaw").
		Params(
			jen.
				Id("src").
				Id(typeName),
		).
		BlockFunc(
			func(code *jen.Group) {
				code.
					Qual("google.golang.org/protobuf/proto", "Reset").
					Call(
						jen.Id("x"),
					)

				code.
					Qual("google.golang.org/protobuf/proto", "Merge").
					Call(
						jen.
							Id("x"),
						jen.
							Qual("google.golang.org/protobuf/proto", "Clone").
							Call(
								jen.
									Id("src").
									Dot("m"),
							).
							Op(".").
							Call(
								jen.
									Op("*").
									Id(m.GoTypeName),
							),
					)
			},
		).
		Line()
}
