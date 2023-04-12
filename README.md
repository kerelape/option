<!--
MIT License

Copyright (c) 2023 kerelape

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
-->
# Option

Option is an object-oriented library for optional value in Go.

[![EO principles respected here](https://www.elegantobjects.org/badge.svg)](https://www.elegantobjects.org)

[![Build](https://github.com/kerelape/option/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/kerelape/option/actions/workflows/build.yml)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/kerelape/option/blob/main/LICENSE.txt)

## How to use this library

Add it to your Go module:
```bash
$ go get -u gihub.com/kerelape/option
```

Create an optional with some value:
```go
package main

import option "github.com/kerelape/pkg"

func main() {
  message := option.NewSome("Hello, World!")
  println(message.Value())
}
```

Or an optional without a value:
```go
package main

import option "github.com/kerelape/pkg"

func main() {
  message := option.NewNone[string]()
  println(message.Value()) // will panic with option.ErrNoValue
  
  // Or with a reason
  message = option.NewNoneReason[string](errors.New("nothing to say"))
  println(message.Value()) // will panic with the "nothing to say" error
}
```

## How to Contribute

Fork this repository, make changes, send us a pull request. We will review your changes and apply them to the `main` branch shortly, provided they don't violate our quality standards. To avoid frustration, before sending us your pull request please run the tests:

```bash
$ go test -v ./...
```

You will need Go 1.20+.
