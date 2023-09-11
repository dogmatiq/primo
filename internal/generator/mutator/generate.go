package mutator

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates mutator methods for each field in each message.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		for _, f := range m.Fields() {
			if f.OneOfOption == nil {
				generateForDiscreteField(code, f)
			} else {
				generateForOneOfMember(code, f)
			}
		}
	}

	return nil
}

// generateForDiscreteField generates a mutator method for a "discrete" field, that is,
// a field that is not part of a one-of group.
func generateForDiscreteField(code *jen.File, f *scope.Field) {
	methodName := "Set" + f.GoFieldName

	code.
		Commentf(
			"%s sets the x.%s field to v.",
			methodName,
			f.GoFieldName,
		)

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(f.Message.GoTypeName),
		).
		Id(methodName).
		Params(
			jen.
				Id("v").
				Add(f.GoType()),
		).
		Block(
			jen.
				Id("x").
				Dot(f.GoFieldName).
				Op("=").
				Id("v"),
		)
}

// generateForOneOfMember generates a mutator method for a field that is a member of a
// one-of group.
func generateForOneOfMember(code *jen.File, f *scope.Field) {
	methodName := "Set" + f.OneOfOption.DiscriminatorFieldName

	code.
		Commentf(
			"%s sets the x.%s field to a [%s] value containing v",
			methodName,
			f.GoFieldName,
			f.OneOfOption.Group.GoFieldName,
		)

	code.
		Func().
		Params(
			jen.
				Id("x").
				Op("*").
				Id(f.Message.GoTypeName),
		).
		Id(methodName).
		Params(
			jen.
				Id("v").
				Add(f.GoType()),
		).
		Block(
			jen.
				Id("x").
				Dot(f.GoFieldName).
				Op("=").
				Op("&").
				Id(f.OneOfOption.DiscriminatorTypeName).
				Values(
					jen.Id(f.OneOfOption.DiscriminatorFieldName).Op(":").Id("v"),
				),
		)
}
