// Package problem provides a problem details implementation.
//
// See the specification at https://www.rfc-editor.org/rfc/rfc9457.html
//
// The details type defined by this package is made to be robust and simple.
// It provides a very rudimentary extension mechanism that serves as a base
// for more sophisticated solutions.
//
// Currently, only JSON is supported. XML support may be added in the future.
//
// Both "type" and "instance" are defined to be URIs. Checking their correctness,
// i.e. that they really are URIs, is not done by this package.
package problem
