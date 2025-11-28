# Lesson 2: Creating Errors

In this lesson, we'll explore the two main ways to create errors in this package: `errors.New()` and `errors.Errorf()`. We'll also understand the intelligent stack trace management that makes this package special.

## Creating Simple Errors with `errors.New()`

The simplest way to create an error is using `errors.New()`:

```go
err := errors.New("file not found")
```

This works exactly like the standard library `errors.New()`, but with one important addition: it automatically captures a stack trace at the point where it's called.

### When Stack Traces Are Added

Here's where it gets interesting. The package is smart about when to add stack traces:

```go
package main

import (
	"fmt"
	"github.com/frederik-jatzkowski/errors"
)

// Sentinel error - defined at package level
var ErrNotFound = errors.New("not found")

func main() {
	// Runtime error - created during execution
	runtimeErr := errors.New("runtime error")
	
	fmt.Println("Sentinel error (%+v):")
	fmt.Printf("%+v\n\n", ErrNotFound)
	
	fmt.Println("Runtime error (%+v):")
	fmt.Printf("%+v\n", runtimeErr)
}
```

Run this, and you'll notice:
- **Sentinel errors** (defined at package level) don't get stack traces
- **Runtime errors** (created during execution) do get stack traces

This is intentional! Sentinel errors are meant to be compared with `errors.Is()`, not debugged. Stack traces are only useful when errors occur at runtime.

## Formatting Errors with `errors.Errorf()`

For more complex error messages, use `errors.Errorf()`:

```go
err := errors.Errorf("failed to process user %d: %s", userID, reason)
```

This works like `fmt.Errorf()`, but again with automatic stack trace support. More importantly, it fully supports the `%w` verb for error wrapping (which we'll cover in detail in the [next lesson](./03-wrapping-errors.md)).

### Simple Formatting (No Wrapping)

When you don't use `%w`, `errors.Errorf()` creates a simple formatted error with a stack trace (if needed):

```go
err := errors.Errorf("user %d not found", 123)
fmt.Printf("%v\n", err)  // Output: user 123 not found
```

Even without wrapping, `errors.Errorf()` will add a stack trace unless the error is created in a sentinel context.

### With Error Wrapping

When you use `%w`, the behavior becomes more sophisticated:

```go
baseErr := errors.New("database connection failed")
err := errors.Errorf("operation failed: %w", baseErr)
```

This wraps the error, preserving the error chain for `errors.Is()` and `errors.As()`. We'll explore this in detail in the [next lesson](./03-wrapping-errors.md).

## Key Differences from Standard Library

| Feature | Standard Library | This Package |
|---------|----------------|--------------|
| `errors.New()` | Creates error | Creates error + stack trace |
| `fmt.Errorf()` | Formats error | Formats error + stack trace |
| Stack traces | Not available | Automatic when needed |

## Best Practices

1. **Use `errors.New()`** for simple error messages
2. **Use `errors.Errorf()`** when you need formatting or error wrapping
3. **Define sentinel errors** at package level (they won't clutter logs with stack traces)
4. **Create runtime errors** in functions (they'll have helpful stack traces)

## What's Next?

Now that you know how to create errors, let's learn about wrapping them and creating an error tree.
In the [next lesson](./03-wrapping-errors.md), we'll explore the `%w` verb, error chains, and how to preserve error identity.

---

**Previous**: [Getting Started](./01-getting-started.md) | **Next**: [Wrapping Errors](./03-wrapping-errors.md)
