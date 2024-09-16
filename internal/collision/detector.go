package collision

import (
	"errors"
	"go/ast"
	"go/token"
	"path/filepath"
	"slices"
	"strconv"
)

type Collision struct {
	Pos   token.Pos
	Pkg   string
	Token ast.Node
}

func NewDetector(f *ast.File) *Detector {
	return &Detector{
		astfile: f,
	}
}

type Detector struct {
	astfile *ast.File
}

func (d *Detector) Detect() ([]Collision, error) {
	pkgs, err := d.getImports()
	if err != nil {
		return nil, err
	}

	return d.findCollisions(pkgs), nil
}

func (d *Detector) findCollisions(target []string) []Collision {
	var found []Collision

	ast.Inspect(d.astfile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.DeclStmt:
			found = append(found, d.findGenericDeclarationCollisions(node, target)...)
		case *ast.AssignStmt:
			found = append(found, d.findShortDeclarationCollisions(node, target)...)
		}

		return true
	})

	return found
}

func (d *Detector) findGenericDeclarationCollisions(
	node *ast.DeclStmt,
	target []string,
) []Collision {
	genDecl, ok := node.Decl.(*ast.GenDecl)
	if !ok {
		return nil
	}

	collisions := make([]Collision, 0, 1)
	for _, spec := range genDecl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		for _, ident := range valueSpec.Names {
			if slices.Contains(target, ident.Name) {
				collisions = append(
					collisions,
					Collision{Pos: genDecl.Pos(), Pkg: ident.Name, Token: node},
				)
			}
		}
	}

	return collisions
}

func (d *Detector) findShortDeclarationCollisions(
	node *ast.AssignStmt,
	target []string,
) []Collision {
	collisions := make([]Collision, 0, 1)
	for _, expr := range node.Lhs {
		if ident, ok := expr.(*ast.Ident); ok {
			if slices.Contains(target, ident.Name) {
				collisions = append(
					collisions,
					Collision{Pos: ident.Pos(), Pkg: ident.Name, Token: node},
				)
			}
		}
	}

	return collisions
}

func (d *Detector) getImports() ([]string, error) {
	if d.astfile == nil {
		return nil, errors.New("ast file can not be nil")
	}

	imports := make([]string, 0, 1)
	for _, importdec := range d.astfile.Imports {
		if importdec.Name != nil {
			imports = append(imports, importdec.Name.Name)
			continue
		}

		pkg, err := strconv.Unquote(importdec.Path.Value)
		if err != nil {
			return nil, err
		}

		imports = append(imports, filepath.Base(pkg))
	}

	return imports, nil
}
