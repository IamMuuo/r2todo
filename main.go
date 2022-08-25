//////////////////////////////////////////////////////////////////
//	Project: Godo
//
//	Author:	Erick Muuo
//	File  : main.go
//	Description: Contains program entry pint
//	//////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"example.com/packages/util"
)

func main() {
	var t util.Todos

	t.Add("hello")

	fmt.Println(t)
}
