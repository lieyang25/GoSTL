package vector

import "fmt"

// select Options for vector
type Options struct {
	InitialCapacity int
	GrowFactor      float64
}

// OptionFuncs is a function that modifies the Options struct
type OptionFuncs func(*Options)

func NewOptions(opts ...OptionFuncs) *Options {
	options := &Options{
		InitialCapacity: 0,
		GrowFactor:      1.25,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// check options for vector
func (o *Options) Validate() error {
	if o.InitialCapacity < 0 {
		return fmt.Errorf("initial cap cannot be negative : %d", o.InitialCapacity)
	}
	if o.GrowFactor <= 1.0 {
		return fmt.Errorf("grow factor must be greater than 1 : %f", o.GrowFactor)
	}
	return nil
}

// set capacity for vector
func WithInitialCapacity(capacity int) OptionFuncs {
	return func(opts *Options) {
		opts.InitialCapacity = capacity
	}
}

//set growfactor for vector
func WithGrowFactor(factor float64) OptionFuncs {
	return func(opts *Options) {
		opts.GrowFactor = factor
	}
}

type Vector[T any] struct {
	data []T
}

// New create a new vector with options
func New[T any](opts ...OptionFuncs) *Vector[T] {
	options := NewOptions(opts...)

	if err := options.Validate(); err != nil {
		panic(fmt.Sprintf("invalid options: %v", err))
	}

	return &Vector[T]{
		data: make([]T, 0, options.InitialCapacity),
	}
}

// create a new vector from another vector
func NewFromVector[T any](other *Vector[T]) *Vector[T] {
	if other == nil {
		return New[T]()
	}
	v := &Vector[T]{
		data: make([]T, len(other.data), cap(other.data)),
	}
	copy(v.data, other.data)
	return v
}

// push item back to vector
func (v *Vector[T]) PushBack(item T) {
	v.data = append(v.data, item)
}

// push item front to vector
func (v *Vector[T]) PushFront(item T) {
	v.data = append([]T{item}, v.data...)
}

// pop item back from vector
func (v *Vector[T]) PoPBack() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	item := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return item, true
}

// pop item front from vector
func (v *Vector[T]) PoPFront() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	item := v.data[0]
	v.data = v.data[1:]
	return item, true
}

// insert item at specific index
func (v *Vector[T]) Inesert(index int, item T) error {
	if index < 0 || index > len(v.data) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	v.data = append(v.data[:index], append([]T{item}, v.data[index:]...)...)
	return nil
}

// insert a range of items at specific index
func (v *Vector[T]) InesertRange(index int, items []T) error {
	if index < 0 || index >= len(v.data) {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	if len(items) == 0 {
		return nil // nothing to insert
	}
	v.data = append(v.data[:index], append(items, v.data[index:]...)...)
	return nil
}

// erase item at specific index
func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index > len(v.data) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	v.data = append(v.data[:index], v.data[index+1:]...)
	return nil
}

// erase a range of items from the vector
func (v *Vector[T]) EraseRange(start, end int) error {
	if start < 0 || end > len(v.data) || start >= end {
		return fmt.Errorf("invalid range: %d to %d", start, end)
	}
	v.data = append(v.data[:start], v.data[end:]...)
	return nil
}

// makes a new space for the vector with passed capacity
func (v *Vector[T]) Reserve(capacity int) {
	if cap(v.data) >= capacity {
		return
	}
	data := make([]T, v.Size(), capacity)
	for i := 0; i < len(v.data); i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

// shrinks the capacity of the vector to the fit size
func (v *Vector[T]) ShrinkToFit() {
	if len(v.data) == cap(v.data) {
		return
	}
	len := v.Size()
	data := make([]T, len, len)
	for i := 0; i < len; i++ {
		data[i] = v.data[i]
	}
	v.data = data
}

// search item in vector
func (v *Vector[T]) At(index int) (T, error) {
	if index < 0 || index >= len(v.data) {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", index)
	}
	return v.data[index], nil
}

// set item at specific index
func (v *Vector[T]) Set(index int, item T) error {
	if index < 0 || index >= len(v.data) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	v.data[index] = item
	return nil
}

// search item front
func (v *Vector[T]) Front() (T, error) {
	if len(v.data) == 0 {
		var zero T
		return zero, fmt.Errorf("vector is empty")
	}
	return v.data[0], nil
}

// serach item back
func (v *Vector[T]) Back() (T, error) {
	if len(v.data) == 0 {
		var zero T
		return zero, fmt.Errorf("vector is empty")
	}
	return v.data[len(v.data)-1], nil
}

// Data returns the underlying slice of the vector
func (v *Vector[T]) Data() []T {
	return v.data
}

// return size
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// return cap
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

//return true if vector is empty
func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

// clear the vector
func (v *Vector[T]) Clear() {
	v.data = make([]T, 0, cap(v.data))
}
