package ArgOverwritten

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "ArgOverwritten",
	Doc:  doc,
	Run:  run,
}

const (
	doc = "ArgOverwritten finds function arguments being overwritten"
)

func run(pass *analysis.Pass) (interface{}, error) {
	visitor := func(node ast.Node) bool {
		var typ *ast.FuncType
		var body *ast.BlockStmt
		switch fn := node.(type) {
		case *ast.FuncDecl: // Regular function
			typ = fn.Type
			body = fn.Body
		case *ast.FuncLit: // Anonymous function
			typ = fn.Type
			body = fn.Body
		}
		if typ == nil || body == nil { // Exclude other types but also external functions with missing body
			return true
		}
		if len(typ.Params.List) == 0 {
			return true
		}

		for _, field := range typ.Params.List {
			for _, arg := range field.Names {
				obj := pass.TypesInfo.ObjectOf(arg)
				ast.Inspect(body, func(node ast.Node) bool {
					assign, ok := node.(*ast.AssignStmt)
					if !ok {
						return true
					}
					for _, lhs := range assign.Lhs {
						ident, ok := lhs.(*ast.Ident)
						if !ok {
							continue
						}
						if pass.TypesInfo.ObjectOf(ident) == obj {
							message := fmt.Sprintf("\"%s\" overwrites func parameter \"%s\"", ident.Name, obj.Name())
							pass.Report(analysis.Diagnostic{
								Pos:     ident.Pos(),
								Message: message,
							})
						}
					}
					return true
				})
			}
		}
		return true
	}
	for _, f := range pass.Files {
		ast.Inspect(f, visitor)
	}
	return nil, nil
}
