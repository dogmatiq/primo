package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForOneOfOption generates a builder method for a field that is a
// option in a one-of group.
func generateForOneOfOption(code *jen.File, f *scope.Field) {
	typeName := builderTypeName(f.Message)
	methodName := "With" + f.OneOfOption.GoIdentifiers.Base

	if f.Message.File.IsOpaqueAPI() {
		code.
			Commentf(
				"%s configures the builder to set the %s field to v,",
				methodName,
				f.OneOfOption.GoIdentifiers.Base,
			)
		code.
			Commentf(
				"then returns %s.",
				receiverName,
			)
	} else {
		code.
			Commentf(
				"%s configures the builder to set the %s field to a",
				methodName,
				f.OneOfOption.Group.GoIdentifiers.Base,
			)
		code.
			Commentf(
				"[%s] value containing v, then returns %s",
				f.OneOfOption.GoIdentifiers.DiscriminatorType,
				receiverName,
			)
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
				Add(f.GoType()),
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
						Op("&").
						Id(f.OneOfOption.GoIdentifiers.DiscriminatorType).
						Values(
							jen.Id(f.OneOfOption.GoIdentifiers.Base).Op(":").Id("v"),
						)
				}
				code.
					Return().
					Id(receiverName)
			},
		)
}
