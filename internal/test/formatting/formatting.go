package formatting

// AsString returns a custom string representation used by the generated
// Format() method for the %s and %q verbs.
//
// It dereferences x directly so that it panics on a nil receiver, exercising
// the generated Format() method's nil guard.
func (x *Stringable) AsString() string {
	return "as-string(" + x.Value + ")"
}
