# What is context in Go?
The context package in Go provides a way to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines.

It is bassicall a advanced done channel concurrency pattern.

Refer this: [Golang Context](https://youtu.be/8omcakb31xQ?si=aT5YkgqHZFZlGu9r)

**Use Cases:**

* Cancel goroutines when work is no longer needed

* Enforce timeouts (e.g., HTTP requests, DB calls)

* Pass request-scoped values like user ID, trace ID, etc.

####  Context is Immutable
Every time you add a value or timeout, you get a new derived context. The original is unchanged.

### Types of Contexts
All contexts are built from context.Background() or context.TODO():

**✅ context.Background()**
* The root context.
* use it when there's no other context to use (e.g., in main() or init).

**❓ context.TODO()**
* Placeholder context.
* Use it during development when you don’t yet have logic to determine the context.

## Context Functions

**1. context.WithCancel(parent)**

Returns a derived context that can be canceled manually.

```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
    <-ctx.Done()
    fmt.Println("cancelled")
}()
cancel() // stops the context
```

**2. context.WithTimeout(parent, duration)**

Returns a derived context that auto-cancels after timeout.

``` go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
<-ctx.Done() // wait for timeout
```

**3. context.WithDeadline(parent, time.Time)**

Like timeout, but for a fixed point in time.

```go
deadline := time.Now().Add(1 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

**4. context.WithValue(parent, key, value)**

Add request-scoped values.

``` go
ctx := context.WithValue(context.Background(), "userID", 42)
userID := ctx.Value("userID").(int)
```

⚠️ Don’t use WithValue for passing optional config/data — use explicit parameters. Keep it request scoped only.


**ctx.Done(), ctx.Err()**

**ctx.Done()** returns a channel that's closed when:
* context is canceled

* timeout/deadline expires

**ctx.Err()** returns:

* context.Canceled

* context.DeadlineExceeded