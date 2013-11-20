# codederror

A small utility that provdes a way to encode errors with namespaces and integers.

## Usage

```go
import "github.com/ngerakines/codederror"

func awesome() string, err {
  // ...
  return nil, NewCodedError([]string{""}, 24)
  // ...
}
```
