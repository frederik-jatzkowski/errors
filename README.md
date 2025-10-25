# errors

[![Go Reference](https://pkg.go.dev/badge/github.com/frederik-jatzkowski/cantor.svg)](https://pkg.go.dev/github.com/frederik-jatzkowski/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/frederik-jatzkowski/errors)](https://goreportcard.com/report/github.com/frederik-jatzkowski/errors)

A smart and idiomatic error handling library for the Go Programming Language adding stack traces.

![Package Mascot](/docs/assets/gopher.png)

## Explicit Design Goals

- Drop-In replacement for the standard error handling in the "errors" and "fmt" packages
  - Write idiomatic Go code
  - No stack traces in sentinel errors
  - Add stack traces whenever errors without stack traces enter the error tree.
- Statically enforced best practices.
- Locality of behaviour and low cognitive load.
- Sentinel errors can be defined using `errors.Errorf` and `errors.New` or even `errors.Join`.
- Easy conversion fom other error handling libraries:
  - [github.com/pkg/errors](https://github.com/pkg/errors)

This module allows both for human-readable output formatting using `err.Error()` or the `%v`/`%s` verbs
or for full stack trace information using the `%+v` verbs.
It just adds enough stack traces that the origin of root errors can be easily reconstructed while allowing for great, human-readable error messages.

So, the same error can be formatted as 

```
call failed: doing a: something bad happened, doing b: external dependency error
```

or

```
call failed: doing a: something bad happened
    main.main
        github.com/frederik-jatzkowski/errors/examples/errorf/main.go:14
    internal/runtime/atomic.(*Uint32).Load
        internal/runtime/atomic/types.go:194
    runtime.goexit
        runtime/asm_arm64.s:1224
, doing b: external dependency error
    main.main
        github.com/frederik-jatzkowski/errors/examples/errorf/main.go:12
    internal/runtime/atomic.(*Uint32).Load
        internal/runtime/atomic/types.go:194
    runtime.goexit
        runtime/asm_arm64.s:1224
```

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

You may also enable `report-internal-errors` in the `wrapcheck` config to enforce very detailed human-readable error messages:

```yaml
linters:
  settings:
    wrapcheck:
      report-internal-errors: true
```

## Performance

There is some performance overhead to the addition of stack traces due to the necessary `runtime.Callers` call.
However, adding new layers of nested errors is an amortized constant time effort per added layer.

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

Similar behaviour is observable for the other public functions of this package.

# Caveats

- If you use other error libraries in the error tree, deeply nested stack traces will not be formatted.
This is because we have to forward the `%+v` verb to the error containing the stack trace.
If there is an error in the chain that does not implement `fmt.Formatter`,
only `err.Error()` will be called for formatting.
Because of this, you should use linters to prevent the usage of other error handling libraries in your code.
- Many IDEs will warn you that using the `%w` verb is illegal in this libraries `errors.Errorf` function.
This is a false positive and the official `govet` will not complain about this.
This packages `errors.Errorf` fully supports the `%w` verb.

# Acknowledgments

> <sub>Gopher artwork © 2009 Renée French.  
Used under the [Creative Commons Attribution 3.0 License](https://creativecommons.org/licenses/by/3.0/).  
Modified from the original Go gopher design.</sub>
