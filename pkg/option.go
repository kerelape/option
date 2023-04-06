// MIT License
//
// # Copyright (c) 2023 kerelape
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package option

// Option is an optional value.
type Option[T any] [1]*T

// None returns an empty Option.
func None[T any]() Option[T] {
	return Option[T]{nil}
}

// Some returns an Option with value.
func Some[T any](value *T) Option[T] {
	return Option[T]{value}
}

// Empty returns true if the Option is empty.
func (p Option[T]) Empty() bool {
	return p[0] == nil
}

// Empty returns true if the Option has a value.
func (p Option[T]) NotEmpty() bool {
	return !p.Empty()
}

// Value returns the value stored in the Option.
func (p Option[T]) Value() T {
	if p.Empty() {
		panic("option is empty")
	}
	return *p[0]
}
