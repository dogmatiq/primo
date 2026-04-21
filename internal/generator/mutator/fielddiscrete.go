package mutator

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForDiscreteField generates a mutator method for a "discrete" field,
// that is, a field that is not part of a one-of group.
func generateForDiscreteField(code *jen.File, f *scope.Field) {
	methodName := "Set" + f.GoIdentifiers.Base

	code.
		Commentf(
			"%s sets the x.%s field to v, then returns x.",
			methodName,
			f.GoIdentifiers.Base,
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
		Params().
		Block(
			jen.
				Id("x").
				Dot(f.GoIdentifiers.ExportedField).
				Op("=").
				Id("v"),
		)
}
