// Package result is an implementation of functional Results,
// loosely modeled on Rust's `Result`.
package result

import (
	"github.com/jwhittle933/rs.go/option"
)

// Result represents an operation that can succeed or fail.
// It wraps either an `ok` operation or an `error` operation.
// Becuase `Result` methods return `Result` interfaces, you can chain
// your method calls together and "happy path" a procedural chain without
// checking for an error until the end of the procedure.
type Result[T any, E any] struct {
	ok  *T
	err *E
}

// And returns `r` if the result is `ok`. Otherwise
// returns the original result.
func (r Result[T, E]) And(res Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return res
	}

	return r
}

// Map calls `m` on the underlying data of
// `Result`. The return from `m` is wrapped and returned. In
// the event of an error, `m` is not called and the
// error Result is returned unchanged. Go's generics
// don't allow for new parameter introduction in an interface,
// so Map can operate only on `T`.
func (r Result[T, E]) Map(fn func(data T) T) Result[T, E] {
	if r.IsOk() {
		op := fn(*r.ok)
		return Result[T, E]{ok: &op}
	}

	return r
}

// MapErr calls `m` on the underlying error of
// `Result`. The return from `m` is wrapped and returned. In
// the event of an error, `m` is not called and the
// error Result is returned unchanged. Go's generics
// don't allow for new parameter introduction in an interface,
// so Map can operate only on `T` or `E`.
func (r Result[T, E]) MapErr(fn func(e E) E) Result[T, E] {
	if r.IsErr() {
		op := fn(*r.err)
		return Result[T, E]{err: &op}
	}

	return r
}

// Ok returns the underlying data wrapped in Option[T].
// If the Result is an error, an None is returned.
func (r Result[T, E]) Ok() option.Option[T] {
	if r.IsOk() {
		return option.Some(*r.ok)
	}

	return option.None[T]()
}

// IsOk reports whether the Result is ok.
func (r Result[T, E]) IsOk() bool {
	return r.ok != nil && r.err == nil
}

// IsErr reports whether the Result is an error.
func (r Result[T, E]) IsErr() bool {
	return r.ok == nil && r.err != nil
}

// Err returns the underlying error wrapped in an Option[E].
// If the Result is ok, Err returns nil.
func (r Result[T, E]) Err() option.Option[E] {
	if r.IsErr() {
		return option.Some(*r.err)
	}

	return option.None[E]()
}

// Expect is an assertion that the operation was ok that
// returns the underlying data. If not, Expect panics
// with `msg`. Only use this if you intend for your
// program to crash on error or if you `recover`.
func (r Result[T, E]) Expect(msg string) T {
	if r.IsOk() {
		return *r.ok
	}

	panic(msg)
}

// Unwrap returns the underlying data. If the Result is an error,
// Unwrap panics. Only use if you intend for your program to crash
// or if you `recover`.
func (r Result[T, E]) Unwrap() T {
	return r.Expect("attempted to unwrap an error")
}

func Ok[T any](data T) Result[T, error] {
	return Result[T, error]{ok: &data}
}

func Err[T any, E error](e E) Result[T, E] {
	return Result[T, E]{err: &e}
}
