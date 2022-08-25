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
	t.Add("Another")
	fmt.Println("Added an item hello")
	fmt.Println(t)

	fmt.Printf("\nPending activities: %d\n", t.CountPending())

	t.Complete(1)
	fmt.Println("Completed task 1", t)

	//	delete a task
	t.Delete(1)
	fmt.Println("Deleted task hello", t)
}
