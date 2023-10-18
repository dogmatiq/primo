package exhaustiveswitch

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateForEnum(code *jen.File, e *scope.Enum) {
	generateEnumDiscriminators(code, e)
	generateEnumSwitch(code, e)
	generateEnumMap(code, e)
}

func enumCaseFuncName(m *scope.EnumMember) string {
	return "case" + m.Descriptor.GetName()
}

func enumDiscriminatorTypeName(m *scope.EnumMember) string {
	return m.GoConstantName + "_Case"
}
func generateEnumDiscriminators(code *jen.File, e *scope.Enum) {
	code.
		Type().
		DefsFunc(
			func(code *jen.Group) {
				for _, m := range e.Members() {
					name := enumDiscriminatorTypeName(m)

					code.
						Commentf(
							"%s is a type that statically associates a function",
							name,
						)
					code.
						Commentf(
							"with a [%s] value.",
							m.GoConstantName,
						)

					if m.AliasFor == nil {
						code.
							Id(name).
							Struct()
					} else {
						code.
							Id(name).
							Op("=").
							Id(enumDiscriminatorTypeName(m.AliasFor))
					}
				}
			},
		)
}

func generateEnumSwitch(code *jen.File, e *scope.Enum) {
	funcName := "Switch_" + e.GoTypeName

	code.
		Commentf(
			"%s dispatches to a function based on the value of v.",
			funcName,
		)
	code.Comment("")
	code.Comment("It invokes the function that corresponds to v. It panics if v is not a")
	code.
		Commentf(
			"recognized [%s] value.",
			e.GoTypeName,
		)

	code.
		Func().
		Id(funcName).
		ParamsFunc(
			func(code *jen.Group) {
				code.
					Line().
					Id("v").
					Id(e.GoTypeName)

				for _, m := range e.Members() {
					if m.AliasFor == nil {
						code.
							Line().
							Id(enumCaseFuncName(m)).
							Func().
							Params(
								jen.Id(enumDiscriminatorTypeName(m)),
							)
					}
				}

				code.Line()
			},
		).
		Block(
			jen.
				Switch(
					jen.Id("v"),
				).
				BlockFunc(
					func(code *jen.Group) {
						for _, m := range e.Members() {
							if m.AliasFor == nil {
								code.
									Case(jen.Id(m.GoConstantName)).
									Id(enumCaseFuncName(m)).
									Call(
										jen.
											Id(enumDiscriminatorTypeName(m)).
											Values(),
									)
							}
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
												funcName,
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

func generateEnumMap(code *jen.File, e *scope.Enum) {
	funcName := "Map_" + e.GoTypeName

	code.
		Commentf(
			"%s maps a member of the [%s] enumeration to a",
			funcName,
			e.GoTypeName,
		)
	code.Comment("value of type T.")
	code.Comment("")
	code.Comment("It invokes the function that corresponds to v, and returns that function's")
	code.Commentf(
		"result. It panics if v is not a recognized [%s] value.",
		e.GoTypeName,
	)

	code.
		Func().
		Id(funcName).
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
					if m.AliasFor == nil {
						code.
							Line().
							Id(enumCaseFuncName(m)).
							Func().
							Params(
								jen.Id(enumDiscriminatorTypeName(m)),
							).
							Params(
								jen.Id("T"),
							)
					}
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
							if m.AliasFor == nil {
								code.
									Case(jen.Id(m.GoConstantName)).
									Return().
									Id(enumCaseFuncName(m)).
									Call(
										jen.
											Id(enumDiscriminatorTypeName(m)).
											Values(),
									)
							}
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
												funcName,
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
