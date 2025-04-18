package intexpr

import (
	"fmt"
    "go/ast"
    "go/token"
    "go/types"

    "golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/passes/inspect"
    "golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "intexpr",
	Doc: "型がintまたはuntyped intの式を全て見つける",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

// formatNodeで文字列化する
func formatNode(expr ast.Expr, fset *token.FileSet) string {
	pos := fset.Position(expr.Pos()).String()

	switch v := expr.(type) {
	case *ast.BasicLit:
		return pos + " " + v.Value
	case *ast.Ident:
		return pos + " " + v.Name
	case *ast.BinaryExpr:
		return pos + fmt.Sprintf("(%s %s %s)", formatNode(v.X, fset), v.Op, formatNode(v.Y, fset))
	default:
		return pos + fmt.Sprintf(" (unknown type: %T)", expr)
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	typedInt := types.Typ[types.Int]

	inspect.Preorder(nil, func(n ast.Node) {
		expr, ok := n.(ast.Expr)
		if !ok {
			return
		}

		// 型情報を取得
		typ := pass.TypesInfo.TypeOf(expr)
		if typ == nil {
			return
		}

		// typed intを判別
		if types.Identical(typ, typedInt) {
			pass.Reportf(expr.Pos(), "typed int expr: %s", formatNode(expr, pass.Fset))
			return
		}

		// untyped intリテラルを判別
		if tv, ok := pass.TypesInfo.Types[expr]; ok {
			if tv.Value != nil && tv.Type != nil {
				if b, ok2 := tv.Type.(*types.Basic); ok2 && b.Kind() == types.UntypedInt {
					pass.Reportf(expr.Pos(), "untyped int: %s", formatNode(expr, pass.Fset))
				}
			}
		}
	})
	return nil, nil
}