package scope

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/primo/internal/generator/internal/identifier"
	"github.com/dogmatiq/primo/internal/generator/internal/option"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// Request describes a request to generate code from one or more .proto files.
type Request struct {
	GeneratorVersion string
	PluginRequest    *pluginpb.CodeGeneratorRequest
	PluginParameters Parameters

	pbTypes map[string]jen.Code                      // fully-qualified protocol buffers type name -> Go type expression
	pbMaps  map[string]*descriptorpb.DescriptorProto // fully-qualified protocol buffers type name -> map descriptor
	goTypes map[string]map[string]struct{}           // go package path -> go type name -> struct{}
}

// Parameters encapsulates the options passed to the generator via the
// --primo_opt flag.
//
// The are referred to as "options" on the protoc command line, but "parameters"
// within the Protocol Buffers plugin system.
type Parameters struct {
	Module string
}

// typeExpr returns the Go type expression that refers to the given protocol
// buffers type.
//
// n is the fully-qualified (that is, with leading dot) protocol buffers type
// name.
func (r *Request) typeExpr(n string) jen.Code {
	populateTypes(r)

	if expr, ok := r.pbTypes[n]; ok {
		return expr
	}

	panic("unknown type: " + n)
}

// hasGoType returns true if there is a generated go type in the given package
// with the given name.
func (r *Request) hasGoType(pkgPath, name string) bool {
	populateTypes(r)

	_, ok := r.goTypes[pkgPath][name]
	return ok
}

func (r *Request) isMap(fd *descriptorpb.FieldDescriptorProto) bool {
	populateTypes(r)
	_, ok := r.pbMaps[fd.GetTypeName()]
	return ok
}

func (r *Request) typeExprForField(fd *descriptorpb.FieldDescriptorProto) jen.Code {
	populateTypes(r)

	if md, ok := r.pbMaps[fd.GetTypeName()]; ok {
		var key, value jen.Code

		for _, mfd := range md.GetField() {
			if mfd.GetNumber() == 1 {
				key = r.typeExprForField(mfd)
			} else {
				value = r.typeExprForField(mfd)
			}
		}

		return jen.Map(key).Add(value)
	}

	expr := typeExprForFieldType(r, fd)

	if fd.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return jen.Index().Add(expr)
	}

	return expr
}

func typeExprForFieldType(r *Request, fd *descriptorpb.FieldDescriptorProto) jen.Code {
	switch fd.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return jen.Int32()
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return jen.Uint32()
	case descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return jen.Int64()
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return jen.Uint64()
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return jen.Float32()
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return jen.Float64()
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return jen.Bool()
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return jen.String()
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return jen.Index().Byte()
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		return r.typeExpr(fd.GetTypeName())
	default:
		panic(fmt.Sprintf("unsupported field type: %s", fd.GetType()))
	}
}

func populateTypes(r *Request) {
	if r.pbTypes != nil {
		return
	}

	r.pbTypes = map[string]jen.Code{}
	r.pbMaps = map[string]*descriptorpb.DescriptorProto{}
	r.goTypes = map[string]map[string]struct{}{}

	for _, fd := range r.PluginRequest.GetProtoFile() {
		goPackage, _, err := option.GoPackage(fd)
		if err != nil {
			// Ignore any protocol buffers files for which there is no Go
			// code.
			continue
		}

		for _, md := range fd.GetMessageType() {
			buildMessageExpression(
				r,
				fd.GetPackage(),
				goPackage,
				nil, // nesting
				md,
			)
		}

		for _, ed := range fd.GetEnumType() {
			buildEnumExpression(
				r,
				fd.GetPackage(),
				goPackage,
				nil, // nesting
				ed,
			)
		}
	}
}

func buildMessageExpression(
	r *Request,
	protoPackage string,
	goPackage string,
	nesting []string,
	md *descriptorpb.DescriptorProto,
) {
	parts := append(nesting, md.GetName())
	protoName := "." + protoPackage + "." + strings.Join(parts, ".")

	if md.GetOptions().GetMapEntry() {
		r.pbMaps[protoName] = md
	} else {
		goType := identifier.Exported(strings.Join(parts, "_"))
		r.pbTypes[protoName] = jen.Op("*").Qual(goPackage, goType)
		registerGoType(r, goPackage, goType)
	}

	for _, nd := range md.GetNestedType() {
		buildMessageExpression(
			r,
			protoPackage,
			goPackage,
			append(nesting, md.GetName()),
			nd,
		)
	}

	for _, ed := range md.GetEnumType() {
		buildEnumExpression(
			r,
			protoPackage,
			goPackage,
			append(nesting, md.GetName()),
			ed,
		)
	}
}

func buildEnumExpression(
	r *Request,
	protoPackage string,
	goPackage string,
	nesting []string,
	ed *descriptorpb.EnumDescriptorProto,
) {
	parts := append(nesting, ed.GetName())
	pbName := "." + protoPackage + "." + strings.Join(parts, ".")
	goType := identifier.Exported(strings.Join(parts, "_"))

	r.pbTypes[pbName] = jen.Qual(goPackage, goType)
	registerGoType(r, goPackage, goType)
}

func registerGoType(r *Request, pkgPath, name string) {
	types, ok := r.goTypes[pkgPath]
	if !ok {
		types = map[string]struct{}{}
		r.goTypes[pkgPath] = types
	}
	types[name] = struct{}{}
}
