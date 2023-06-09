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
	"errors"
	"testing"

	option "github.com/kerelape/option/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNone_Present(t *testing.T) {
	assert.Equal(
		t,
		false,
		option.NewNone[any]().Present(),
		"Present is expected to always return false.",
	)
}

func TestNone_Value(t *testing.T) {

	t.Run("without reason", func(t *testing.T) {
		assert.Panics(
			t,
			func() {
				option.NewNone[any]().Value()
			},
			"Value is expected to panic.",
		)
	})
	t.Run("with reason", func(t *testing.T) {
		assert.PanicsWithError(
			t,
			"because I want so",
			func() {
				option.NewNoneReason[any](
					errors.New("because I want so"),
				).Value()
			},
			"Value must panic with the error passed as reason",
		)
	})
}
