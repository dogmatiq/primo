package exhaustiveswitch

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// generateForEnum generates a switch function for the members of an enum type.
func generateForEnum(code *jen.File, e *scope.Enum) {
	const (
		paramPrefix         = "case"
		discriminatorSuffix = "_Case"
	)

	name := "Switch_" + e.GoTypeName

	code.
		Type().
		DefsFunc(
			func(code *jen.Group) {
				for _, m := range e.Members() {
					caseName := m.GoConstantName + discriminatorSuffix

					code.
						Commentf(
							"%s is a type that identifies a function that is",
							caseName,
						)
					code.
						Commentf(
							"invoked by [%s] when the [%s] is [%s].",
							name,
							e.GoTypeName,
							m.GoConstantName,
						)

					if m.AliasFor == nil {
						code.
							Id(caseName).
							Struct()
					} else {
						code.
							Id(caseName).
							Op("=").
							Id(m.AliasFor.GoConstantName + discriminatorSuffix)
					}
				}
			},
		)

	code.
		Commentf(
			"%s accepts a set of functions that correspond to the",
			name,
		)
	code.
		Commentf(
			"members of the [%s] enumeration.",
			e.GoTypeName,
		)
	code.Comment("")
	code.Comment("It invokes the function that corresponds to v, and returns that function's")
	code.Comment("return value. If no return value is required, use a return type of [error]")
	code.Comment("and always return nil.")

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
					Id("v").
					Id(e.GoTypeName)

				for _, m := range e.Members() {
					if m.AliasFor != nil {
						continue
					}

					code.
						Line().
						Id(paramPrefix + m.Descriptor.GetName()).
						Func().
						Params(
							jen.Id(m.GoConstantName + discriminatorSuffix),
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
					jen.Id("v"),
				).
				BlockFunc(
					func(code *jen.Group) {
						for _, m := range e.Members() {
							if m.AliasFor != nil {
								continue
							}
							code.
								Case(jen.Id(m.GoConstantName)).
								Return().
								Id(paramPrefix + m.Descriptor.GetName()).
								Call(
									jen.
										Id(m.GoConstantName + discriminatorSuffix).
										Values(),
								)
						}

						code.
							Default().
							Panic(
								jen.
									Qual("fmt", "Sprintf").
									Call(
										jen.Lit(
											fmt.Sprintf(
												"%s: %%d is not a valid %s",
												name,
												e.GoTypeName,
											),
										),
										jen.Id("v"),
									),
							)
					},
				),
		)
}
