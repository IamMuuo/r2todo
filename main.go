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
	"strconv"

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

	switch args[1] {
	case "-list":
		t.Print()

	case "-add":
		if args[2] != "" {
			t.Add(args[2])
		} else {
			fmt.Println("Use godo -add \"Task\"")
			os.Exit(1)
		}

	case "-delete":
		x, _ := strconv.Atoi(args[2])

		t.Delete(x)

	case "-complete":
		val, err := strconv.Atoi(args[2])

		if err != nil {
			fmt.Println(err)
		}

		t.Complete(val)

	default:
		fmt.Printf("Usage of: " + args[0] +
			"\n-add\n\tadd a new todo\n" +
			"-complete <int>\n\tmark todo as completed\n" +
			"-del <int>\n\tdelete a todo\n" +
			"-list\n\tlist all todos\n")
	}
	t.Store(filename)
}
