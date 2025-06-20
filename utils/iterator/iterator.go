package iterator

// Iterator is an basic iterator
type ConstIterator[T any] interface {
	HasNext() bool
	Next() ConstIterator[T]
	Value() T
	Clone() ConstIterator[T]
	Equal(other ConstIterator[T]) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any] interface {
	ConstIterator[T]
	SetValue(value T)
}

// ConstKvIterator is an iterator that returns key-value pairs
type ConstKvIterator[K, V any] interface {
	ConstIterator[V]
	Key() K
}

// KvIterator is an interface of mutable key-value iterator
type KvIterator[K, V any] interface {
	ConstKvIterator[K, V]
	SetValue(value V)
}

// ConstBidIterator is an iterator that supports bidirectional traversal
type ConstBidIterator[T any] interface {
	ConstIterator[T]
	HasPrev() bool
	Prev() ConstBidIterator[T]
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator[T any] interface {
	ConstBidIterator[T]
	SetValue(value T)
}

// ConstKvBidIterator is an iterator that returns key-value pairs and supports bidirectional traversal
type ConstKvBidIterator[K, V any] interface {
	ConstKvIterator[K, V]
	BidIterator[V]
}

// KvBidIterator is an interface of mutable key-value iterator that supports bidirectional traversal
type KvBidIterator[K, V any] interface {
	KvIterator[K, V]
	BidIterator[V]
}

// RandomAccessIterator is an iterator that supports random access
type RandomAccessIterator[T any] interface {
	BidIterator[T]
	IteratorAt(index int) RandomAccessIterator[T]
	Index() int
}
