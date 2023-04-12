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

import (
	"errors"
)

var ErrNoValue = errors.New("no value present")

// None is an Option without a value.
type None[T any] struct {
	reason error // reason why the option has no value.
}

// NewNoneReason creates a new None with reason.
func NewNoneReason[T any](reason error) None[T] {
	return None[T]{reason}
}

// NewNone creates a new None without any reason.
func NewNone[T any]() None[T] {
	return NewNoneReason[T](nil)
}

func (n None[T]) Present() bool {
	return false
}
func (n None[T]) Value() T {
	dummy, _ := any(nil).(T)
	if n.reason != nil {
		panic(n.reason)
	} else {
		panic(ErrNoValue)
	}
	return dummy
}
