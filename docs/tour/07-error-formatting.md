# Lesson 7: Error Formatting

One of the most powerful features of this package is its flexible formatting system. The same error can be displayed in multiple ways depending on your needs. In this lesson, we'll explore all the formatting options and when to use each one.

## The Two Modes: Human-Readable vs. Debug

This package provides two distinct formatting modes:

1. **Human-readable**: Clean, concise error messages (for users and logs)
2. **Debug mode**: Full stack traces (for debugging and development)

## Format Verbs

### `%v` - Default Format

The `%v` verb shows just the error message, without stack traces:

```go
err := errors.New("file not found")
fmt.Printf("%v\n", err)
// Output: file not found
```

### `%s` - String Format

The `%s` verb behaves the same as `%v`:

```go
err := errors.New("file not found")
fmt.Printf("%s\n", err)
// Output: file not found
```

### `%+v` - Verbose Format (With Stack Traces)

The `%+v` verb shows the error message plus the full stack trace:

```go
err := errors.New("file not found")
fmt.Printf("%+v\n", err)
// Output:
// file not found
//     main.main
//         /path/to/main.go:8
//     runtime.main
//         runtime/proc.go:250
//     runtime.goexit
//         runtime/asm_arm64.s:1224
```

## Key Takeaways

1. **`%v` and `%s`**: Show just the error message (human-readable)
2. **`%+v`**: Shows error message plus stack trace (debug mode)
3. **Never expose stack traces to users**: Always use `%v` or `err.Error()` for user-facing errors

## What's Next?

Now that you understand error formatting, let's learn how to enforce best practices with linters. In the [next lesson](./08-linter-integration.md), we'll configure `golangci-lint` to ensure consistent error handling across your team.

---

**Previous**: [Sentinel Errors](./06-sentinel-errors.md) | **Next**: [Linter Integration](./08-linter-integration.md)
