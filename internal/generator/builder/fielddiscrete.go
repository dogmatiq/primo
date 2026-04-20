package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForDiscreteField generates a builder method for a "discrete" field,
// that is, a field that is not part of a one-of group.
func generateForDiscreteField(code *jen.File, f *scope.Field) {
	typeName := builderTypeName(f.Message)
	methodName := "With" + f.GoIdentifiers.Base

	code.
		Commentf(
			"%s configures the builder to set the %s field to v,",
			methodName,
			f.GoIdentifiers.Base,
		)
	code.
		Commentf(
			"then returns %s.",
			receiverName,
		)

	paramType := f.GoType()
	if f.Message.File.IsOpaqueAPI() {
		paramType = f.GoElemType()
	}

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
				Add(paramType),
		).
		Params(
			jen.
				Op("*").
				Id(typeName),
		).
		BlockFunc(
			func(code *jen.Group) {
				if f.Message.File.IsOpaqueAPI() {
					code.
						Id(receiverName).
						Dot(prototypeFieldName).
						Dot(f.GoIdentifiers.SetMethod).
						Call(jen.Id("v"))
				} else {
					code.
						Id(receiverName).
						Dot(prototypeFieldName).
						Dot(f.GoIdentifiers.ExportedField).
						Op("=").
						Id("v")
				}
				code.
					Return().
					Id(receiverName)
			},
		)
}
