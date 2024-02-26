package fp

type _cons[T any] struct {
	a T
	b *_cons[T]
}

func cons[T any](a T, b *_cons[T]) *_cons[T] {
	return &_cons[T]{
		a, b,
	}
}
