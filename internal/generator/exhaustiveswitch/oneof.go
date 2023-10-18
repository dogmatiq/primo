package exhaustiveswitch

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForOneOf generates a switch function for the options within a one-of
// group.
func generateForOneOf(code *jen.File, g *scope.OneOfGroup) {
	name := "Switch_" + g.Message.GoTypeName + "_" + g.GoFieldName
	const paramPrefix = "case"

	code.
		Commentf(
			"%s dispatches to one of the given functions based on",
			name,
		)
	code.
		Commentf(
			"which value of the [%s] message's %q one-of group is populated.",
			g.Message.GoTypeName,
			g.GoFieldName,
		)
	code.Comment("")
	code.Commentf(
		"It panics if x.%s field is nil; otherwise, it returns the value",
		g.GoFieldName,
	)
	code.Comment("returned by the called function. If no return value is required, use a return")
	code.Comment("type of [error] and always return nil.")

	code.
		Func().
		Id(name).
		Types(
			jen.Id("T").Any(),
		).
		ParamsFunc(
			func(code *jen.Group) {
				code.
					Line().
					Id("x").
					Op("*").
					Id(g.Message.GoTypeName)

				for _, o := range g.Options {
					code.
						Line().
						Id(paramPrefix + o.DiscriminatorFieldName).
						Func().
						Params(
							o.Field.GoType(),
						).
						Params(
							jen.Id("T"),
						)
				}

				code.Line()
			},
		).
		Params(
			jen.Id("T"),
		).
		Block(
			jen.
				Switch(
					jen.
						Id("v").
						Op(":=").
						Id("x").
						Dot(g.GoFieldName).
						Op(".").
						Call(jen.Type()),
				).
				BlockFunc(
					func(code *jen.Group) {
						for _, o := range g.Options {
							code.
								Case(
									jen.
										Op("*").
										Id(o.DiscriminatorTypeName),
								).
								Return().
								Id(paramPrefix + o.DiscriminatorFieldName).
								Call(
									jen.
										Id("v").
										Dot(o.DiscriminatorFieldName),
								)
						}

						code.
							Default().
							Panic(
								jen.Lit(
									fmt.Sprintf(
										"%s: x.%s is nil",
										name,
										g.GoFieldName,
									),
								),
							)
					},
				),
		)
}
