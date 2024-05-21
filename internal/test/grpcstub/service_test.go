package grpcstub_test

import (
	context "context"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/grpcstub"
	grpc "google.golang.org/grpc"
)

func TestClientStub(t *testing.T) {
	t.Parallel()

	stub := &ServiceClientStub{}

	t.Run("when the implementation function is nil", func(t *testing.T) {
		t.Run("it returns an error", func(t *testing.T) {
			_, err := stub.Unary(context.Background(), &Request{})
			if err == nil {
				t.Fatal("expected an error")
			}

			want := "rpc error: code = Unimplemented desc = method Unary not implemented"
			if err.Error() != want {
				t.Fatalf("unexpected error: got %q, want %q", err, want)
			}
		})
	})

	t.Run("when the implementation function is non-nil", func(t *testing.T) {
		t.Run("it returns the result of the implementation function", func(t *testing.T) {
			wantReq := &Request{}
			wantRes := &Response{}

			stub.UnaryFunc = func(
				_ context.Context,
				gotReq *Request,
				_ ...grpc.CallOption,
			) (*Response, error) {
				if gotReq != wantReq {
					t.Fatalf("unexpected request: got %v, want %v", gotReq, wantReq)
				}
				return wantRes, nil
			}

			gotRes, err := stub.Unary(context.Background(), wantReq)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotRes != wantRes {
				t.Fatalf("unexpected response: got %v, want %v", gotRes, wantRes)
			}
		})
	})
}
