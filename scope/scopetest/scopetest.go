package scopetest

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "scopetest",
	Doc: "ノードの中のスコープを見つける",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		scope := pass.TypesInfo.Scopes[n]
		if scope != nil {
			fmt.Printf("Scpoe %T\n", n)
			fmt.Printf("Pos: %s\n", pass.Fset.Position(n.Pos()))
			for _, name := range scope.Names() {
				obj := scope.Lookup(name)
				fmt.Printf("%s %s (%T)\n", name, obj.Type(), obj)
			}
		}
	})
	return nil, nil
}