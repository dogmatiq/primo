// Package accessor generates an accessor method for each type in "one-of
// groups". While "protoc-gen-go" plugin already generates accessor methods for
// for each type in "one-of groups", the generated methods lack the ability to
// differentiate between the absence of a value and the presence of a zero
// value.
package accessor
