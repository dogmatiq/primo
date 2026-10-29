package option_test

import (
	"testing"

	. "github.com/dogmatiq/primo/internal/generator/internal/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestGoPackage(t *testing.T) {
	cases := []struct {
		Name     string
		Option   string
		WantPath string
		WantName string
		WantErr  string
	}{
		{
			Name:     "implicit package name",
			Option:   "example.com/foo",
			WantPath: "example.com/foo",
			WantName: "foo",
		},
		{
			Name:     "implicit package name containing a hyphen",
			Option:   "example.com/my-module",
			WantPath: "example.com/my-module",
			WantName: "my_module",
		},
		{
			Name:     "implicit package name starting with a digit",
			Option:   "example.com/2fa",
			WantPath: "example.com/2fa",
			WantName: "_2fa",
		},
		{
			Name:     "implicit package name that is a Go keyword",
			Option:   "example.com/func",
			WantPath: "example.com/func",
			WantName: "_func",
		},
		{
			Name:     "explicit package name",
			Option:   "example.com/my-module;foo",
			WantPath: "example.com/my-module",
			WantName: "foo",
		},
		{
			Name:    "explicit package name containing a hyphen",
			Option:  "example.com/foo;my-pkg",
			WantErr: `file.proto specifies an invalid Go package name in its 'go_package' option ("example.com/foo;my-pkg"): "my-pkg" is not a valid Go identifier`,
		},
		{
			Name:    "explicit package name that is a Go keyword",
			Option:  "example.com/foo;func",
			WantErr: `file.proto specifies an invalid Go package name in its 'go_package' option ("example.com/foo;func"): "func" is a Go keyword`,
		},
		{
			Name:    "missing option",
			Option:  "",
			WantErr: `file.proto does not specify a 'go_package' option`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			d := &descriptorpb.FileDescriptorProto{
				Name: proto.String("file.proto"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String(tc.Option),
				},
			}

			gotPath, gotName, err := GoPackage(d)

			if tc.WantErr != "" {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				if err.Error() != tc.WantErr {
					t.Fatalf("unexpected error:\nwant: %s\n got: %s", tc.WantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			if gotPath != tc.WantPath {
				t.Fatalf("unexpected package path: want %q, got %q", tc.WantPath, gotPath)
			}
			if gotName != tc.WantName {
				t.Fatalf("unexpected package name: want %q, got %q", tc.WantName, gotName)
			}
		})
	}
}
