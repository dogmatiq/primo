package immutable

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates mutator methods for each field in each message.
func Generate(code *jen.File, f *scope.File) error {
	for _, m := range f.Messages() {
		generateImmutableWrapper(code, m)
	}

	return nil
}
