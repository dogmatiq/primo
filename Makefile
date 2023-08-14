# Note, these files depend on the protoc-gen-go-primo binary from this repo to
# be built.
#
# They have been added to .gitignore so that they are excluded from the
# GO_SOURCE_FILES variable, as otherwise it would create a circular dependency.
GO_TEST_REQ += internal/test/oneof/oneof_primo.pb.go

-include .makefiles/Makefile
-include .makefiles/pkg/protobuf/v2/Makefile
-include .makefiles/pkg/go/v1/Makefile

%_primo.pb.go: %.proto $(PROTOC_COMMAND) artifacts/protobuf/bin/go.mod artifacts/protobuf/args/common artifacts/protobuf/args/go $(GO_DEBUG_DIR)/protoc-gen-go-primo
	PATH="$(GO_DEBUG_DIR):$(MF_PROJECT_ROOT)/artifacts/protobuf/bin:$$PATH" $(PROTOC_COMMAND) \
		--proto_path="$(dir $(PROTOC_COMMAND))../include" \
		--go-primo_opt=module=$$(go list -m) \
		--go-primo_out=. \
		$$(cat artifacts/protobuf/args/common artifacts/protobuf/args/go) \
		$(MF_PROJECT_ROOT)/$(@D)/*.proto

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"
