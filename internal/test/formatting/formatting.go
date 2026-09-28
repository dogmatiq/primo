package formatting

// AsString returns a custom string representation used by the generated
// Format() method for the %s and %q verbs.
//
// It dereferences x directly, so it panics on a nil receiver, exercising the
// generated Format() method's panic-recovery behavior.
func (x *Stringable) AsString() string {
	return "as-string(" + x.Value + ")"
}

// GoString returns a custom Go-syntax representation used by the generated
// Format() method for the %#v verb.
//
// It dereferences x directly, so it panics on a nil receiver, exercising the
// generated Format() method's panic-recovery behavior.
func (x *Stringable) GoString() string {
	return "go-string(" + x.Value + ")"
}

// AsString always panics, exercising the panic marker produced by the
// generated Format() method for a panic unrelated to a nil receiver.
func (x *Panicky) AsString() string {
	panic("boom")
}

// GoString always panics, exercising the panic marker produced by the
// generated Format() method for a panic unrelated to a nil receiver.
func (x *Panicky) GoString() string {
	panic("boom")
}

// AsString handles a nil receiver itself, exercising the generated Format()
// method's ability to call through to a nil-tolerant implementation.
func (x *NilSafe) AsString() string {
	if x == nil {
		return "as-string(nil)"
	}
	return "as-string(" + x.Value + ")"
}

// GoString handles a nil receiver itself, exercising the generated Format()
// method's ability to call through to a nil-tolerant implementation.
func (x *NilSafe) GoString() string {
	if x == nil {
		return "go-string(nil)"
	}
	return "go-string(" + x.Value + ")"
}
