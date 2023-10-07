package mutator

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForDiscreteField generates a mutator method for a "discrete" field,
// that is, a field that is not part of a one-of group.
func generateForDiscreteField(code *jen.File, f *scope.Field) {
	methodName := "Set" + f.GoFieldName

	code.
		Commentf(
			"%s sets the x.%s field to v, then returns x.",
			methodName,
			f.GoFieldName,
		)

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(f.Message.GoTypeName),
		).
		Id(methodName).
		Params(
			jen.
				Id("v").
				Add(f.GoType()),
		).
		Params(
			jen.
				Op("*").
				Id(f.Message.GoTypeName),
		).
		Block(
			jen.If(
				jen.
					Id("x").
					Op("==").
					Nil(),
			).
				Block(
					jen.
						Id("x").
						Op("=").
						Op("&").
						Id(f.Message.GoTypeName).
						Values(),
				),
			jen.
				Id("x").
				Dot(f.GoFieldName).
				Op("=").
				Id("v"),
			jen.
				Return().
				Id("x"),
		)
}
