# pkgcollision

pkgcollision is a program to check whether you have variable name collision with the packages you imported.

## Install

```sh
go install github.com/hrvadl/pkgcollision/cmd/pkgcollision
```

## Usage

To scan all packages run:

```sh
pkgcollision ./...
```

To scan specific packages run:

```sh
pkgcollision pkgName
```

## Example

Let's say you have following code:

```go
package main

import (
	"fmt"

	"github.com/hrvadl/pkgcollision/internal/app"
)

func main() {
	app := app.New()
	fmt.Println(app)
}
```

`pkgcollision` will produce the following output:

```sh
/Users/vadym.hrashchenko/go/pkgcollision/cmd/pkgcollision/main.go:10:2: found collision with package 'app': app := app.New()
```

## Rules

It forbids to name variables with the same name as imported packages. Package name collision can be annoying and can even lead
to unexpected errors, therefore it'd be better to avoid it.

## Inspired by

- [100 Go mistakes and how to avoid them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them)

## TODO

- Add ignore comments
- Add ignore path options
- Enhace README.md
- Add tests
- More meaningfull error message
