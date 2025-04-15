package main

import (
	"fmt"
	"go/types"
)

func main() {
	// 型情報を用意
	t1 := types.Typ[types.Int] // int
	t2 := types.Typ[types.Int] // int
	t3 := types.Typ[types.String] // string

	// 同じ true
	fmt.Println(types.Identical(t1, t2))

	// 違う false
	fmt.Println(types.Identical(t1, t3))

	// ② 名前付き型 (type MyInt int)
	myInt := types.NewNamed(
		types.NewTypeName(0, nil, "MyInt", nil), // 型名: MyInt
		t1,
		nil,
	)

	// 名前付き型は元の型と違うと見なされる → false
	fmt.Println(types.Identical(t1, myInt)) // false

}