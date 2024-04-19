package accessor

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateForOneOf(code *jen.File, g *scope.OneOfGroup) {
	for _, o := range g.Options {
		oneOfAccessorTryFunc(code, o)
	}
}

func oneOfAccessorTryFunc(code *jen.File, o *scope.OneOfOption) {
	methodName := "Try" + o.DiscriminatorFieldName

	code.
		Commentf(
			"%s returns the value of [%s] in one-of field x.%s.",
			methodName,
			o.DiscriminatorFieldName,
			o.Group.GoFieldName,
		)
	code.Comment("")
	code.Comment("ok returns false if the value of this type is not set.")

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
