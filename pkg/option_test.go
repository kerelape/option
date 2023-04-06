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

package option_test

import (
	"math/rand"
	"testing"

	option "github.com/kerelape/option/pkg"
)

// TestOption_NoneIsEmpty tests that None returns an empty Option.
func TestOption_NoneIsEmpty(t *testing.T) {
	subject := option.None[any]()
	if subject.NotEmpty() {
		t.Error("Expected None to be empty")
	}
}

// TestOption_SomeIsNotEmpty tests that Some returns not an empty Option.
func TestOption_SomeIsNotEmpty(t *testing.T) {
	v := 0
	subject := option.Some(&v)
	if subject.Empty() {
		t.Error("Expected Some to have a value")
	}
}

// TestOption_ReturnsCorrectValue tests that an Option with value returns it correctly.
func TestOption_ReturnsCorrectValue(t *testing.T) {
	want := rand.Float64()
	subject := option.Some(&want)
	actual := subject.Value()
	if actual != want {
		t.Errorf("Expected %f, got %f", want, actual)
	}
}
