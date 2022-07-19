package convert

type From[F, T any] interface {
	From(F) T
}

type Into[T any] interface {
	Into() T
}

type AsRef[T any] interface {
	AsRef(T) *T
}

type Converter[From any, To any] interface {
	Convert(From) To
}
