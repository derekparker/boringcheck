package boringcheck

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type funcVisitor struct {
	curpkg string
	curfn  string
	fns    []string
}

func (v *funcVisitor) Visit(node ast.Node) ast.Visitor {
	if pkg, ok := node.(*ast.Package); ok {
		v.curpkg = pkg.Name
		return v
	}
	if fn, ok := node.(*ast.FuncDecl); ok {
		if v.curfn != "" {
			v.fns = append(v.fns, fmt.Sprintf("%s.%s", v.curpkg, v.curfn))
		}
		if fn.Recv == nil {
			v.curfn = fn.Name.Name
		} else {
			switch t := fn.Recv.List[0].Type.(type) {
			case *ast.StarExpr:
				v.curfn = fmt.Sprintf("(*%s).%s", t.X, fn.Name.Name)
			case *ast.Ident:
				v.curfn = fmt.Sprintf("(%s).%s", t, fn.Name.Name)
			}
		}
	}
	if ifstmt, ok := node.(*ast.IfStmt); ok {
		if ifstmt.Cond != nil {
			if call, ok := ifstmt.Cond.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if x, ok := sel.X.(*ast.Ident); ok {
						if x.Name == "boring" && sel.Sel.Name == "Enabled" {
							v.curfn = ""
							return nil
						}
					}
				}
			}
		}
		return nil
	}
	return v
}

func BoringCheck(path string) []string {
	v := &funcVisitor{}
	filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, 0)
		if err != nil {
			return err
		}
		for _, pkg := range pkgs {
			ast.Walk(v, pkg)
		}
		return nil
	})

	return v.fns
}
