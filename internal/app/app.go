package app

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hrvadl/pkgcollision/internal/collision"
)

func New() *App {
	return &App{}
}

type App struct{}

func (a *App) Run(pkgPath string) error {
	if err := filepath.Walk(pkgPath, a.walk); err != nil {
		return err
	}

	return nil
}

func (a *App) walk(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	filename := filepath.Base(path)
	if !isGoFile(filename) {
		return nil
	}

	f, err := parser.ParseFile(fset, filename, content, 0)
	if err != nil {
		return err
	}

	detector := collision.NewDetector(f)
	detector.Detect()

	return nil
}

func isGoFile(path string) bool {
	filename := filepath.Base(path)
	return strings.HasSuffix(filename, ".go")
}
