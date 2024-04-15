package immutable

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateAccessor(code *jen.File, typeName string, f *scope.Field) {
	originalMethodName := "Get" + f.GoFieldName
	if f.OneOfOption != nil {
		originalMethodName = "Get" + f.OneOfOption.DiscriminatorFieldName
	}

	methodName := originalMethodName
	if f.HasReferenceSemantics() {
		methodName = "Clone" + strings.TrimPrefix(originalMethodName, "Get")

		code.
			Commentf(
				"%s returns a clone of the %s field.",
				methodName,
				f.GoFieldName,
			)
		code.Comment("")
		code.Commentf(
			"It cannot return the original value because [%s] is mutable.",
			f.GoType(),
		)
	} else {
		code.
			Commentf(
				"%s returns the value of the %s field.",
				methodName,
				f.GoFieldName,
			)
	}

	code.
		Func().
		Params(
			jen.
				Id("x").
				Id(typeName),
		).
		Id(methodName).
		Params().
		Params(
			jen.Add(f.GoType()),
		).
		Block(
			jen.
				ReturnFunc(
					func(code *jen.Group) {
						if !f.HasReferenceSemantics() {
							code.
								Id("x").
								Dot("m").
								Dot(originalMethodName).
								Call()
							return
						}

						switch f.Kind() {
						case reflect.Slice:
							code.
								Qual("slices", "Clone").
								Call(
									jen.Id("x").
										Dot("m").
										Dot(originalMethodName).
										Call(),
								)
						case reflect.Map:
							code.
								Qual("maps", "Clone").
								Call(
									jen.Id("x").
										Dot("m").
										Dot(originalMethodName).
										Call(),
								)
						case reflect.Pointer:
							code.
								Qual("google.golang.org/protobuf/proto", "Clone").
								Call(
									jen.
										Id("x").
										Dot("m").
										Dot(originalMethodName).
										Call(),
								).
								Op(".").
								Call(
									jen.Add(f.GoType()),
								)
						default:
							panic(fmt.Sprintf("unsupported kind: %v", f.Kind()))
						}
					},
				),
		).
		Line()
}
