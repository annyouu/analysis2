package main

import (
	"fmt"
	"go/types"
)

func main() {
	// 組み込み関数printlnを取得
	printlnObj := types.Universe.Lookup("println")
	fmt.Printf("nName: %s, nType:%s, nKind: %T", printlnObj.Name(), printlnObj.Type(), printlnObj)

	// 組み込み型 errorを取得
	errorObj := types.Universe.Lookup("error")
	fmt.Printf("nName: %s, nType: %s, nKind: %T\n", errorObj.Name(), errorObj.Type(), errorObj)

	// 組み込み型 intを取得
	intObj := types.Universe.Lookup("int")
	fmt.Printf("nName: %s, nType: %s, nKind: %T\n", intObj.Name(), intObj.Type(), intObj)

	// 組み込み定数trueを取得
	trueObj := types.Universe.Lookup("true")
	fmt.Printf("nName: %s nType: %s, nKind: %T\n", trueObj.Name(), trueObj.Type(), trueObj)
}