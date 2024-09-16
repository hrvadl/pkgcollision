package main

import (
	"github.com/hrvadl/pkgcollision/internal/collision"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(&collision.Analyzer)
}
