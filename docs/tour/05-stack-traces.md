# Lesson 5: Stack Traces Deep Dive

## When Stack Traces Are Added

The package is intelligent about when to add stack traces, avoiding it if they provide no benefit.

### Rule 1: Runtime Errors Get Stack Traces

Errors created during program execution get stack traces:

```go
func processFile(filename string) error {
	if filename == "" {
		return errors.New("filename is required")  // Gets stack trace
	}
	return nil
}
```

### Rule 2: Sentinel Errors Don't Get Stack Traces

Errors defined at package level (sentinel errors) don't get stack traces:

```go
var ErrNotFound = errors.New("not found")  // No stack trace

func lookup(id int) error {
	return ErrNotFound  // Still no stack trace
}
```

This is intentional!
There is no benefit to adding an `init` stack trace to your error message.
We'll cover this in detail in the [next lesson](./06-sentinel-errors.md).

### Rule 3: Stack Traces Are Added When Needed

When wrapping or joining errors, a stack trace is only added if at least one error in the chain lacks one:

```go
err := errors.New("runtime error") // err has a stack trace
errExternal := fmt.Errorf("external error")

// joined will have a stack trace because errExternal lacks one
joined := errors.Join(err, errExternal)
```

This helps you identify where external errors enter your codebase.

### Rule 4: Duplicate Stack Traces Are Prevented

When wrapping an error that already has a stack trace, no new stack trace is added:

```go
err1 := errors.New("first error")        // Has stack trace
err2 := errors.Errorf("wrapped: %w", err1)  // No new stack trace
err3 := errors.Errorf("wrapped again: %w", err2)  // Still no new stack trace
```

This keeps stack traces clean and focused on the original error location.
The users of this library should feel no hesitation when wrapping errors with meaningful messages.

## Sentinel Detection

The package detects sentinel errors by checking if the stack trace was captured during package initialization (`.init` function). This is why sentinel errors don't get stack traces.

```go
// This is detected as a sentinel
var ErrSentinel = errors.New("sentinel")

// This is NOT a sentinel (created in a function)
func createError() error {
	return errors.New("runtime error")  // Gets stack trace
}
```

Note that this relies on a convention within the go runtime.

## Performance Considerations

Stack trace capture has a performance cost because it calls `runtime.Callers()`. However:

1. **Stack traces are only captured when needed**: If an error already has a stack trace, no new one is captured
2. **Wrapping is fast**: Once a stack trace exists, wrapping it multiple times is an amortized constant-time operation
3. **The cost is upfront**: The capture happens once when the error is created

## Best Practices

1. **Wrap external errors**: Always wrap errors from other packages to add stack traces
2. **Don't worry about duplicates**: The package prevents duplicate stack traces automatically
3. **Use sentinel errors**: Feel free to define sentinel errors at package level without cluttering your logs

## What's Next?

Now that you understand stack traces, let's explore sentinel errors in detail. In the [next lesson](./06-sentinel-errors.md), we'll learn how to define and use sentinel errors effectively.

---

**Previous**: [Joining Multiple Errors](./04-joining-errors.md) | **Next**: [Sentinel Errors](./06-sentinel-errors.md)

