// Package result is an implementation of functional Results,
// loosely modeled on Rust's `Result`.
package result

import (
	"reflect"

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

// AndThen calls `fn` on `r` if the result is ok. Otherwise
// returns the original result.
func (r Result[T, E]) AndThen(fn func(data T) Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return fn(*r.ok)
	}

	return r
}

// Or returns the `res` if `r` is an error, otherwise returns `r`.
func (r Result[T, E]) Or(res Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return r
	}

	return res
}

// OrElse returns the `res` if `r` is an error, otherwise calls `fn` on the error.
func (r Result[T, E]) OrElse(fn func(e E) Result[T, E]) Result[T, E] {
	if r.IsErr() {
		return fn(*r.err)
	}

	return r
}

// Contains compares the wrapped data to the data
// parameter.
func (r Result[T, E]) Contains(data T) bool {
	if r.IsOk() {
		// Without further constraining T,
		// it may not be possible to compare
		// without reflection. Constraining T
		// would severly hinder the API.
		if reflect.DeepEqual(*r.ok, data) {
			return true
		}
	}

	return false
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

// MapOr returns the default if error, or applies the `fn` to
// to the wrapped value.
func (r Result[T, E]) MapOr(def T, fn func(data T) T) T {
	if r.IsOk() {
		return fn(*r.ok)
	}

	return def
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

// IsOkAnd returns true if the Result is ok and the predicate
// returns true.
func (r Result[T, E]) IsOkAnd(fn func(data T) bool) bool {
	if r.IsOk() {
		return fn(*r.ok)
	}

	return false
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

// ExpectErr is an assertion that the operation was error
// that returns the underlying error. If not, ExpectErr
// panics with `msg`. Only use this if you intend for your
// program to crash on error or if you `recover`.
func (r Result[T, E]) ExpectErr(msg string) E {
	if !r.IsErr() {
		panic(msg)
	}

	return *r.err
}

// Unwrap returns the underlying data. If the Result is an error,
// Unwrap panics. Only use if you intend for your program to crash
// or if you `recover`.
func (r Result[T, E]) Unwrap() T {
	return r.Expect("called Unwrap an on an error")
}

// UnwrapErr returns the underlying error. If the Result is ok,
// UnwrapErr panics. Only use if you intend for your program to crash
// or if you `recover`.
func (r Result[T, E]) UnwrapErr() E {
	return r.ExpectErr("called UnwrapErr an ok")
}

func Ok[T any](data T) Result[T, error] {
	return Result[T, error]{ok: &data}
}

func Err[T any, E error](e E) Result[T, E] {
	return Result[T, E]{err: &e}
}

// Match accepts data and an error (the return from an ioutil.ReadAll, for example),
// matches on the values, and returns the appropriate result.
func Match[T any](data T, e error) Result[T, error] {
	if e != nil {
		return Err[T](e)
	}

	return Ok(data)
}
