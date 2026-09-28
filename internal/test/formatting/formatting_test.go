package formatting_test

import (
	"fmt"
	"testing"

	. "github.com/dogmatiq/primo/internal/test/formatting"
)

var (
	_ fmt.Formatter = (*Stringable)(nil)
	_ fmt.Formatter = (*Plain)(nil)
	_ fmt.Formatter = (*Panicky)(nil)
	_ fmt.Formatter = (*NilSafe)(nil)
)

func TestFormat(t *testing.T) {
	cases := []struct {
		name   string
		value  any
		format string
		want   string
	}{
		{
			"uses AsString() for %s",
			&Stringable{Value: "hello"},
			"%s",
			"as-string(hello)",
		},
		{
			"uses AsString() for %q",
			&Stringable{Value: "hello"},
			"%q",
			`"as-string(hello)"`,
		},
		{
			"uses String() for %v even though AsString() is implemented",
			&Stringable{Value: "hello"},
			"%v",
			`value:"hello"`,
		},
		{
			"uses GoString() for %#v",
			&Stringable{Value: "hello"},
			"%#v",
			"go-string(hello)",
		},
		{
			"recovers from AsString() panicking on a nil receiver for %s",
			(*Stringable)(nil),
			"%s",
			"<nil>",
		},
		{
			"recovers from AsString() panicking on a nil receiver for %q",
			(*Stringable)(nil),
			"%q",
			"<nil>",
		},
		{
			"recovers from GoString() panicking on a nil receiver for %#v",
			(*Stringable)(nil),
			"%#v",
			"<nil>",
		},
		{
			"calls a nil-tolerant AsString() for %s",
			(*NilSafe)(nil),
			"%s",
			"as-string(nil)",
		},
		{
			"calls a nil-tolerant AsString() for %q",
			(*NilSafe)(nil),
			"%q",
			`"as-string(nil)"`,
		},
		{
			"calls a nil-tolerant GoString() for %#v",
			(*NilSafe)(nil),
			"%#v",
			"go-string(nil)",
		},
		{
			"formats an AsString() panic like fmt formats a panicking Stringer for %s",
			&Panicky{Value: "hello"},
			"%s",
			"%!s(PANIC=AsString method: boom)",
		},
		{
			"formats an AsString() panic like fmt formats a panicking Stringer for %q",
			&Panicky{Value: "hello"},
			"%q",
			"%!q(PANIC=AsString method: boom)",
		},
		{
			"formats a GoString() panic like fmt formats a panicking GoStringer for %#v",
			&Panicky{Value: "hello"},
			"%#v",
			"%!v(PANIC=GoString method: boom)",
		},
		{
			"delegates to String() for %v when AsString() is absent",
			&Plain{Value: "hello"},
			"%v",
			`value:"hello"`,
		},
		{
			"delegates to String() for %+v when AsString() is absent",
			&Plain{Value: "hello"},
			"%+v",
			`value:"hello"`,
		},
		{
			"delegates to String() for %s when AsString() is absent",
			&Plain{Value: "hello"},
			"%s",
			`value:"hello"`,
		},
		{
			"delegates to String() for %x when AsString() is absent",
			&Plain{Value: "hello"},
			"%x",
			"76616c75653a2268656c6c6f22",
		},
		{
			"delegates to String() for %X when AsString() is absent",
			&Plain{Value: "hello"},
			"%X",
			"76616C75653A2268656C6C6F22",
		},
		{
			"delegates to String() for %q when AsString() is absent",
			&Plain{Value: "hello"},
			"%q",
			`"value:\"hello\""`,
		},
		{
			"delegates to String() for %#q when AsString() is absent",
			&Plain{Value: "hello"},
			"%#q",
			"`value:\"hello\"`",
		},
		{
			"delegates to String() for %#x when AsString() is absent",
			&Plain{Value: "hello"},
			"%#x",
			"0x76616c75653a2268656c6c6f22",
		},
		{
			"delegates to String() for %#X when AsString() is absent",
			&Plain{Value: "hello"},
			"%#X",
			"0X76616C75653A2268656C6C6F22",
		},
		{
			"prefers AsString() over String() for %#q",
			&Stringable{Value: "hello"},
			"%#q",
			"`as-string(hello)`",
		},
		{
			"prefers String() over AsString() for %#x",
			&Stringable{Value: "hello"},
			"%#x",
			"0x76616c75653a2268656c6c6f22",
		},
		{
			"prefers String() over AsString() for %#X",
			&Stringable{Value: "hello"},
			"%#X",
			"0X76616C75653A2268656C6C6F22",
		},
		{
			"falls back to default formatting for %d",
			&Plain{Value: "hello"},
			"%d",
			"&{{{} [] [] 0} %!d(string=hello) [] 0}",
		},
		{
			"preserves the type name for %#v",
			&Plain{Value: "hello"},
			"%#v",
			`&formatting.Plain{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, Value:"hello", unknownFields:[]uint8(nil), sizeCache:0}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			if got := fmt.Sprintf(c.format, c.value); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
