package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	// ソースコード文字列を定義
	src := `v+1`

	// ソースコードをASTのExprにパースする
	expr, err := parser.ParseExpr(src)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return
	}

	fset := token.NewFileSet()

	var buf bytes.Buffer

	// 抽象構文木をgofmtと同じスタイルでフォーマットしてbufに書き込む
	err = format.Node(&buf, fset, expr)
	if err != nil {
		fmt.Println("フォーマットエラー:", err)
		return
	}

	// 結果を出力
	fmt.Println(buf.String())
}