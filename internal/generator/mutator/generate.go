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
				generateForOneOfOption(code, f)
			}
		}
	}

	return nil
}
