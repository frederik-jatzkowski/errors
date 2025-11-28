# errors

[![Go Reference](https://pkg.go.dev/badge/github.com/frederik-jatzkowski/errors.svg)](https://pkg.go.dev/github.com/frederik-jatzkowski/errors)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go)](https://golang.org/)
[![Tests](https://github.com/frederik-jatzkowski/errors/actions/workflows/tests.yml/badge.svg)](https://github.com/frederik-jatzkowski/errors/actions/workflows/tests.yml)
[![Linter](https://github.com/frederik-jatzkowski/errors/actions/workflows/linter.yml/badge.svg)](https://github.com/frederik-jatzkowski/errors/actions/workflows/linter.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/frederik-jatzkowski/errors)](https://goreportcard.com/report/github.com/frederik-jatzkowski/errors)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A smart and idiomatic error handling library for the Go Programming Language adding stack traces.

![Package Mascot](/docs/assets/gopher.png)

## Explicit Design Goals

- Drop-In replacement for the standard error handling in the "errors" and "fmt" packages:
- Write idiomatic Go code.
- Add stack traces whenever needed.
- Statically enforced best practices.
- Locality of behavior and low cognitive load.
- Easy conversion from other error handling libraries:
  - [Standard Library](https://pkg.go.dev/errors)
  - [github.com/pkg/errors](https://github.com/pkg/errors)

This module allows both for human-readable output formatting using `err.Error()` or the `%v`/`%s` verbs
or for full stack trace information using the `%+v` verbs.
It just adds enough stack traces that the origin of root errors can be easily reconstructed while allowing for great, human-readable error messages.

So, the same error can be formatted as 

```
call failed: processing id 123: 
    => double errorf: 
        => something bad happened, 
        => hi, abc
    => something else happened
```

or

```
call failed: processing id 123: 
    => double errorf: 
        => something bad happened
            github.com/frederik-jatzkowski/errors/examples/nested/subpackage.SomethingBad
                github.com/frederik-jatzkowski/errors/examples/nested/subpackage/errors.go:6
            main.main
                github.com/frederik-jatzkowski/errors/examples/nested/main.go:23, 
        => hi, abc
            main.main
                github.com/frederik-jatzkowski/errors/examples/nested/main.go:26
    => something else happened
        github.com/frederik-jatzkowski/errors/examples/nested/subpackage.SomethingElse
            github.com/frederik-jatzkowski/errors/examples/nested/subpackage/errors.go:10
        main.main
            github.com/frederik-jatzkowski/errors/examples/nested/main.go:29
```
Note that some internal function calls from the go runtime are already ignored by default (adjustable with `errors.WithIgnoredFunctionPrefixes(...)`).
You probably want to clean up the stack trace even more by setting

```go
errors.GlobalFormatSettings(
    errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
    errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
)
```
as a global format setting (adjusted to your project specifics).  This results in the following:

```
call failed: processing id 123: 
    => double errorf: 
        => something bad happened
            examples/nested/subpackage.SomethingBad
                examples/nested/subpackage/errors.go:6
            main.main
                examples/nested/main.go:23, 
        => hi, abc
            main.main
                examples/nested/main.go:26
    => something else happened
        examples/nested/subpackage.SomethingElse
            examples/nested/subpackage/errors.go:10
        main.main
            examples/nested/main.go:29
```

This message contains all information necessary for debugging the different origins and paths
of the error tree while being minimally clean and straightforward.

## Learn More

Want to dive deeper? Check out our comprehensive [Package Tour](./docs/tour/00-intro.md) that covers:
- Getting started and installation
- Creating and wrapping errors
- Stack trace management
- Sentinel errors
- Error formatting
- Linter integration
- Migration guides
- Advanced topics

## Roadmap

There are some additional features planned before reaching `v1`:
- [x] Forwarding of format verbs to nested errors. This allows for seamless step by step replacement.
- [x] Ignore certain function names in stack traces (e.g., go runtime functions).
- [x] Defining a stripped prefix for function names.
- [ ] Performance optimizations.
- [ ] Hardening for enterprise grade stability.

## Settings

You can configure the package's formatting behavior using `GlobalFormatSettings()` with various format settings. This should only be called once in your `main` function as it has global effects.

### WithAdvancedFormattingOfExternalErrors

If you're mixing this package with other error handling libraries (e.g., during migration), you can enable advanced formatting of external errors. This allows forwarding of the `%+v` verb to external errors, providing more stack trace information from other error handling libraries.

Keep in mind that this adds valuable debugging information but might also lead to redundant stack traces. Additionally, the formatting of wrapped errors might not work as nicely with the formatting of this package.

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithAdvancedFormattingOfExternalErrors(),
    )
    // ... rest of your application
}
```

### WithIgnoredFunctionPrefixes

You can configure which function name prefixes should be ignored in stack traces. This is useful to keep Go internals out of your application logs. The default is already set to `[]string{"runtime", "internal/runtime", "testing"}`.

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithIgnoredFunctionPrefixes("runtime", "internal/runtime", "testing", "myapp/internal"),
    )
    // ... rest of your application
}
```

### WithStrippedFileNamePrefix

You can configure a prefix that will be stripped from file names in stack traces. This is useful to keep stack traces clean and focused, showing only the necessary information. The default is `""` (no stripping).

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
    )
    // ... rest of your application
}
```

This setting transforms file paths like `github.com/frederik-jatzkowski/errors/examples/nested/main.go:21` into `examples/nested/main.go:21`, making stack traces more concise while preserving all essential debugging information.

### WithStrippedFuncNamePrefix

You can configure a prefix that will be stripped from function names in stack traces. This works together with `WithStrippedFileNamePrefix` to create cleaner, more concise stack traces. The default is `""` (no stripping).

```go
func main() {
    errors.GlobalFormatSettings(
        errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
        errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
    )
    // ... rest of your application
}
```

This setting transforms function names like `github.com/frederik-jatzkowski/errors/examples/nested/subpackage.SomethingBad` into `examples/nested/subpackage.SomethingBad`, complementing the file name prefix stripping for even cleaner stack traces.

## Caveats

- Many IDEs will warn you that using the `%w` verb is illegal in this library's `errors.Errorf` function.
This is a false positive and the official `govet` will not complain about this.
This package's `errors.Errorf` fully supports the `%w` verb.

## Acknowledgments

> <sub>Gopher artwork © 2009 Renée French.  
Used under the [Creative Commons Attribution 3.0 License](https://creativecommons.org/licenses/by/3.0/).  
Modified from the original Go gopher design.</sub>
