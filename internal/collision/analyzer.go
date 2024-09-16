package collision

import (
	"bytes"
	"go/printer"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = analysis.Analyzer{
	Name: "pkgcollision",
	Doc:  "reports variable name collision with the package name",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		detector := NewDetector(file)
		collisions, err := detector.Detect()
		if err != nil {
			return nil, err
		}
		render(pass, collisions)
	}

	return nil, nil
}

func render(pass *analysis.Pass, c []Collision) {
	var buf bytes.Buffer
	for _, collision := range c {
		printer.Fprint(&buf, pass.Fset, collision.Token)
		pass.Reportf(
			collision.Pos,
			"found collision with package '%s': %s",
			collision.Pkg,
			buf.String(),
		)
		buf.Reset()
	}
}
