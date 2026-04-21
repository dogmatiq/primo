package mutator

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates mutator methods for each field in each message.
//
// Opaque API files are skipped; protoc-gen-go already generates SetXxx
// methods for them.
func Generate(code *jen.File, f *scope.File) error {
	if f.IsOpaqueAPI() {
		return nil
	}

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
