package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateFromMethod(code *jen.File, m *scope.Message) {
	typeName := builderTypeName(m)
	methodName := "From"

	code.
		Commentf(
			"%s configures the builder to use x as the prototype for new messages,",
			methodName,
		)
	code.
		Commentf(
			"then returns %s.",
			receiverName,
		)

	code.Comment("")
	code.Comment("It performs a shallow copy of x, such that any changes made via the builder")
	code.Comment("do not modify x. It does not make a copy of the field values themselves.")

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
				Id("x").
				Op("*").
				Id(m.GoTypeName),
		).
		Params(
			jen.
				Op("*").
				Id(typeName),
		).
		BlockFunc(
			func(code *jen.Group) {
				if m.File.IsOpaqueAPI() {
					code.Qual("google.golang.org/protobuf/proto", "Reset").
						Call(
							jen.Op("&").Id(receiverName).Dot(prototypeFieldName),
						)
					for _, f := range m.Fields() {
						setCall := jen.
							Id(receiverName).
							Dot(prototypeFieldName).
							Dot(f.GoIdentifiers.SetMethod).
							Call(
								jen.Id("x").Dot(f.GoIdentifiers.GetMethod).Call(),
							)
						if f.GoIdentifiers.HasMethod != "" {
							code.
								If(jen.Id("x").Dot(f.GoIdentifiers.HasMethod).Call()).
								Block(setCall)
						} else {
							code.Add(setCall)
						}
					}
				} else {
					for _, f := range m.Fields() {
						if f.OneOfOption == nil {
							code.
								Id(receiverName).
								Dot(prototypeFieldName).
								Dot(f.GoIdentifiers.ExportedField).
								Op("=").
								Id("x").
								Dot(f.GoIdentifiers.ExportedField)
						}
					}

					for _, g := range m.OneOfGroups() {
						code.
							Id(receiverName).
							Dot(prototypeFieldName).
							Dot(g.GoIdentifiers.ExportedField).
							Op("=").
							Id("x").
							Dot(g.GoIdentifiers.ExportedField)
					}
				}

				code.
					Return().
					Id(receiverName)
			},
		)
}
