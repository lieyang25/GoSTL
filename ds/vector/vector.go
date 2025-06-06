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
		GrowFactor:      1.5,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

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

func (v *Vector[T]) PushBack(item T) {
	v.data = append(v.data, item)
}

func (v *Vector[T]) PushFront(item T) {
	v.data = append([]T{item}, v.data...)
}

func (v *Vector[T]) PoPBack() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	item := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return item, true
}

func (v *Vector[T]) PoPFront() (T, bool) {
	if len(v.data) == 0 {
		var zero T
		return zero, false
	}
	item := v.data[0]
	v.data = v.data[1:]
	return item, true
}

func (v *Vector[T]) Size() int {
	return len(v.data)
}

func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}
