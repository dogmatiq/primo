package grpcstub

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/identifier"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateClientStub(code *jen.File, s *scope.Service) {
	interfaceName := identifier.Exported(s.Name) + "Client"
	structName := interfaceName + "Stub"

	code.
		Commentf(
			"%s is a test implementation of the [%s] interface.",
			structName,
			interfaceName,
		)

	code.
		Var().
		Id("_").
		Id(interfaceName).
		Op("=").
		Parens(
			jen.
				Op("*").
				Id(structName),
		).
		Parens(
			jen.Nil(),
		)

	var methods []jen.Code

	code.
		Type().
		Id(structName).
		StructFunc(
			func(code *jen.Group) {
				for _, m := range s.Methods {
					methodName := identifier.Exported(m.Name)
					fieldName := methodName + "Func"

					inputs := func(code *jen.Group) {
						code.
							Id("ctx").
							Qual("context", "Context")

						if !m.ClientStreaming {
							code.
								Id("req").
								Add(m.InputGoType())
						}

						code.
							Id("options").
							Op("...").
							Qual("google.golang.org/grpc", "CallOption")
					}

					outputs := func(code *jen.Group) {
						if m.ClientStreaming || m.ServerStreaming {
							code.Id(m.StreamingClientGoTypeName)
						} else {
							code.Add(m.OutputGoType())
						}
						code.Error()
					}

					code.
						Commentf("%s is a function that implements the %s RPC method.", fieldName, m.Name).
						Line().
						Id(fieldName).
						Func().
						ParamsFunc(inputs).
						ParamsFunc(outputs).
						Line()

					methods = append(
						methods,
						jen.
							Commentf("%s calls c.%s if it is non-nil. Otherwise, it returns an error.", methodName, fieldName).
							Line().
							Func().
							Params(
								jen.
									Id("c").
									Op("*").
									Id(structName),
							).
							Id(methodName).
							ParamsFunc(inputs).
							ParamsFunc(outputs).
							BlockFunc(
								func(code *jen.Group) {
									code.
										If(
											jen.
												Id("c").
												Dot(fieldName).
												Op("==").
												Nil(),
										).
										Block(
											jen.
												Return(
													jen.Nil(),
													jen.
														Qual(
															"google.golang.org/grpc/status",
															"Error",
														).
														Call(
															jen.Qual("google.golang.org/grpc/codes", "Unimplemented"),
															jen.Lit(fmt.Sprintf("method %s not implemented", m.Name)),
														),
												),
										)

									code.
										Return(
											jen.
												Id("c").
												Dot(fieldName).
												CallFunc(
													func(code *jen.Group) {
														code.
															Id("ctx")

														if !m.ClientStreaming {
															code.
																Id("req")
														}

														code.
															Id("options").
															Op("...")
													},
												),
										)
								},
							),
					)
				}
			},
		)

	for _, m := range methods {
		code.
			Add(m).
			Line()
	}
}
