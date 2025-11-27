# Lesson 6: Sentinel Errors

Sentinel errors are a fundamental pattern in Go error handling.
They're predefined error values that represent specific error conditions.
In this lesson, we'll learn how to define and use sentinel errors effectively with this package.

## What Are Sentinel Errors?

Sentinel errors are error values that are defined once and reused. They're typically defined at package level and represent specific error conditions that can be checked with `errors.Is()`.

```go
var ErrNotFound = errors.New("not found")
var ErrPermissionDenied = errors.New("permission denied")
var ErrInvalidInput = errors.New("invalid input")
```

## Defining Sentinel Errors

You can define sentinel errors using `errors.New()`, `errors.Errorf()`:

```go
package mypackage

import "github.com/frederik-jatzkowski/errors"

// Simple sentinel error
var ErrNotFound = errors.New("not found")

// Formatted sentinel error (added by this package)
// You can build more expressive errors by reusing defined constants
const AllowedRetries = 3
var ErrInvalidUser = errors.Errorf("too many retries (max. %d)", AllowedRetries)
```

All of these are detected as sentinel errors and won't get stack traces.

## Using Sentinel Errors

### Returning Sentinel Errors

```go
func FindUser(id int) (*User, error) {
	if id < 0 {
		return nil, ErrNotFound
	}
	// ... lookup logic
	return user, nil
}
```

### Checking for Sentinel Errors

Use `errors.Is()` to check for sentinel errors.
Sentinel errors can be wrapped, and they'll still be detectable:

```go
func processUser(id int) error {
	user, err := FindUser(id)
	if err != nil {
		return errors.Errorf("failed to process user %d: %w", id, err)
	}
	// ... process user
	return nil
}

func main() {
	err := processUser(123)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User was not found!")
	}
}
```

Even though `ErrNotFound` is wrapped, `errors.Is()` will still find it.

## Why No Stack Traces?

Sentinel errors don't get stack traces because:

1. **They're defined once**: The stack trace would always point to the package initialization, which isn't useful
2. **They're meant to be compared**: Sentinel errors are checked with `errors.Is()`, not debugged
3. **They keep logs clean**: Avoiding stack traces for sentinel errors keeps error logs focused on actual problems

## Best Practices for Sentinel Errors

### 1. Define at Package Level

```go
// ✅ Good: Package-level definition
var ErrNotFound = errors.New("not found")

// ❌ Bad: Function-level (will get stack trace, no identity with errors.Is)
func ErrNotFound() error {
	return errors.New("not found")
}
```

### 2. Export When Needed

```go
// ✅ Good: Export if callers need to check for it
var ErrNotFound = errors.New("not found")

// ✅ Good: Unexported if it's internal
var errInternal = errors.New("internal error")
```

### 3. Group Related Errors

```go
var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrInvalidInput  = errors.New("invalid input")
)
```

## Example

```go
const (
	MinAge = 6
)

var (
	ErrEmptyName  = errors.New("name cannot be empty")
	ErrMinAge     = errors.Errorf("minimum age is %d", MinAge)
)

func ValidateUser(name string, age int) error {
	if name == "" {
		return ErrEmptyName
	}
	if age < MinAge {
		return errors.Errorf("age was %d: %w", age, ErrMinAge)
	}
	return nil
}
```

When you check `errors.Is(err, ErrMinAge)`, it will still work.
The stack trace will show where the error was wrapped, not where the sentinel was defined.

## What's Next?

Now that you understand sentinel errors, let's explore all the formatting options available. In the [next lesson](./07-error-formatting.md), we'll learn about different format verbs and how to use them effectively.

---

**Previous**: [Stack Traces Deep Dive](./05-stack-traces.md) | **Next**: [Error Formatting](./07-error-formatting.md)

