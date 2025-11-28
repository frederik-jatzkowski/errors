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
```

Note that the default behavior of this package already omits Go internals from the stack trace.

## Formatting External Errors

When mixing this package with other error handling libraries (e.g., during migration), you may want to preserve stack trace information from external errors. By default, the `%+v` verb is not forwarded to external errors to maintain consistent formatting.

However, you can enable advanced formatting of external errors using `GlobalFormatSettings()` with `WithAdvancedFormattingOfExternalErrors()` once in your `main` function:

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithAdvancedFormattingOfExternalErrors(),
    )
    // ... rest of your application
}
```

This setting allows forwarding of the `%+v` verb to external errors, which can provide more debugging information when working with errors from other libraries. However, keep in mind:

- **More debugging info**: You'll get additional stack traces from external error libraries
- **Potential redundancy**: You might see duplicate or redundant stack traces
- **Formatting inconsistencies**: The formatting of external errors might not align perfectly with this package's formatting style

This is particularly useful during gradual migration from other error handling libraries, as it helps preserve debugging information while you transition.

## Customizing Stack Trace Display

You can further customize how stack traces are displayed using additional format settings.

### Stripping File Name Prefixes

To make stack traces more concise, you can strip common prefixes from file paths. This is especially useful when your code is in a module with a long import path:

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
    )
    // ... rest of your application
}
```

With this setting, file paths like:
```
github.com/frederik-jatzkowski/errors/examples/nested/main.go:21
```

Will be displayed as:
```
examples/nested/main.go:21
```

This keeps stack traces clean and readable while preserving all essential debugging information.

### Stripping Function Name Prefixes

Similarly, you can strip common prefixes from function names in stack traces. This works together with `WithStrippedFileNamePrefix` to create cleaner, more concise stack traces:

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
        errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
    )
    // ... rest of your application
}
```

With this setting, function names like:
```
github.com/frederik-jatzkowski/errors/examples/nested/subpackage.SomethingBad
```

Will be displayed as:
```
examples/nested/subpackage.SomethingBad
```

This complements the file name prefix stripping to create even cleaner stack traces.

### Ignoring Function Prefixes

You can also configure which function name prefixes should be ignored in stack traces. This is useful to keep Go internals out of your application logs:

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithIgnoredFunctionPrefixes("runtime", "internal/runtime", "testing", "myapp/internal"),
    )
    // ... rest of your application
}
```

The default already ignores `[]string{"runtime", "internal/runtime", "testing"}`, so you typically only need to add your own internal prefixes if desired.

## Key Takeaways

1. **`%v` and `%s`**: Show just the error message (human-readable)
2. **`%+v`**: Shows error message plus stack trace (debug mode)
3. **Never expose stack traces to users**: Always use `%v` or `err.Error()` for user-facing errors
4. **External errors**: Use `GlobalFormatSettings(errors.WithAdvancedFormattingOfExternalErrors())` if you need to preserve stack traces from other error libraries

## What's Next?

Now that you understand error formatting, let's learn how to enforce best practices with linters. In the [next lesson](./08-linter-integration.md), we'll configure `golangci-lint` to ensure consistent error handling across your team.

---

**Previous**: [Sentinel Errors](./06-sentinel-errors.md) | **Next**: [Linter Integration](./08-linter-integration.md)
