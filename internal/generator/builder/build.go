package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateBuildMethod(code *jen.File, m *scope.Message) {
	methodName := "Build"

	code.
		Commentf(
			"%s returns a new [%s] containing the values configured via the builder.",
			methodName,
			m.GoTypeName,
		)
	code.Comment("")
	code.Comment("Each call returns a new message, such that future changes to the builder do")
	code.Comment("not modify previously constructed messages.")

	code.
		Func().
		Params(
			jen.
				Id(receiverName).
				Op("*").
				Id(builderTypeName(m)),
		).
		Id(methodName).
		Params().
		Params(
			jen.
				Op("*").
				Id(m.GoTypeName),
		).
		BlockFunc(
			func(code *jen.Group) {
				if m.File.IsOpaqueAPI() {
					code.
						Id("m").
						Op(":=").
						Op("&").
						Id(m.GoTypeName).
						Values()

					for _, f := range m.Fields() {
						setCall := jen.
							Id("m").
							Dot(f.GoIdentifiers.SetMethod).
							Call(
								jen.Id(receiverName).Dot(prototypeFieldName).Dot(f.GoIdentifiers.GetMethod).Call(),
							)
						if f.GoIdentifiers.HasMethod != "" {
							code.
								If(jen.Id(receiverName).Dot(prototypeFieldName).Dot(f.GoIdentifiers.HasMethod).Call()).
								Block(setCall)
						} else {
							code.Add(setCall)
						}
					}

					code.Return(jen.Id("m"))
				} else {
					code.
						Return().
						Op("&").
						Id(m.GoTypeName).
						ValuesFunc(
							func(code *jen.Group) {
								for _, f := range m.Fields() {
									if f.OneOfOption == nil {
										code.
											Line().
											Id(f.GoIdentifiers.ExportedField).
											Op(":").
											Id(receiverName).
											Dot(prototypeFieldName).
											Dot(f.GoIdentifiers.ExportedField)
									}
								}

								for _, g := range m.OneOfGroups() {
									code.
										Line().
										Id(g.GoIdentifiers.ExportedField).
										Op(":").
										Id(receiverName).
										Dot(prototypeFieldName).
										Dot(g.GoIdentifiers.ExportedField)
								}

								code.Line()
							},
						)
				}
			},
		)
}
