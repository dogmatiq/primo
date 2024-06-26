package generator

import (
	"fmt"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/accessor"
	"github.com/dogmatiq/primo/internal/generator/builder"
	"github.com/dogmatiq/primo/internal/generator/exhaustiveswitch"
	"github.com/dogmatiq/primo/internal/generator/grpcstub"
	"github.com/dogmatiq/primo/internal/generator/internal/option"
	"github.com/dogmatiq/primo/internal/generator/internal/scope"
	"github.com/dogmatiq/primo/internal/generator/mutator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	protoPackage = "google.golang.org/protobuf/proto"
	rootPackage  = "github.com/dogmatiq/primo"
)

// generator is a function that generates Go code for a given file.
type generator func(*jen.File, *scope.File) error

// Generate generates Go code for the given request.
func Generate(
	ver string,
	req *pluginpb.CodeGeneratorRequest,
) *pluginpb.CodeGeneratorResponse {
	res := &pluginpb.CodeGeneratorResponse{}

	if err := generate(
		ver,
		req,
		res,
		accessor.Generate,
		builder.Generate,
		exhaustiveswitch.Generate,
		grpcstub.Generate,
		mutator.Generate,
	); err != nil {
		res.Error = proto.String(err.Error())
	}

	return res
}

func generate(
	ver string,
	req *pluginpb.CodeGeneratorRequest,
	res *pluginpb.CodeGeneratorResponse,
	generators ...generator,
) error {
	params, err := parseParameters(req.GetParameter())
	if err != nil {
		return err
	}

	r := &scope.Request{
		GeneratorVersion: ver,
		PluginRequest:    req,
		PluginParameters: params,
	}

	for _, n := range req.GetFileToGenerate() {
		for _, d := range req.GetProtoFile() {
			if d.GetName() != n {
				continue
			}

			fr, err := generateFile(r, d, generators)
			if err != nil {
				return err
			}

			res.File = append(res.File, fr)
		}
	}

	return nil
}

// parseParameters parses the parameters passed to the generator.
func parseParameters(params string) (scope.Parameters, error) {
	var p scope.Parameters

	for _, k := range strings.Split(params, ",") {
		if k == "" {
			continue
		}

		var v string
		if i := strings.Index(k, "="); i != -1 {
			v = k[i+1:]
			k = k[:i]
		}

		switch k {
		case "module":
			p.Module = v
		default:
			return scope.Parameters{}, fmt.Errorf("unrecognized option: %s", k)
		}
	}

	return p, nil
}

func generateFile(
	r *scope.Request,
	d *descriptorpb.FileDescriptorProto,
	generators []generator,
) (*pluginpb.CodeGeneratorResponse_File, error) {
	pkgPath, pkgName, err := option.GoPackage(d)
	if err != nil {
		return nil, err
	}

	f := &scope.File{
		Request:       r,
		Descriptor:    d,
		GoPackagePath: pkgPath,
		GoPackageName: pkgName,
	}

	code := jen.NewFilePathName(pkgPath, pkgName)
	code.HeaderComment("Code generated by protoc-gen-go-primo. DO NOT EDIT.")
	code.HeaderComment("versions:")
	code.HeaderComment(fmt.Sprintf("// 	protoc-gen-go-primo v%s", r.GeneratorVersion))
	code.HeaderComment(fmt.Sprintf("// 	protoc              v%s", compilerVersion(r)))
	code.HeaderComment(fmt.Sprintf("// source: %s", d.GetName()))

	for _, g := range generators {
		if err := g(code, f); err != nil {
			return nil, err
		}
	}

	w := &strings.Builder{}
	if err := code.Render(w); err != nil {
		return nil, err
	}

	return &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String(outputFileName(f)),
		Content: proto.String(w.String()),
	}, nil
}

// compilerVersion returns the protoc version provided in the request for use in
// a file header.
func compilerVersion(r *scope.Request) string {
	ver := r.PluginRequest.GetCompilerVersion()

	str := fmt.Sprintf(
		"%d.%d.%d",
		ver.GetMajor(),
		ver.GetMinor(),
		ver.GetPatch(),
	)

	if suffix := ver.GetSuffix(); suffix != "" {
		str += "-" + suffix
	}

	return str
}

// outputFileName returns the name of the file to generate.
func outputFileName(f *scope.File) string {
	n := strings.TrimPrefix(
		f.Descriptor.GetName(),
		f.Request.PluginParameters.Module,
	)

	if ext := path.Ext(n); ext == ".proto" || ext == ".protodevel" {
		n = strings.TrimSuffix(n, ext)
	}

	return n + "_primo.pb.go"
}
