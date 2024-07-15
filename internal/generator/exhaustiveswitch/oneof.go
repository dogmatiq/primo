package exhaustiveswitch

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateForOneOf(code *jen.File, g *scope.OneOfGroup) {
	generateOneOfSwitch(code, g, true)
	generateOneOfSwitch(code, g, false)
	generateOneOfMap(code, g, true)
	generateOneOfMap(code, g, false)
}

func oneOfCaseFuncName(o *scope.OneOfOption) string {
	return "case" + o.DiscriminatorFieldName
}

func generateOneOfSwitch(code *jen.File, g *scope.OneOfGroup, panicOnNil bool) {
	funcName := "Switch_" + g.Message.GoTypeName + "_" + g.GoFieldName
	nilBehavior := "calls none()"

	if panicOnNil {
		funcName = "Must" + funcName
		nilBehavior = "panics"
	}

	code.
		Commentf(
			"%s invokes one of the given functions based on",
			funcName,
		)
	code.
		Commentf(
			"the value of x.%s.",
			g.GoFieldName,
		)
	code.Comment("")
	code.
		Commentf(
			"It %s if x.%s is nil.",
			nilBehavior,
			g.GoFieldName,
		)

	code.
		Func().
		Id(funcName).
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
						Id(oneOfCaseFuncName(o)).
						Func().
						Params(
							o.Field.GoType(),
						)
				}

				if !panicOnNil {
					code.
						Line().
						Id("none").
						Func().
						Params()
				}

				code.Line()
			},
		).
		Block(
			jen.
				Switch(
					jen.
						Id("v").
						Op(":=").
						Id("x").
						Dot("Get" + g.GoFieldName).
						Call().
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
								Id(oneOfCaseFuncName(o)).
								Call(
									jen.
										Id("v").
										Dot(o.DiscriminatorFieldName),
								)
						}

						if panicOnNil {
							code.
								Default().
								Panic(
									jen.Lit(
										fmt.Sprintf(
											"%s: x.%s is nil",
											funcName,
											g.GoFieldName,
										),
									),
								)
						} else {
							code.
								Default().
								Id("none").
								Call()
						}
					},
				),
		)
}

func generateOneOfMap(code *jen.File, g *scope.OneOfGroup, panicOnNil bool) {
	funcName := "Map_" + g.Message.GoTypeName + "_" + g.GoFieldName
	nilBehavior := "calls none()"

	if panicOnNil {
		funcName = "Must" + funcName
		nilBehavior = "panics"
	}

	code.
		Commentf(
			"%s maps x.%s to a value of type T by invoking",
			funcName,
			g.GoFieldName,
		)
	code.Comment("one of the given functions.")
	code.Comment("")
	code.Commentf(
		"It invokes the function that corresponds to the value of x.%s,",
		g.GoFieldName,
	)
	code.Commentf(
		"and returns that function's result. It %s if x.%s is nil.",
		nilBehavior,
		g.GoFieldName,
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
					Id("x").
					Op("*").
					Id(g.Message.GoTypeName)

				for _, o := range g.Options {
					code.
						Line().
						Id(oneOfCaseFuncName(o)).
						Func().
						Params(
							o.Field.GoType(),
						).
						Params(
							jen.Id("T"),
						)
				}

				if !panicOnNil {
					code.
						Line().
						Id("none").
						Func().
						Params().
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
						Dot("Get" + g.GoFieldName).
						Call().
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
								Id(oneOfCaseFuncName(o)).
								Call(
									jen.
										Id("v").
										Dot(o.DiscriminatorFieldName),
								)
						}

						if panicOnNil {
							code.
								Default().
								Panic(
									jen.Lit(
										fmt.Sprintf(
											"%s: x.%s is nil",
											funcName,
											g.GoFieldName,
										),
									),
								)
						} else {
							code.
								Default().
								Return().
								Id("none").
								Call()
						}
					},
				),
		)
}
