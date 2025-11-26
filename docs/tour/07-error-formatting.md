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

## Formatting External Errors

When mixing this package with other error handling libraries (e.g., during migration), you may want to preserve stack trace information from external errors. By default, the `%+v` verb is not forwarded to external errors to maintain consistent formatting.

However, you can enable advanced formatting of external errors by calling `EnableAdvancedFormattingOfExternalErrors()` once in your `main` function:

```go
func main() {
    errors.EnableAdvancedFormattingOfExternalErrors()
    // ... rest of your application
}
```

This setting allows forwarding of the `%+v` verb to external errors, which can provide more debugging information when working with errors from other libraries. However, keep in mind:

- **More debugging info**: You'll get additional stack traces from external error libraries
- **Potential redundancy**: You might see duplicate or redundant stack traces
- **Formatting inconsistencies**: The formatting of external errors might not align perfectly with this package's formatting style

This is particularly useful during gradual migration from other error handling libraries, as it helps preserve debugging information while you transition.

## Key Takeaways

1. **`%v` and `%s`**: Show just the error message (human-readable)
2. **`%+v`**: Shows error message plus stack trace (debug mode)
3. **Never expose stack traces to users**: Always use `%v` or `err.Error()` for user-facing errors
4. **External errors**: Use `EnableAdvancedFormattingOfExternalErrors()` if you need to preserve stack traces from other error libraries

## What's Next?

Now that you understand error formatting, let's learn how to enforce best practices with linters. In the [next lesson](./08-linter-integration.md), we'll configure `golangci-lint` to ensure consistent error handling across your team.

---

**Previous**: [Sentinel Errors](./06-sentinel-errors.md) | **Next**: [Linter Integration](./08-linter-integration.md)
