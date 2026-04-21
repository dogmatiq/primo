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
	methodName := "TryGet" + o.GoIdentifiers.Base

	code.
		Commentf(
			"%s returns the %s value if x.%s is set to %s.",
			methodName,
			o.GoIdentifiers.Base,
			o.Group.GoIdentifiers.Base,
			o.GoIdentifiers.Base,
		)
	code.
		Commentf(
			"Otherwise, ok is false and v is the zero-value.",
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
			jen.Id("v").Add(o.Field.GoType()),
			jen.Id("ok").Bool(),
		).
		BlockFunc(func(c *jen.Group) {
			if o.Group.Message.File.IsOpaqueAPI() {
				oneOfOptionBodyOpaque(c, o)
			} else {
				oneOfOptionBodyOpen(c, o)
			}
		})
}

// oneOfOptionBodyOpaque populates the body for TryGetXxx in opaque API files.
// It uses HasXxx() + GetXxx() instead of type-asserting on a single
// GetXxx()-returned interface value.
func oneOfOptionBodyOpaque(c *jen.Group, o *scope.OneOfOption) {
	c.If(
		jen.Id("x").
			Dot(o.GoIdentifiers.HasMethod).
			Call(),
	).Block(
		jen.Return(
			jen.Id("x").Dot(o.GoIdentifiers.GetMethod).Call(),
			jen.True(),
		),
	)
	c.Return(jen.Id("v"), jen.False())
}

// oneOfOptionBodyOpen populates the body for TryGetXxx in open API files.
// It type-asserts on the interface value returned by GetXxx().
func oneOfOptionBodyOpen(c *jen.Group, o *scope.OneOfOption) {
	c.
		If(
			jen.List(
				jen.Id("x"),
				jen.Id("ok"),
			).
				Op(":=").
				Id("x").
				Dot(o.Group.GoIdentifiers.GetMethod).
				Call().
				Assert(
					jen.Op("*").
						Id(o.GoIdentifiers.DiscriminatorType),
				).
				Op(";").
				Id("ok"),
		).Block(
		jen.Return(
			jen.Id("x").
				Dot(o.GoIdentifiers.Base),
			jen.True(),
		),
	)
	c.Return(
		jen.Id("v"),
		jen.False(),
	)
}
