# generigo &nbsp;[![GoDoc](https://godoc.org/github.com/dlampsi/generigo?status.svg)](https://godoc.org/github.com/dlampsi/generigo) [![Build Status](https://travis-ci.org/dlampsi/gogen.svg?branch=master)](https://travis-ci.org/dlampsi/gogen)

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