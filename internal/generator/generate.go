package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	protoPackage = "google.golang.org/protobuf/proto"
	rootPackage  = "github.com/dogmatiq/primo"
)

// Generate generates Go code for the given request.
func Generate(
	ver string,
	req *pluginpb.CodeGeneratorRequest,
) *pluginpb.CodeGeneratorResponse {
	res := &pluginpb.CodeGeneratorResponse{}

	if err := generate(ver, req, res); err != nil {
		res.Error = proto.String(err.Error())
	}

	return res
}

func generate(
	ver string,
	req *pluginpb.CodeGeneratorRequest,
	res *pluginpb.CodeGeneratorResponse,
) error {
	params, err := parseParameters(req.GetParameter())
	if err != nil {
		return err
	}

	types := map[string]jen.Code{}

	for _, fd := range req.GetProtoFile() {
		pkgPath, _, err := goPackage(fd)
		if err != nil {
			continue
		}

		for _, md := range fd.GetMessageType() {
			q := fmt.Sprintf(
				".%s.%s",
				fd.GetPackage(),
				md.GetName(),
			)
			types[q] = jen.Op("*").Qual(pkgPath, camelCase(md.GetName()))
		}

		for _, ed := range fd.GetEnumType() {
			q := fmt.Sprintf(
				".%s.%s",
				fd.GetPackage(),
				ed.GetName(),
			)
			types[q] = jen.Qual(pkgPath, camelCase(ed.GetName()))
		}
	}

	for _, n := range req.GetFileToGenerate() {
		for _, fd := range req.GetProtoFile() {
			if fd.GetName() != n {
				continue
			}

			fs := &fileScope{
				GeneratorVersion: ver,
				Request:          req,
				Types:            types,
				Parameters:       params,
				FileDescriptor:   fd,
			}

			fr, err := fs.Generate()
			if err != nil {
				return err
			}
			res.File = append(res.File, fr)
		}
	}

	return nil
}
