package mutator

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForOneOfOption generates a mutator method for a field that is a
// option in a one-of group.
func generateForOneOfOption(code *jen.File, f *scope.Field) {
	methodName := "Set" + f.OneOfOption.DiscriminatorFieldName

	code.
		Commentf(
			"%s sets the x.%s field to a [%s] value containing v",
			methodName,
			f.GoFieldName,
			f.OneOfOption.Group.GoFieldName,
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
				Op("&").
				Id(f.OneOfOption.DiscriminatorTypeName).
				Values(
					jen.Id(f.OneOfOption.DiscriminatorFieldName).Op(":").Id("v"),
				),
		)
}