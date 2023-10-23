# Option

Option is a small Go library, that provides single struct - `Option`,
that can simply have a value, or have no value (technically it will
always have a value, but not always will allow you to access the value).

The `Option` can be created via two functions: `Some` and `None`, however
`None` is not going to be used as frequently, because the nil-value of
`Option` is already a 'none'.

After having an `Option` in your code, make sure to check it via `(Option[T]).HasValue()`
method, which reports whether there is a value or not, otherwise if the value is not present
the `(Option[T]).Value()` method will rise a runtime panic.

## How to use

Add the library as a dependency to your module:
```bash
$ go get github.com/kerelape/option
```
