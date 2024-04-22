package accessor

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateForOneOf(code *jen.File, g *scope.OneOfGroup) {
	for _, o := range g.Options {
		generateForOneOfOption(code, o)
	}
}

func generateForOneOfOption(code *jen.File, o *scope.OneOfOption) {
	methodName := "TryGet" + o.DiscriminatorFieldName

	code.
		Commentf(
			"%s returns x.%s.%s if x.%s is a [%s].",
			methodName,
			o.Group.GoFieldName,
			o.DiscriminatorFieldName,
			o.Group.GoFieldName,
			o.DiscriminatorTypeName,
		)
	code.
		Commentf(
			"Otherwise, ok is false v is the zero-value.",
		)

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(o.Group.Message.GoTypeName),
		).
		Id(methodName).
		Params().
		Params(
			jen.
				Id("v").
				Add(o.Field.GoType()),
			jen.
				Id("ok").
				Add(jen.Bool()),
		).
		Block(
			jen.
				If(
					jen.List(
						jen.Id("x"),
						jen.Id("ok"),
					).
						Op(":=").
						Id("x").
						Dot("Get"+o.Group.GoFieldName).
						Call().
						Assert(
							jen.Op("*").
								Id(o.DiscriminatorTypeName),
						).
						Op(";").
						Id("ok"),
				).Block(
				jen.Return(
					jen.Id("x").
						Dot(o.DiscriminatorFieldName),
					jen.True(),
				),
			),
			jen.Return(
				jen.Id("v"),
				jen.False(),
			),
		)
}
