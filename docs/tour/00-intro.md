# Welcome to the frederik-jatzkowski/errors Package Tour

Welcome! You're about to dive deep into this thoughtfully designed error handling library for Go.

## Why This Package Exists

Go's standard error handling is elegant in its simplicity, but when things go wrong in production,
you often find yourself asking: "Where did this error come from?"
Traditional solutions either add stack traces everywhere (cluttering your logs)
or nowhere (leaving you in the dark).

This package solves that problem with **much more refined stack trace management**:
it adds stack traces only when they're needed, keeps sentinel errors clean,
and gives you the flexibility to format good errors for both humans and machines.

## Key Highlights

### 🎯 Drop-In Replacement
Every function in this package is a drop-in replacement for the standard library.
Replace standard library `errors.New` with this package's `errors.New`, `fmt.Errorf` with `errors.Errorf`, and `errors.Join` with `errors.Join`.
Your code stays idiomatic and you get stack traces for free.
We will also add future std library error handling APIs to this package with minimal delay.

### 🧠 Smart Stack Trace Management
The package is intelligent about when to add stack traces:
- **Adds** stack traces when errors without them enter your error tree
- **Skips** stack traces for sentinel errors (defined at package level)
- **Prevents** duplicate stack traces in deeply nested error chains

### 📊 Dual Formatting Modes
The same error can be formatted two ways:
- **Human-readable**: `%s` or `%v` gives you clean, concise error messages
- **Full context**: `%+v` gives you complete stack traces for debugging

### ⚡ Performance Conscious
Stack traces are only captured when necessary.
Once an error has a stack trace, wrapping it multiple times is an amortized constant-time operation.
The package is designed to have minimal overhead.

### 🔒 Statically Enforced Best Practices
With the right linter configuration, you can ensure that:
- Standard library error functions are never used
- External errors are always wrapped with stack traces
- Your team follows consistent error handling patterns

## What You'll Learn

This tour is structured as a series of lessons, each building on the previous one. By the end, you'll understand not just how to use this package, but why it's designed the way it is, and how to integrate it effectively into your projects.

## Table of Contents

1. **[Getting Started](./01-getting-started.md)**
   - Installation and basic setup
   - Your first error with a stack trace
   - Understanding the difference between `%v` and `%+v`

2. **[Creating Errors](./02-creating-errors.md)**
   - `errors.New()`: Creating simple errors
   - `errors.Errorf()`: Formatting errors with context
   - When stack traces are added (and when they're not)

3. **[Wrapping Errors](./03-wrapping-errors.md)**
   - Using the `%w` verb effectively
   - Single vs. multiple error wrapping
   - Preserving error identity with `errors.Is()` and `errors.As()`

4. **[Joining Multiple Errors](./04-joining-errors.md)**
   - When to use `errors.Join()`
   - Combining errors from parallel operations
   - Error tree traversal and inspection

5. **[Stack Traces Deep Dive](./05-stack-traces.md)**
   - How stack traces are captured
   - When stack traces are added vs. skipped
   - Understanding the stack trace format
   - Preventing duplicate stack traces
   - Explicitly adding stack traces with `WithStack`

6. **[Sentinel Errors](./06-sentinel-errors.md)**
   - Defining sentinel errors correctly
   - Why sentinel errors don't get stack traces
   - Best practices for error constants

7. **[Error Formatting](./07-error-formatting.md)**
   - All the formatting verbs and their behavior
   - Human-readable vs. machine-readable output
   - Structured logging with slog
   - Custom formatting strategies

8. **[Linter Integration](./08-linter-integration.md)**
   - Configuring `golangci-lint` for this package
   - Using `wrapcheck` to enforce wrapping
   - Using `forbidigo` to prevent stdlib usage
   - Setting up your team's error handling standards

9. **[Migration Guide](./09-migration.md)**
    - Migrating from standard library errors
    - Migrating from `github.com/pkg/errors`
    - Common migration patterns and pitfalls

10. **[Advanced Topics](./10-advanced.md)**
    - Custom error types and this package
    - Integration with other error libraries
    - Edge cases and caveats
    - Contributing to the project

## Prerequisites

Before starting this tour, you should be familiar with:
- Go's standard error handling (`errors` package)
- The `fmt` package and formatting verbs
- Basic understanding of Go's error wrapping (`%w` verb, `errors.Is`, `errors.As`)
- Go modules and dependency management

---

Ready to get started? Let's begin with [Getting Started](./01-getting-started.md)!
