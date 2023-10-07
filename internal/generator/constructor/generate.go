package constructor

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates constructor functions for each message.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		generate(code, m)
	}

	return nil
}

// generate generates a constructor function for a message.
func generate(code *jen.File, m *scope.Message) {
	name := "New" + m.GoTypeName

	code.
		Commentf(
			"%s returns a new [%s].",
			name,
			m.GoTypeName,
		)

	code.
		Func().
		Id(name).
		Params().
		Params(
			jen.
				Op("*").
				Id(m.GoTypeName),
		).
		Block(
			jen.
				Return().
				Op("&").
				Id(m.GoTypeName).
				Values(),
		)
}
