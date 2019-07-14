# gogen &nbsp;[![GoDoc](https://godoc.org/github.com/dlampsi/gogen?status.svg)](https://godoc.org/github.com/dlampsi/gogen)

Simple common methods for Go

## Examples

```go
import "github.com/dlampsi/gogen"
```

String in slice

```go
str := "apple"
slice := []string{"pineapple", "apple", "onion"}

if gogen.StringInSlice(str, slice) {
    // Proccess case
}
```

Compare slices

```go
slice1 := []string{"one", "two"}
slice2 := []string{"two", "one"}

if gogen.CompareStringSlices(slice1, slice2) {
    // Proccess case
}
```