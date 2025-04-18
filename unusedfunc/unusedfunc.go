package unusedfunc

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/passes/inspect"
    "golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "unusedfunc",
	Doc: "使われていない関数を検出する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// 使用されている関数を集める
	usedFuncs := map[types.Object]bool{}
	for _, obj := range pass.TypesInfo.Uses {
		if _, ok := obj.(*types.Func); ok {
			usedFuncs[obj] = true
		}
	}

	// ASTを走査して全FUncDeclをチェック
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
		fn := n.(*ast.FuncDecl)

		// レシーバを除外(レシーバがあるからメソッド)
		if fn.Recv != nil {
			return
		}

		// エクスポート関数も除外
		if fn.Name.IsExported() {
			return
		}

		// initも除外
		if fn.Name.Name == "init" {
			return
		}

		// 関数名に対応する types.Objectを取得
		obj := pass.TypesInfo.ObjectOf(fn.Name)
		if obj == nil {
			return
		}

		// 使用されている関数でなければ、出力する
		if !usedFuncs[obj] {
			pos := pass.Fset.Position(fn.Name.Pos())
			pass.Reportf(fn.Name.Pos(), "unused function: %s at %s", fn.Name.Name, pos)
		}
	})
	return nil, nil
}