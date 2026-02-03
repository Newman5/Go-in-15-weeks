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

**Day 4 — Maps and Zero Values**

* A map is Go’s built-in key–value data structure, written as `map[K]V`.
* Map keys must be **comparable types** (types that support `==` and `!=`).
* Common comparable key types include basic types, arrays, and structs whose fields are all comparable.
* Non-comparable types (slices, maps, functions) cannot be used as map keys.
* Declaring a map with `var m map[K]V` creates a **nil map** with no allocated storage.
* A map must be initialized with `make(map[K]V)` or `map[K]V{}` before writing to it.
* The zero value of a map is `nil`.
* Reading from a nil map is allowed; writing to a nil map causes a runtime panic.
* Writing to a map is done via assignment: `m[key] = value`.
* Assigning to a new key inserts it; assigning to an existing key overwrites its value.
* Maps have no order; there is no concept of position or append.
* `len(m)` returns the number of key–value pairs in the map (0 for empty or nil maps).
* Reading a missing key returns the **zero value** of the map’s value type.
* The “comma ok” pattern (`value, ok := m[key]`) is used to check whether a key exists.
* `ok == true` means the key exists; `ok == false` means it does not.
* Keys can be safely deleted using `delete(m, key)`.
* Deleting a missing key is safe and has no effect.
* Maps are **reference types**; passing a map to a function does not copy its contents.
* No pointer is required to modify a map inside a function.
* Map iteration is done using `for k, v := range m`, and iteration order is not guaranteed.
* Use maps for fast key-based lookup when order does not matter; use slices when order matters.

---

**Day 5 — Methods, Receivers, and Interfaces**

* A **method** is a function with a receiver: `func (r T) Method()`.
* Methods attach behavior to **types**, not to values or pointers.
* Any named type (not just structs) can have methods.
* Methods are used when behavior logically belongs to data.

* A **value receiver (`T`)** receives a copy of the value.
* Value receivers cannot mutate the original value.
* A **pointer receiver (`*T`)** receives the address of the value.
* Pointer receivers can mutate the original value.
* Pointer receivers should be used when a method mutates state or when copying would be expensive.
* If any method on a type uses a pointer receiver, pointer receivers should be used consistently.
* Go automatically takes the address when calling pointer-receiver methods on addressable values.

* Method sets are defined per type:
  * `T` has methods with value receivers.
  * `*T` has methods with both value and pointer receivers.
* Pointer-receiver methods cannot be called on non-addressable values (e.g. temporary literals).

* An **interface** defines behavior as a set of method signatures.
* Interfaces contain no data and no implementation.
* A type implements an interface **implicitly** by implementing its methods.
* There is no `implements` keyword in Go.
* Interfaces should be small and focused.
* Idiomatic rule: **accept interfaces, return concrete types**.

* Functions can take interfaces as parameters.
* Any value whose type satisfies the interface can be passed in.
* Interfaces are primarily used at **boundaries** (APIs, packages, testing).
* Concrete types are preferred inside packages.

* A single type can implement multiple interfaces.
* Interface implementation is independent per interface.
* The same concrete value can be wrapped by multiple interface values.

* An interface value consists of two parts: the concrete type and the concrete value.
* An interface is `nil` only if both its type and value are `nil`.
* An interface holding a nil concrete value is not `nil`.
* Calling a method on such an interface may cause a runtime panic.
* Practical rule: do not return nil concrete values as interfaces.

* Interfaces decouple code from concrete implementations.
* Interfaces enable substitution and testing with fake implementations.
* Interfaces define architectural boundaries without inheritance.

* Structs hold data.
* Methods define behavior.
* Pointer receivers control mutation.
* Interfaces define contracts.
* Functions operate on abstractions at boundaries.

---

**Day 6 — Concurrency Fundamentals (Goroutines & Channels)**

* **Concurrency** is about structuring a program to handle multiple tasks that overlap in time.
* **Parallelism** is about executing tasks simultaneously on multiple CPU cores.
* Concurrency does **not** require parallel hardware; parallelism is optional.

* Every Go program starts with exactly **one goroutine**: the **main goroutine**.
* The main goroutine runs `main()`.
* When `main()` returns, the **entire program exits immediately**.
* All other goroutines are terminated when the main goroutine exits.
* Goroutines do **not** keep a program alive on their own.

* A **goroutine** is a lightweight unit of execution managed by the Go runtime.
* New goroutines are started explicitly using the `go` keyword: `go f()`.
* Without `go`, a function runs in the current goroutine.
* Goroutines are cheap compared to OS threads and are scheduled by the Go runtime.

* **Blocking** means a goroutine pauses execution until a condition is met.
* Blocking affects **only the goroutine performing the operation**, not the whole program.
* Blocking is how Go coordinates concurrent work safely.

* A **channel** is a typed conduit used to communicate between goroutines.
* Channels are created with `make(chan T)`.
* Sending syntax: `ch <- value`
* Receiving syntax: `value := <-ch`
* Send and receive operations **block by default**.

* **Unbuffered channels** have no capacity.
* A send blocks until a receive happens.
* A receive blocks until a send happens.
* Unbuffered channels provide **synchronization (handshake)** between goroutines.

* **Buffered channels** have capacity: `make(chan T, n)`.
* Sending blocks only when the buffer is full.
* Receiving blocks only when the buffer is empty.
* Buffered channels **decouple sender and receiver timing**.
* A buffered channel with capacity 1 is **not the same** as an unbuffered channel.

* Channels are used to **share data by communicating**, not by shared memory.
* This avoids data races and makes ownership explicit.

* Goroutines often require explicit waiting.
* A common waiting pattern uses a **signal channel** (e.g. `done`).
* The empty struct type `struct{}` is used for signals because it carries no data.
* `struct{}` is a type; `struct{}{}` is its only value.

* Receiving without assignment (`<-ch`) is valid when only blocking/synchronization matters.
* Receiving a value without `ok` discards closure information.

* **Closing a channel** signals that no more values will be sent.
* Closing is a **state change**, not a send.
* Only the goroutine responsible for sending values should close the channel (convention).
* Sending on a closed channel causes a runtime panic.
* Closing a channel unblocks all receivers.
* Receiving from a closed channel yields zero values **after buffered values are drained**.

* The **comma-ok** form (`v, ok := <-ch`) detects channel closure.
  * `ok == true` → real value received
  * `ok == false` → channel closed and empty

* Ranging over a channel (`for v := range ch`) receives values until the channel is closed.
* `range` blocks waiting for values.
* `range` exits only when the channel is closed.
* If a channel is never closed, `range` never terminates.

* Channel **direction** can restrict usage:
  * `chan T` — send and receive
  * `chan<- T` — send-only
  * `<-chan T` — receive-only
* Direction is a **type restriction at the usage site**, not a property of the channel itself.
* Directional channels are commonly used at function boundaries to enforce correctness.

* Sending and receiving on an unbuffered channel in the **same goroutine** causes deadlock.
* Unbuffered channel operations that must rendezvous require **multiple goroutines**.

* Goroutines define concurrency.
* Channels define synchronization and communication.
* Correct Go concurrency relies on clear ownership, blocking, and explicit signaling.

---

**Day 7 — Concurrency Control & Coordination**

* Concurrency introduces two separate problems: **correctness** (safe memory access) and **lifecycle** (knowing when work is finished).
* These problems are orthogonal and require different tools.

* `sync.WaitGroup` solves a **waiting / lifecycle** problem.
* A WaitGroup lets one goroutine (usually `main`) wait until a fixed number of goroutines have finished.
* Internally, a WaitGroup maintains a counter.
* `wg.Add(n)` increments the counter by `n`.
* Each goroutine must call `wg.Done()` exactly once to decrement the counter.
* `wg.Wait()` blocks until the counter reaches zero.
* `wg.Add()` must be called before goroutines start.
* A WaitGroup does **not** protect shared memory and does **not** make code safe.
* A WaitGroup only ensures the program does not demonstrate incomplete results or exit early.

* `select` is used when a goroutine must wait on **multiple possible events** and react to whichever happens first.
* `select` blocks until one case is ready.
* If multiple cases are ready simultaneously, Go chooses one non-deterministically.
* This non-determinism is intentional and prevents hidden timing assumptions.
* `select` is not a replacement for `WaitGroup`.
* `select` is not used for ordering or correctness.
* `select` is primarily used for cancellation, timeouts, and multiplexing signals.
* If a goroutine waits on only one thing, `select` is unnecessary.
* `select` becomes necessary when waiting must be interruptible (e.g., shutdown, timeout).

* A `sync.Mutex` solves a **memory correctness** problem.
* A mutex ensures only one goroutine accesses a critical section at a time.
* A data race occurs when:
  * two or more goroutines access the same variable concurrently
  * at least one access is a write
  * and no synchronization is used
* Data races cause undefined behavior, even if results “look correct”.
* `mu.Lock()` must be called before accessing shared data.
* `mu.Unlock()` must be called after.
* Forgetting to unlock or locking twice in the same goroutine causes deadlock.
* Mutexes do not wait for goroutines to finish.
* Mutexes do not coordinate execution order.
* Mutexes only protect memory.

* Mutexes and WaitGroups are commonly used together:
  * mutex → correctness of shared data
  * waitgroup → completeness of execution
* Using one does not replace the other.

* Channels coordinate **communication and ownership**.
* Mutexes protect **shared state**.
* Idiomatic rule:
  * Share memory by communicating when possible.
  * Use mutexes only when shared memory is unavoidable.

* The Go race detector can detect data races at runtime.
* Run with: `go run -race main.go` or `go test -race`.
* If the race detector reports an issue, the program is incorrect regardless of output.

* Correct concurrent programs require:
  * proper synchronization (mutex or channels)
  * explicit waiting (WaitGroup or equivalent)
  * no reliance on timing or sleeps

--- 




