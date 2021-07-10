# generigo 

[![build-img]][build-url]
[![doc-img]][doc-url]
[![coverage-img]][coverage-url]

Simple common methods for Go

## Examples

```go
import "github.com/dlampsi/generigo"
```

String in slice

```go
str := "apple"
slice := []string{"pineapple", "apple", "onion"}

if generigo.StringInSlice(str, slice) {
    // Proccess case
}
```

Compare slices

```go
slice1 := []string{"one", "two"}
slice2 := []string{"two", "one"}

if generigo.CompareStringSlices(slice1, slice2) {
    // Proccess case
}
```

[build-img]: https://github.com/dlampsi/generigo/workflows/build/badge.svg
[build-url]: https://github.com/dlampsi/generigo/actions
[coverage-img]: https://codecov.io/gh/dlampsi/generigo/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/dlampsi/generigo
[doc-img]: https://pkg.go.dev/badge/dlampsi/generigo
[doc-url]: https://pkg.go.dev/github.com/dlampsi/generigo