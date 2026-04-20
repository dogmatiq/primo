package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates builders for each message type.
//
// For open API files, primo generates WithXxx/From/Build using direct struct
// field access. For opaque API files, the same interface is generated using
// the SetXxx/HasXxx/GetXxx methods that protoc-gen-go provides, since the
// underlying xxx_hidden_* fields are inaccessible.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		generateType(code, m)
		generateConstructor(code, m)
		generateFromMethod(code, m)
		generateBuildMethod(code, m)
		generateWithMethods(code, m)
	}

	return nil
}

const (
	receiverName       = "b"
	prototypeFieldName = "prototype"
)

func builderTypeName(m *scope.Message) string {
	return m.GoTypeName + "Builder"
}

func generateConstructor(code *jen.File, m *scope.Message) {
	typeName := builderTypeName(m)
	funcName := "New" + typeName

	code.
		Commentf(
			"%s returns a builder that constructs [%s] messages.",
			funcName,
			m.GoTypeName,
		)

	code.
		Func().
		Id(funcName).
		Params().
		Params(
			jen.
				Op("*").
				Id(typeName),
		).
		Block(
			jen.
				Return().
				Op("&").
				Id(typeName).
				Values(),
		)
}

func generateType(code *jen.File, m *scope.Message) {
	code.
		Type().
		Id(builderTypeName(m)).
		Struct(
			jen.
				Id(prototypeFieldName).
				Id(m.GoTypeName),
		)
}

func generateWithMethods(code *jen.File, m *scope.Message) {
	for _, f := range m.Fields() {
		if f.OneOfOption == nil {
			generateForDiscreteField(code, f)
		} else {
			generateForOneOfOption(code, f)
		}
	}
}
