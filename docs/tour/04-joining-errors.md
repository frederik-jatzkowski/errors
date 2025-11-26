# Lesson 4: Joining Multiple Errors

Sometimes, multiple things can go wrong independently, and you want to capture all of them. That's where `errors.Join()` comes in. In this lesson, we'll learn when and how to use `errors.Join()` effectively.

## What is `errors.Join()`?

`errors.Join()` combines multiple errors into a single error.
Unlike wrapping (which creates a hierarchy), joining creates a flat collection of errors that all happened at the same level.

```go
err1 := errors.New("first error")
err2 := errors.New("second error")
joined := errors.Join(err1, err2)
```

## When to Use `errors.Join()`

Use `errors.Join()` when:
- Multiple independent operations fail
- You want to report all failures, not just the first one
- Errors are at the same logical level (not nested)

### Example: Parallel Operations

```go
package main

import (
	"fmt"
	"github.com/frederik-jatzkowski/errors"
)

func saveUser(userID int) error {
	return errors.New("failed to save user")
}

func sendEmail(userID int) error {
	return errors.New("failed to send email")
}

func processUser(userID int) error {
	return errors.Join(
		saveUser(userID), 
		sendEmail(userID), 
	)
}

func main() {
	err := processUser(123)
	fmt.Printf("%v\n", errors.Errorf("failed to process user: %w", err))
	// Output:
	// failed to process user: 
	//   => failed to save user
	//   => failed to send email
}
```

## Error Inspection with `errors.Join()`

Joined errors can be inspected using `errors.Is()` and `errors.As()`:

```go
var ErrSaveFailed = errors.New("failed to save user")
var ErrEmailFailed = errors.New("failed to send email")

func main() {
	err := processUser(123)
	
	// Check for specific errors
	if errors.Is(err, ErrSaveFailed) {
		fmt.Println("Save operation failed")
	}
	if errors.Is(err, ErrEmailFailed) {
		fmt.Println("Email operation failed")
	}
}
```

Both checks will return `true` if the joined error contains either error.

## Stack Trace Behavior

`errors.Join()` intelligently manages stack traces:

- **If all errors have stack traces**: No new stack trace is added
- **If any error lacks a stack trace**: A stack trace is added at the join point
- **Nil errors are filtered out**: `errors.Join(nil, err1, nil, err2)` is equivalent to `errors.Join(err1, err2)`

## Edge Cases

### All Nil Errors

```go
result := errors.Join(nil, nil, nil)
fmt.Println(result == nil)  // true
```

If all errors are nil, `errors.Join()` returns `nil`.

### Single Error

```go
err := errors.New("single error")
joined := errors.Join(err)
fmt.Println(errors.Is(joined, err))  // true
```

Joining a single error preserves its identity.

## Best Practices

1. **Use `errors.Join()` for independent failures**: When multiple things can fail independently
2. **Use wrapping with `Errorf` for dependent failures**: When one failure causes another
3. **Filter nil errors**: The package does this automatically, but be aware of it
4. **Combine with wrapping**: You can wrap a joined error to add context

## What's Next?

Now that you understand how to create, wrap, and join errors, let's dive deep into stack traces. In the [next lesson](./05-stack-traces.md), we'll explore how stack traces are captured, when they're added, and how to interpret them.

---

**Previous**: [Wrapping Errors](./03-wrapping-errors.md) | **Next**: [Stack Traces Deep Dive](./05-stack-traces.md)
