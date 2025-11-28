# Lesson 10: Advanced Topics

In this final lesson, we'll explore advanced topics including custom error types and caveats.

## Custom Error Types

This package works with custom error types as long as they are in the leaves of error trees.
You can define your own error types and use them with all the package's features.

### Defining Custom Error Types

```go
type ValidationError struct {
	Field string
	Value interface{}
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Msg)
}
```

### Using Custom Errors

```go
func validateEmail(email string) error {
	if !isValidEmail(email) {
		return &ValidationError{
			Field: "email",
			Value: email,
			Msg:   "invalid format",
		}
	}
	return nil
}
```

### Wrapping Custom Errors

```go
err := validateEmail("invalid")
wrapped := errors.Errorf("user validation failed: %w", err)

// Still works with errors.As
var valErr *ValidationError
if errors.As(wrapped, &valErr) {
	fmt.Printf("Field: %s\n", valErr.Field)
}
```

## Edge Cases and Caveats

### Caveat 1: Mixing With External Errors

Using external errors is fine as long as you don't wrap this package's errors with external error wrappers.
If you wrap this package's errors with external error wrappers, stack traces may be lost.

If you need to preserve stack trace information from external error libraries (e.g., during gradual migration), you can enable advanced formatting:

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithAdvancedFormattingOfExternalErrors(),
    )
    // ... rest of your application
}
```

This setting forwards the `%+v` verb to external errors, providing more debugging information. However, this comes with trade-offs:
- **More debugging info**: Stack traces from external libraries are preserved
- **Potential redundancy**: You may see duplicate or redundant stack traces
- **Formatting inconsistencies**: External error formatting may not align perfectly with this package's style

Use this setting judiciously, primarily during migration periods or when you need to debug issues involving external error libraries.

### Caveat 2: IDE Warnings About `%w`

Some IDEs may warn that `%w` is illegal in `errors.Errorf()`.
This is a false positive. The package fully supports `%w` and `govet` will not complain.

## Performance Considerations

There is some performance overhead to the addition of stack traces due to the necessary `runtime.Callers` call.
However, adding new layers of nested errors is an amortized constant time effort per added layer.

### Benchmark Results

Standard Library Join without checks for stack traces:
```
goos: darwin
goarch: arm64
pkg: github.com/frederik-jatzkowski/errors
cpu: Apple M2 Max
BenchmarkStdJoin
BenchmarkStdJoin-12    	57709878	        20.80 ns/op
```

When no stack exists, we have to do a `runtime.Callers` call:

```
BenchmarkJoin-12    	 1966750	       611.5 ns/op
```

After 5 nested layers and stack already exists:
```
BenchmarkJoin_Deep5-12    	14539718	        81.43 ns/op
```

After 50 nested layers:
```
BenchmarkJoin_Deep50-12    	14128707	        82.86 ns/op
```

Similar behavior is observable for the other public functions of this package.

## Conclusion

Congratulations! You've completed the tour of the errors package. You now understand:

- How to create and format errors
- How to wrap and join errors
- How stack traces work
- How to use sentinel errors
- How to format errors for different contexts
- Linter integration
- Migration strategies
- Advanced patterns

You're ready to use this package effectively in your Go projects!

## Additional Resources

- [Package Documentation](https://pkg.go.dev/github.com/frederik-jatzkowski/errors)
- [GitHub Repository](https://github.com/frederik-jatzkowski/errors)
- [Go Error Handling Best Practices](https://go.dev/blog/error-handling-and-go)

---

**Previous**: [Migration Guide](./09-migration.md) | **Back to**: [Introduction](./00-intro.md)
