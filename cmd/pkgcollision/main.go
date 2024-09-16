package main

import (
	"fmt"

	"github.com/hrvadl/pkgcollision/internal/app"
	"github.com/hrvadl/pkgcollision/internal/collision"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	app := app.New()
	fmt.Println(app)
	singlechecker.Main(&collision.Analyzer)
}
