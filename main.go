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
	"os"
	"strings"

	"example.com/packages/util"
)

const (
	filename string = "Tasks.json"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Printf("Usage of: " + args[0] +
			"\n-add\n\tadd a new todo\n" +
			"-complete <int>\n\tmark todo as completed\n" +
			"-del <int>\n\tdelete a todo\n" +
			"-list\n\tlist all todos\n")

		os.Exit(1)
	}

	var t util.Todos

	t.Load(filename)

	if strings.ToLower(args[1]) == "-list" {
		fmt.Println(t)
		os.Exit(0)
	}

	if strings.ToLower(args[1]) == "-add" {
		if args[2] != "" {
			t.Add(args[2])
		} else {
			os.Exit(2)
		}
	}

	//t.Add("hello")
	//t.Add("Another")
	//fmt.Println("Added an item hello")
	fmt.Println(t)

	fmt.Printf("\nPending activities: %d\n", t.CountPending())

	//t.Complete(1)
	//fmt.Println("Completed task 1", t)

	//	delete a task
	//t.Delete(1)
	//fmt.Println("Deleted task hello", t)

	t.Store(filename)
}
