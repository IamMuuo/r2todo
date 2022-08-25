//////////////////////////////////////////////////////
//	Project:	Godo
//	Author:		Erick Muup
//	File:		Todo.go
//	Descr:		Contains struct definition for usage
//				in the application
//////////////////////////////////////////////////////

package util

import "time"

// A single todo item.
type Todo struct {
	Task      string
	Date      time.Time
	Done      bool
	Completed time.Time
}

type Todos []Todo // A typedef for an array of todos

func (t *Todos) Add(task string) {
	// Append a task to the list of items
	todo := Todo{
		Task:      task,
		Date:      time.Now(),
		Done:      false,
		Completed: time.Time{},
	}

	*t = append(*t, todo)
}
