package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForOneOfOption generates a builder method for a field that is a
// option in a one-of group.
func generateForOneOfOption(code *jen.File, f *scope.Field) {
	typeName := builderTypeName(f.Message)
	methodName := "With" + f.OneOfOption.DiscriminatorFieldName

	code.
		Commentf(
			"%s configures the builder to set the %s field to a",
			methodName,
			f.OneOfOption.Group.GoFieldName,
		)
	code.
		Commentf(
			"[%s] value containing v, then returns %s",
			f.OneOfOption.DiscriminatorTypeName,
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
				Op("&").
				Id(f.OneOfOption.DiscriminatorTypeName).
				Values(
					jen.Id(f.OneOfOption.DiscriminatorFieldName).Op(":").Id("v"),
				),
			jen.
				Return().
				Id(receiverName),
		)
}
