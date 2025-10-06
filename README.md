# Go (Golang) — 5-minute cheat-sheet for Node.js developers

This short README is designed to get a Node.js developer up to speed with Go syntax, concepts, and common patterns in about 5 minutes. It focuses on differences, quick examples, and things to remember.

## Quick setup

- Install Go: https://go.dev/dl
- Verify:

```bash
go version
```

- Create a module (project) in your folder:

```bash
go mod init example.com/quick
```

- Build and run:

```bash
go run main.go
# or
go build -o app .
./app
```

## File layout

- Single-file program: `main.go` with `package main` and `func main()` as the entrypoint.
- Modules: `go.mod` tracks module path and dependency versions.

## Hello world

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}
```

## Basic differences from Node.js / JavaScript

- Statically typed, compiled language. Types are explicit (recommended) or inferred with `:=`.
- No `this` or prototype-based OO. Use structs and methods.
- No exceptions like JS `throw`/`try` (there is `panic` but avoid); use multi-value returns for errors.
- Concurrency via goroutines and channels, not promises.

## Variables and types

```go
var x int = 10
var s string = "hello"
// type inference
y := 20 // compiler infers int

// multiple assignment
a, b := 1, "two"

// zero values (uninitialized)
var z int    // 0
var flag bool // false
```

Common types: `int`, `int64`, `float64`, `string`, `bool`, `complex128`, `byte` (alias for uint8), `rune` (alias for int32, a Unicode code point)

## Functions

```go
func add(a int, b int) int {
    return a + b
}

// multiple returns (used for errors)
func div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }
    return a / b, nil
}
```

Callsite:

```go
res, err := div(10, 2)
if err != nil {
    // handle
}
```

## Structs & methods (like classes)

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string { // value receiver
    return "hi, " + p.Name
}

func (p *Person) SetAge(a int) { // pointer receiver
    p.Age = a
}
```

Note: methods can have pointer receivers to mutate state.

## Slices, arrays, maps

```go
// array (fixed size)
var arr [3]int

// slice (dynamic)
s := []int{1,2,3}
s = append(s, 4)

// map
m := map[string]int{"a": 1}
val := m["a"]

// check presence
v, ok := m["b"]
if !ok {
    // not present
}
```

## Control flow

```go
if x > 0 {
    // ...
} else if x == 0 {
    // ...
} else {
    // ...
}

for i := 0; i < 5; i++ {
    // like JS for
}

// like while
for i < 10 {
    i++
}

// iterate over slice or map
for idx, val := range s {
    // idx, val
}
```

## Error handling

- Idiomatic Go returns (value, error). Always check errors.

```go
res, err := doSomething()
if err != nil {
    // handle or return
}
```

- Avoid `panic` except for unrecoverable errors.

## Concurrency: goroutines and channels

```go
// goroutine (fire-and-forget)
go doWork()

// channels
ch := make(chan int)

// sender
go func() { ch <- 42 }()

// receiver
v := <-ch

// buffered channel
buf := make(chan string, 2)
```

Sync with WaitGroup (from `sync`):

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

## Package imports

- Standard library is rich: `fmt`, `net/http`, `io`, `os`, `encoding/json`, `sync`, `context`.
- To import third-party use module path with `go get` or just `go build`/`run` (Go will fetch):

```bash
go get github.com/some/repo
```

## JSON & HTTP quick example

```go
// simple http server
package main

import (
    "encoding/json"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"hello": "world"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Tooling

- go fmt: formats code. Run before commits.

```bash
go fmt ./...
```

- go vet: static analysis for suspicious code.

```bash
go vet ./...
```

- go test: run tests.

```bash
go test ./...
```

- go mod tidy: clean unused deps.

```bash
go mod tidy
```

## Quick JS -> Go mapping (mental model)

- JS `async/await` -> goroutines + channels or `sync` primitives
- JS objects -> Go `struct`
- JS arrays -> Go `[]T` (slice)
- JS Map -> Go `map[K]V`
- JS functions as first-class -> same in Go; use function types
- `null`/`undefined` -> zero values and `nil` for pointers, slices, maps, channels, interfaces

## Performance & compilation

- Go compiles to a single binary. Fast startup compared to Node.
- Cross-compile with `GOOS`/`GOARCH` env vars.

## Gotchas & best practices

- Always check errors.
- Use `go fmt`.
- Prefer explicit types on exported APIs.
- Use pointers when you want to mutate or avoid copies for large structs.
- Avoid goroutine leaks: cancel via `context.Context` or use done channels.
- `defer` runs when function returns—use for cleanup (close files, unlock mutexes).

## Small examples (copy/paste)

1) multiple return (error handling)

```go
package main

import (
    "fmt"
)

func safeDiv(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("b is zero")
    }
    return a / b, nil
}

func main() {
    r, err := safeDiv(10, 0)
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    fmt.Println(r)
}
```

2) goroutine + channel

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch <- "done"
    }()
    fmt.Println(<-ch)
}
```

## Learning path (next steps)

- Read the tour: https://go.dev/tour
- Official docs: https://go.dev/doc/
- Concurrency patterns: Go blog (rob pike), Effective Go

## Closing: quick checklist

- [ ] Install Go and run `go version`
- [ ] Create `go.mod` with `go mod init` in project
- [ ] Try `go run main.go` with examples above
- [ ] Learn `defer`, `error` handling, and goroutines

Happy Go-coding! 🎯
