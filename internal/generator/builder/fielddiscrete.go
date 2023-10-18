package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForDiscreteField generates a builder method for a "discrete" field,
// that is, a field that is not part of a one-of group.
func generateForDiscreteField(code *jen.File, f *scope.Field) {
	typeName := builderTypeName(f.Message)
	methodName := "With" + f.GoFieldName

	code.
		Commentf(
			"%s configures the builder to set the %s field to v,",
			methodName,
			f.GoFieldName,
		)
	code.
		Commentf(
			"then returns %s.",
			receiverName,
		)

	code.
		Func().
		Params(
			jen.
				Id(receiverName).
				Op("*").
				Id(typeName),
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
				Id(typeName),
		).
		Block(
			jen.
				Id(receiverName).
				Dot(prototypeFieldName).
				Dot(f.GoFieldName).
				Op("=").
				Id("v"),
			jen.
				Return().
				Id(receiverName),
		)
}
