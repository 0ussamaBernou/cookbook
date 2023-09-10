# Go's basic types

```go
bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
// represents a Unicode code point

float32 float64

complex64 complex128
```

# Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.

# Type inference

if a variable declaration is accompanied with initializer the type is infered
but for numeric types it depends on the precision

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```
