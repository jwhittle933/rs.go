// Package option is an implementation of functional Options,
// loosely modeled on Rust's `Option`.
package option

type Option[T any] struct {
	some *T
}

func (o Option[T]) And(other Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}

	return other
}

func (o Option[T]) AndThen(fn func(data T) Option[T]) Option[T] {
	if o.IsSome() {
		return fn(*o.some)
	}

	return o
}

func (o Option[T]) IsSome() bool {
	return o.some != nil
}

func (o Option[T]) IsNone() bool {
	return o.some == nil
}

func (o Option[T]) Expect(msg string) T {
	if o.IsNone() {
		panic(msg)
	}

	return *o.some
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		o.Expect("unwrapped a none")
	}

	return *o.some
}

func Some[T any](data T) Option[T] {
	return Option[T]{some: &data}
}

func None[T any]() Option[T] {
	return Option[T]{}
}
