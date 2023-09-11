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
			"%s sets the x.%s field to v.",
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
		Block(
			jen.
				Id("x").
				Dot(f.GoFieldName).
				Op("=").
				Id("v"),
		)
}
