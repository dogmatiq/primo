package option

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
)

// GoPackage returns the qualified name of the Go package, based on the .proto
// file's 'go_package' option.
func GoPackage(d *descriptorpb.FileDescriptorProto) (string, string, error) {
	pkg := d.GetOptions().GetGoPackage()
	if pkg == "" {
		return "", "", fmt.Errorf(
			"%s does not specify a 'go_package' option",
			d.GetName(),
		)
	}

	if i := strings.Index(pkg, ";"); i != -1 {
		// If a semi-colon is present, the part after the semi-colon is the
		// actual package name. Used when the import path and package name
		// differ.
		//
		// Use of this option is discouraged. See
		// https://developers.google.com/protocol-buffers/docs/reference/go-generated
		return pkg[:i], pkg[i+1:], nil
	}

	return pkg, path.Base(pkg), nil
}
