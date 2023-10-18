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
				for _, f := range m.Fields() {
					if f.OneOfOption == nil {
						code.
							Id(receiverName).
							Dot(prototypeFieldName).
							Dot(f.GoFieldName).
							Op("=").
							Id("x").
							Dot(f.GoFieldName)
					}
				}

				for _, g := range m.OneOfGroups() {
					code.
						Id(receiverName).
						Dot(prototypeFieldName).
						Dot(g.GoFieldName).
						Op("=").
						Id("x").
						Dot(g.GoFieldName)
				}

				code.
					Return().
					Id(receiverName)
			},
		)
}
