# errors

[![Go Reference](https://pkg.go.dev/badge/github.com/frederik-jatzkowski/cantor.svg)](https://pkg.go.dev/github.com/frederik-jatzkowski/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/frederik-jatzkowski/errors)](https://goreportcard.com/report/github.com/frederik-jatzkowski/errors)

A minimal and idiomatic error handling library for the Go Programming Language adding stack traces.

## Explicit Design Goals

- Drop-In replacement for the standard "errors" and "fmt" packages
  - Write idiomatic Go code
  - No stack traces in sentinel errors
  - No duplicate stack traces
- Statically enforced best practices.
- Locality of behaviour and low cognitive load.
- Easy conversion fom other error handling libraries:
  - [github.com/pkg/errors](https://github.com/pkg/errors)

## Enforce Proper Usage with Linters

This showcase uses [golangci-lint v2](https://golangci-lint.run/docs/).

In order to make sure, stack traces are properly applied at the boundaries to external modules,
we can use the [wrapcheck](https://golangci-lint.run/docs/linters/configuration/#wrapcheck) linter.
The default settings of wrapcheck already allow the recommended public api of this package.

Also, configure `forbidigo` forbid usage of the standard `errors.*` and `fmt.Errorf` functionality:

```yaml
linters:
  settings:
    forbidigo:
      forbid:
        # disallow the entire stdlib errors package
        - pattern: ^errors\..*$
          pkg: ^errors$
          msg: Use github.com/frederik-jatzkowski/errors.* instead.
        # disallow fmt.Errorf
        - pattern: ^fmt\.Errorf$
          pkg: ^fmt$
          msg: Use github.com/frederik-jatzkowski/errors.Errorf instead.
      analyze-types: true
```
