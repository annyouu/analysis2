package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	threshold := flag.Int("n", 5, "引数の個数")
	filename := flag.String("file", "", "解析対象のGoソースファイルパス")
	flag.Parse()
	fmt.Println("DEBUG: filename =", *filename)

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, *filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
		os.Exit(1)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// 関数宣言ノードだけを対象にする
		fd, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// 引数リストの長さを取得
		numParams := 0
		if fd.Type.Params != nil {
			for _, field := range fd.Type.Params.List {
				// field.Namesがnilなら無名関数とカウントする
				if len(field.Names) == 0 {
					numParams++
				} else {
					numParams += len(field.Names)
				}
			}
		}

		if numParams >= *threshold {
			pos := fset.Position(fd.Pos())
			fmt.Printf("関数 %s が引数 %d 個 (>= %d)です: %s\n",
				fd.Name.Name, numParams, *threshold, pos)
		}
		return true
	})
}