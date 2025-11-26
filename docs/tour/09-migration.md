# Lesson 9: Migration Guide

Migrating to this package is straightforward because it's designed as a drop-in replacement. In this lesson, we'll cover migrating from standard library errors and other error handling packages.

## Migration from Standard Library

The standard library `errors` package is the easiest to migrate from because this package is designed as a drop-in replacement.

### Step 1: Update `errors` Imports

For this, you can use find and replace:

```go
// ❌ Before
import "errors"

// ✅ After
import "github.com/frederik-jatzkowski/errors"
```

### Step 2: Update `fmt.Errorf` to `errors.Errorf`

For this, you can use find and replace:

```go
// ❌ Before
import "fmt"
err := fmt.Errorf("failed: %w", baseErr)

// ✅ After
err := errors.Errorf("failed: %w", baseErr)
```

Make sure the right import is used:
```go
import "github.com/frederik-jatzkowski/errors"
```

### Step 3: Run Linters

After migration, run your linters to catch any remaining standard library usage:

```bash
golangci-lint run
```

This will also highlight incorrect error handling practices.

## Migration from `github.com/pkg/errors`

There are two ways of migrating away from `github.com/pkg/errors`.

1. Add a replace directive to your `go.mod` file: 
```
replace github.com/pkg/errors => github.com/frederik-jatzkowski/errors
```
2. If you want a full replacement:
   1. Do a string replace of the package files and update the imports.
   2. Uninstall the old package `go get github.com/pkg/errors@none`.

This package provides compatibility functions for `pkg/errors`.
Fade them out gradually.

### Optional - Enable Advanced Formatting

During gradual migration, if you want to preserve stack trace information from `pkg/errors` or other external error libraries, you can enable advanced formatting:

```go
func main() {
    errors.EnableAdvancedFormattingOfExternalErrors()
    // ... rest of your application
}
```

This setting forwards the `%+v` verb to external errors, which can be helpful during the transition period but may result in redundant stack traces and formatting inconsistencies.

## What's Next?

In the final lesson, we'll explore advanced topics including custom error types and edge cases.
Let's continue to the [next lesson](10-advanced.md).

---

**Previous**: [Linter Integration](./08-linter-integration.md) | **Next**: [Advanced Topics](./10-advanced.md)
