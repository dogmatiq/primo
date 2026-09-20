package formatting

// AsString returns a custom string representation used by the generated
// Format() method for the %s and %q verbs.
func (x *Stringable) AsString() string {
	return "as-string(" + x.GetValue() + ")"
}
