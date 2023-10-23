// Package option provides struct with optional value.
package option

// Option is an optional value.
type Option[T any] struct {
	value T
	has   bool
}

// Some returns an option with the value.
func Some[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		has:   true,
	}
}

// None returns an empty option.
func None[T any]() Option[T] {
	return Option[T]{}
}

// HasValue reports whether the option has value or not.
func (o Option[T]) HasValue() bool {
	return o.has
}

// Value returns the value or panics if no value is present.
func (o Option[T]) Value() T {
	if !o.has {
		panic("option has not value")
	}
	return o.value
}
