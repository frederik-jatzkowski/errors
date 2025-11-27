# Lesson 3: Wrapping Errors

Error wrapping is one of Go's most powerful features for building error chains.
In this lesson, we'll learn how to use the `%w` verb with `errors.Errorf()` to create meaningful error trees
while preserving the ability to inspect the original errors.

## The `%w` Verb

The `%w` verb is Go's standard way to wrap errors.
This package fully supports it, and also allows you to wrap multiple errors in a single call.

The most common pattern is wrapping a single error with additional context:

```go
_, err := os.ReadFile(filename)
if err != nil {
    return errors.Errorf("failed to process %s: %w", filename, err)
}
```

## Error Identity Preservation

One of the key benefits of error wrapping is that it preserves error identity.
This means you can check for specific errors even when they're deeply wrapped.

### Using `errors.Is()`

```go
package main

import (
	"fmt"
	
	"github.com/frederik-jatzkowski/errors"
)

var ErrDatabase = errors.New("not authorized")
var ErrNotFound = errors.New("resource not found")

func operation() error {
	return errors.Errorf("query failed: %w", ErrDatabase)
}

func main() {
	err := operation()
	
	if errors.Is(err, ErrDatabase) {
		fmt.Println("You are not authorized to access this operation.")
	} else if errors.Is(err, ErrNotFound) {
		fmt.Println("The required resource does not exist.")
	} else {
		fmt.Println("Something unexpected went wrong.")
	}
}
```

This allows you to present meaningful and secure messages to your user.

### Using `errors.As()`

For extracting custom error types from deep within the error tree, use `errors.As()`:

```go
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func main() {
	err := errors.Errorf("validation failed: %w", &ValidationError{
		Field: "email",
		Msg:   "invalid format",
	})
	
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Msg)
	}
}
```

## Stack Trace Behavior with Wrapping

When you wrap errors, the package intelligently manages stack traces:

1. **If the wrapped error already has a stack trace**: No new stack trace is added
2. **If the wrapped error lacks a stack trace**: A stack trace is added at the wrapping point
3. **If wrapping multiple errors**: A stack trace is added if any wrapped error lacks one

This prevents duplicate stack traces while ensuring every error chain has at least one stack trace.

## What's Next?

Now that you understand error wrapping, let's explore `errors.Join()` for combining multiple independent errors. In the [next lesson](./04-joining-errors.md), we'll learn when and how to use `errors.Join()` effectively.

---

**Previous**: [Creating Errors](./02-creating-errors.md) | **Next**: [Joining Multiple Errors](./04-joining-errors.md)

