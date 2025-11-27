# Lesson 1: Getting Started

Welcome to your first lesson! In this lesson, we'll get you up and running with the errors package and understand the fundamental difference it makes to your error handling.

## Installation

Add the package to your Go module:

```bash
go get github.com/frederik-jatzkowski/errors
```

That's it! No configuration needed. The package is designed to be a drop-in replacement, so you can start using it immediately.

## Your First Error

Let's start with the simplest possible example. Create a file `main.go`:

```go
package main

import (
	"fmt"
	"github.com/frederik-jatzkowski/errors"
)

func main() {
	err := errors.New("something went wrong")
	fmt.Println(err)
}
```

Run it, and you'll see:

```
something went wrong
```

Looks familiar, right? That's because `errors.New` works exactly like the standard library version. But there's more beneath the surface.

## The Magic: `%+v` vs `%v`

The real power of this package becomes visible when you use the `%+v` formatting verb. Let's modify our example:

```go
package main

import (
	"fmt"
	"github.com/frederik-jatzkowski/errors"
)

func main() {
	err := errors.New("something went wrong")
	
	fmt.Println("Standard format (%v):")
	fmt.Printf("%v\n\n", err)
	
	fmt.Println("With stack trace (%+v):")
	fmt.Printf("%+v\n", err)
}
```

Now run it. You'll see something like:

```
Standard format (%v):
something went wrong

With stack trace (%+v):
something went wrong
    main.main
        /path/to/your/main.go:8
```

**This is the core feature**: The same error can be formatted two ways:
- **`%v` or `%s`**: Clean, human-readable error messages (perfect for user-facing output)
- **`%+v`**: Full stack trace information (perfect for debugging and logging)

## Why This Matters

In production, you often need both:
- **User-facing errors**: Clean messages without technical details
- **Debug logs**: Full context including where the error originated

With this package, you get both from the same error object. No need to maintain separate error types or custom formatting logic.

## What's Next?

Now that you understand the basics, let's dive deeper into creating errors. In the [next lesson](./02-creating-errors.md), we'll explore `errors.New()` and `errors.Errorf()` in detail, and learn when stack traces are added (and when they're not).

---

**Previous**: [Introduction](./00-intro.md) | **Next**: [Creating Errors](./02-creating-errors.md)
