# Lesson 8: Linter Integration

One of the best ways to ensure consistent error handling across your team is to use linters. In this lesson, we'll configure `golangci-lint` to enforce best practices and prevent common mistakes.

Linters help you:
1. **Enforce consistency**: Everyone uses the same error handling patterns
2. **Catch mistakes early**: Find issues before they reach production
3. **Prevent regressions**: Stop accidental use of standard library errors
4. **Improve code quality**: Ensure errors are always wrapped at boundaries

## Required Linters

For this package, you'll want to enable two key linters:

1. **`forbidigo`**: Prevents use of standard library error functions
2. **`wrapcheck`**: Ensures external errors are wrapped

## Configuring `forbidigo`

The `forbidigo` linter prevents usage of forbidden patterns. Here's how to configure it:

```yaml
linters:
  enable:
    - forbidigo
  settings:
    forbidigo:
      forbid:
        # Disallow the entire stdlib errors package
        - pattern: ^errors\..*$
          pkg: ^errors$
          msg: Use github.com/frederik-jatzkowski/errors.* instead.
        # Disallow fmt.Errorf
        - pattern: ^fmt\.Errorf$
          pkg: ^fmt$
          msg: Use github.com/frederik-jatzkowski/errors.Errorf instead.
      analyze-types: true
```

### Example Violations

```go
// ❌ This will be flagged by forbidigo
import "errors"
err := errors.New("test")  // Error: Use github.com/frederik-jatzkowski/errors.* instead.

// ❌ This will also be flagged
import "fmt"
err := fmt.Errorf("test")  // Error: Use github.com/frederik-jatzkowski/errors.Errorf instead.

// ✅ This is allowed
import "github.com/frederik-jatzkowski/errors"
err := errors.New("test")  // OK
```

## Configuring `wrapcheck`

The `wrapcheck` linter ensures that errors from external packages are wrapped. Here's the configuration:

```yaml
linters:
  enable:
    - wrapcheck
  settings:
    wrapcheck:
```

### Example Violations

```go
// ❌ This will be flagged by wrapcheck
func process() error {
	err := externalPackage.DoSomething()
	return err  // Error: error returned from external package must be wrapped
}

// ✅ This is correct
func process() error {
	err := externalPackage.DoSomething()
	if err != nil {
		return errors.Errorf("processing failed: %w", err)
	}
	return nil
}
```

## What's Next?

Now that you know how to enforce best practices with linters, let's learn how to migrate existing code. In the [next lesson](./09-migration.md), we'll cover migrating from standard library errors and other error handling packages.

---

**Previous**: [Error Formatting](./07-error-formatting.md) | **Next**: [Migration Guide](./09-migration.md)
