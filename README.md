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
- Locality of behaviour and low cognitive load.
- Easy conversion fom other error handling libraries:
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
            main.main
                github.com/frederik-jatzkowski/errors/examples/nested/main.go:17
            runtime/internal/atomic.(*Uint32).Load
                runtime/internal/atomic/types.go:194
            runtime.goexit
                runtime/asm_arm64.s:1198, 
        => hi, abc
            main.main
                github.com/frederik-jatzkowski/errors/examples/nested/main.go:20
            runtime/internal/atomic.(*Uint32).Load
                runtime/internal/atomic/types.go:194
            runtime.goexit
                runtime/asm_arm64.s:1198
    => something else happened
        main.main
            github.com/frederik-jatzkowski/errors/examples/nested/main.go:23
        runtime/internal/atomic.(*Uint32).Load
            runtime/internal/atomic/types.go:194
        runtime.goexit
                runtime/asm_arm64.s:1198
```

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
- Forwarding of format verbs to nested errors. This allows for seamless step by step replacement.
- Ignore certain function names in stack traces (eg. go runtime functions).
- Performance optimizations.
- Hardening for enterprise grade stability.
- Defining a stripped prefix for function names.

## Caveats

- If you use other error libraries in the error tree,
stack trace information deeper in the tree might not be formatted.
Because of this, you should use linters to prevent the usage of other error handling libraries in your code.
- Many IDEs will warn you that using the `%w` verb is illegal in this libraries `errors.Errorf` function.
This is a false positive and the official `govet` will not complain about this.
This packages `errors.Errorf` fully supports the `%w` verb.

## Acknowledgments

> <sub>Gopher artwork © 2009 Renée French.  
Used under the [Creative Commons Attribution 3.0 License](https://creativecommons.org/licenses/by/3.0/).  
Modified from the original Go gopher design.</sub>
