package immutable

import (
	"fmt"
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

func generateMutateCloneMethod(code *jen.File, typeName string, m *scope.Message) {
	methodName := "MutateClone"

	code.
		Commentf(
			"%s returns a clone of x with the given set of mutations applied.",
			methodName,
		)

	code.
		Func().
		Params(
			jen.
				Id("x").
				Id(typeName),
		).
		Id(methodName).
		Params(
			jen.
				Id("mutators").
				Op("...").
				Func().
				Params(
					jen.
						Op("*").
						Id(m.GoTypeName),
				),
		).
		Params(
			jen.Id(typeName),
		).
		BlockFunc(
			func(code *jen.Group) {
				// If there are no mutators then this is essentially a no-op,
				// just return the original message without actually cloning.
				code.
					If(
						jen.
							Len(jen.Id("mutators")).
							Op("==").
							Lit(0),
					).
					Block(
						jen.Return(
							jen.Id("x"),
						),
					).
					Line()

				// Otherwise, make a clone BEFORE applying the mutators so we
				// don't mess with any other immutable wrappers that may
				// reference the same message.
				code.
					Id("m").
					Op(":=").
					Qual("google.golang.org/protobuf/proto", "Clone").
					Call(
						jen.
							Id("x").
							Dot("m"),
					).
					Op(".").
					Call(
						jen.Op("*").
							Id(m.GoTypeName),
					).
					Line()

				// Apply the mutations.
				code.
					For(
						jen.
							Id("_").
							Op(",").
							Id("fn").
							Op(":=").
							Range().
							Id("mutators"),
					).
					Block(
						jen.
							Id("fn").
							Call(
								jen.Id("m"),
							),
					).
					Line()

				generateMutateCloneForOneOf(code, m)

				code.
					Return(
						jen.
							Id(typeName).
							Values(jen.Id("m")),
					)

			},
		).
		Line()
}

func generateMutateCloneForOneOf(code *jen.Group, m *scope.Message) {
	for _, g := range m.OneOfGroups() {
		code.
			If(
				jen.
					Id("m").
					Dot(g.GoFieldName).
					Op("!=").
					Nil().
					Op("&&").
					Id("m").
					Dot(g.GoFieldName).
					Op("!=").
					Id("x").
					Dot("m").
					Dot(g.GoFieldName),
			).
			BlockFunc(
				func(code *jen.Group) {
					for _, g := range m.OneOfGroups() {
						code.
							Switch(
								jen.
									Id("v").
									Op(":=").
									Id("m").
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
											BlockFunc(
												func(code *jen.Group) {
													code.
														Id("m").
														Dot(g.GoFieldName).
														Op("=").
														Op("&").
														Id(o.DiscriminatorTypeName).
														ValuesFunc(
															func(code *jen.Group) {
																if !o.Field.HasReferenceSemantics() {
																	code.
																		Id("v").
																		Dot(o.DiscriminatorFieldName)
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
														)
												},
											)
									}
								},
							)
					}
				},
			).
			Line()
	}
}

// // If any of the fields with reference semantics no longer point
// // to the same value we make A SECOND CLONE, as these values may
// // be shared with other non-mutable messages.

// for _, f := range m.Fields() {
// 	if f.OneOfOption != nil {
// 		continue
// 	}

// 	if !f.HasReferenceSemantics() {
// 		continue
// 	}

// 	var cond jen.Code
// 	switch f.Kind() {
// 	case reflect.Slice, reflect.Map:
// 		cond = jen.
// 			Parens(
// 				jen.
// 					Id("x").
// 					Dot("m").
// 					Dot(f.GoFieldName).
// 					Op("!=").
// 					Nil().
// 					Op("||").
// 					Qual("reflect", "ValueOf").
// 					Call(
// 						jen.
// 							Id("m").
// 							Dot(f.GoFieldName),
// 					).
// 					Dot("Pointer").
// 					Call().
// 					Op("!=").
// 					Qual("reflect", "ValueOf").
// 					Call(
// 						jen.
// 							Id("x").
// 							Dot("m").
// 							Dot(f.GoFieldName),
// 					).
// 					Dot("Pointer").
// 					Call(),
// 			)
// 	case reflect.Pointer:
// 		cond = jen.
// 			Id("m").
// 			Dot(f.GoFieldName).
// 			Op("!=").
// 			Id("x").
// 			Dot("m").
// 			Dot(f.GoFieldName)
// 	default:
// 		panic(fmt.Sprintf("unsupported kind: %v", f.Kind()))
// 	}

// 	code.
// 		If(
// 			jen.
// 				Id("m").
// 				Dot(f.GoFieldName).
// 				Op("!=").
// 				Nil().
// 				Op("&&").
// 				Add(cond),
// 		).
// 		BlockFunc(
// 			func(code *jen.Group) {
// 				switch f.Kind() {
// 				case reflect.Slice:
// 				case reflect.Map:
// 				// 	cond = jen.
// 				// 		Parens(
// 				// 			jen.
// 				// 				Id("x").
// 				// 				Dot("m").
// 				// 				Dot(f.GoFieldName).
// 				// 				Op("!=").
// 				// 				Nil().
// 				// 				Op("||").
// 				// 				Qual("reflect", "ValueOf").
// 				// 				Call(
// 				// 					jen.
// 				// 						Id("m").
// 				// 						Dot(f.GoFieldName),
// 				// 				).
// 				// 				Dot("Pointer").
// 				// 				Call().
// 				// 				Op("!=").
// 				// 				Qual("reflect", "ValueOf").
// 				// 				Call(
// 				// 					jen.
// 				// 						Id("x").
// 				// 						Dot("m").
// 				// 						Dot(f.GoFieldName),
// 				// 				).
// 				// 				Dot("Pointer").
// 				// 				Call(),
// 				// 		)
// 				case reflect.Pointer:
// 					// 	cond = jen.
// 					// 		Id("m").
// 					// 		Dot(f.GoFieldName).
// 					// 		Op("!=").
// 					// 		Id("x").
// 					// 		Dot("m").
// 					// 		Dot(f.GoFieldName)
// 					// default:
// 					// 	panic(fmt.Sprintf("unsupported kind: %v", f.Kind()))
// 				}
// 			},
// 		)
// }

// for _, f := range m.Fields() {
// 	if f.OneOfOption != nil {
// 		continue
// 	}

// 	if f.HasReferenceSemantics() {
// 		continue
// 	}

// }
