**DAY 1 — Go Fundamentals**

* Go is a compiled language that produces native executables (binaries).
* Compilation time checks code correctness; runtime is when the program executes.
* A binary/executable runs directly on the OS without an interpreter.
* `go run` compiles and runs immediately (no saved binary).
* `go build` compiles and produces a reusable executable.
* A Go program must use `package main` to be executable.
* `func main()` is the single entry point; it takes no arguments and returns nothing.
* Programs communicate via side effects (printing, files, network), not return values.
* Go code is organized into packages; all `.go` files in a directory form one package.
* Variables are statically typed; types are fixed at compile time.
* `var` declares explicitly; `:=` declares + assigns inside functions only.
* Basic types learned: `int`, `string`, `bool`.
* `if` requires a boolean condition; no truthy/falsy values.
* `for` is the only loop construct (classic, while-style, infinite).
* `for` init/condition/post each allow exactly one expression.
* Go enforces formatting via `gofmt` / `go fmt`.
* `gofmt` enforces structure, not personal spacing preferences.
* A Go module is a directory with a `go.mod` file.
* `go mod init` is run once per project; optional for throwaway snippets.

---

**Day 2 — Core Go Building Blocks**

* Functions can be defined outside `main`; `main` should only orchestrate.
* Functions can return multiple values.
* Errors are ordinary return values; `nil` means no error.
* A group of values of the same type can be created using `[]T{...}`.
* `len(x)` returns how many values are in a group.
* Repetition over a group uses `for _, v := range group`.
* `_` is used to ignore a value you don’t need.
* Functions should return early when an error occurs.
* Empty input must be handled explicitly.
* Multiple results (count, sum, min, max) can be computed in a single loop.
* Logic belongs in functions; printing belongs in `main`.

---

**Day 3 — Structs, Pointers, and Slice Mutation**

* A `struct` type defines a data shape; a struct value allocates memory and holds concrete data.
* Struct values are created using **struct literals** (e.g. `Book{}`, `Book{Field: value}`).
* A **literal** is a value written directly in code; struct literals are composite literals.
* Passing a struct to a function copies the entire value.
* Passing a pointer to a struct (`*T`) allows the function to modify the original value.
* `&x` produces the address of a value; `*T` is a type that can store such an address.
* Go automatically dereferences struct pointers for field access (`p.Field` works for `*p`).
* `for ... range` assigns values by position, not variable name.
* In `for i, v := range slice`, `i` is the index and `v` is a copy of the element.
* Modifying the `range` value (`v`) does not modify the slice.
* To mutate slice elements, iterate by index: `for i := range slice { slice[i] = ... }`.
* A slice value is a small descriptor (pointer, length, capacity) that refers to an underlying array.
* Passing a slice to a function copies the descriptor, but both copies refer to the same array.
* Modifying slice elements via `slice[i]` affects the original data without using pointers.
* A pointer to a slice (`*[]T`) is only needed to change the slice itself (length/capacity), not its elements.
* Idiomatic Go prefers `[]T` over `*[]T` unless structural slice changes are required.

---



