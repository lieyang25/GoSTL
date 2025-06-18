package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next() Iterator[T]
	Value() T
	Clone() Iterator[T]
	Equal(other Iterator[T]) bool
}
