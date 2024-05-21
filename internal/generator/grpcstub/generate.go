package grpcstub

import (
	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
)

// Generate generates stub implementations for gRPC client interfaces.
func Generate(code *jen.File, f *scope.File) error {
	for _, s := range f.Services() {
		generateClientStub(code, s)
	}

	return nil
}
