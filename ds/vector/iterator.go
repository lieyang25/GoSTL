package vector

import (
	"github.com/lieyang25/GoSTL/utils/iterator"
)

type T any

var _ iterator.RandomAccessIterator[T] = (*VectorIterator[T])(nil)

type VectorIterator[T any] struct {
	vec   *Vector[T]
	index int
}

func (iter *VectorIterator[T]) Clone() iterator.ConstIterator[T] {
	return &VectorIterator[T]{vec: iter.vec, index: iter.index}
}

func (iter *VectorIterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	if otherIter, ok := other.(*VectorIterator[T]); ok {
		return iter.vec == otherIter.vec && iter.index == otherIter.index
	}
	return false
}

func (iter *VectorIterator[T]) HasNext() bool {
	return iter.index < len(iter.vec.data)
}

func (iter *VectorIterator[T]) Next() iterator.ConstIterator[T] {
	if iter.HasNext() {
		iter.index++
		return iter
	}
	return nil
}

func (iter *VectorIterator[T]) HasPrev() bool {
	return iter.index > 0
}

func (iter *VectorIterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.HasPrev() {
		iter.index--
		return iter
	}
	return nil
}

func (iter *VectorIterator[T]) SetValue(value T) {
	if iter.index < 0 || iter.index >= len(iter.vec.data) {
		return // out of range, do nothing
	}
	iter.vec.data[iter.index] = value
}

func (iter *VectorIterator[T]) Value() T {
	if iter.index < 0 || iter.index >= len(iter.vec.data) {
		var zero T
		return zero
	}
	return iter.vec.data[iter.index]
}

func (iter *VectorIterator[T]) Index() int {
	return iter.index
}

func (iter *VectorIterator[T]) IteratorAt(index int) iterator.RandomAccessIterator[T] {
	if index < 0 || index >= len(iter.vec.data) {
		return nil // out of range, return nil
	}
	return &VectorIterator[T]{vec: iter.vec, index: index}
}
