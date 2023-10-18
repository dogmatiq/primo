package exhaustiveswitch

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates switch and map functions for each enum type and one-of
// group.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		for _, g := range m.OneOfGroups() {
			generateForOneOf(code, g)
		}
	}

	for _, e := range f.Enums() {
		generateForEnum(code, e)
	}

	return nil
}
