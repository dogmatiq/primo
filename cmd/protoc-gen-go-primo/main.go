package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dogmatiq/primo/internal/generator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var version string // populated automatically

func main() {
	if err := generate(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// generate reads a code generation request from r, invokes gen() and writes its
// response to w.
func generate(
	r io.Reader,
	w io.Writer,
) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("unable to read request: %w", err)
	}

	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		return fmt.Errorf("unable to unmarshal request: %w", err)
	}

	if len(req.FileToGenerate) == 0 {
		return fmt.Errorf("no files to generate")
	}

	res := generator.Generate(version, req)

	data, err = proto.Marshal(res)
	if err != nil {
		return fmt.Errorf("unable to marshal response: %w", err)
	}

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("unable to write response: %w", err)
	}

	return nil
}
