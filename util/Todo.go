//////////////////////////////////////////////////////
//	Project:	Godo
//	Author:		Erick Muup
//	File:		Todo.go
//	Descr:		Contains struct definition for usage
//				in the application
//////////////////////////////////////////////////////

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/alexeyco/simpletable"
)

// A single todo item.
type Todo struct {
	Task      string    `"task":`
	Date      time.Time `"Date":`
	Done      bool      `"Done":`
	Completed time.Time `"Completed":`
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

func (t *Todos) Complete(index int) error {
	// complete a task
	ls := *t

	if index <= 0 || index >= len(ls) {
		return errors.New("invalid index")
	}

	(ls)[index-1].Completed = time.Now()
	(ls)[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {

	// Delete an item
	ls := *t

	if index <= 0 || index >= len(*t) {
		return errors.New("invalid index")
	}

	*t = append((ls)[:index-1], (ls)[index:]...)

	return nil
}

func (t *Todos) CountPending() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}

// file io

func (t *Todos) Load(filename string) error {
	// load items from disk
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil

}

func (t *Todos) Store(filename string) error {
	// store the todos into disk

	content, err := json.MarshalIndent(t, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, content, 0644)

	if err != nil {
		return nil
	}

	return nil

}

// printing to console mechanism

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created Date"},
			{Align: simpletable.AlignCenter, Text: "Completed Date"},
		},
	}

	var cells [][]*simpletable.Cell

	for index, item := range *t {
		index++

		done := func(bool) string {
			if !item.Done {
				return "Pending"
			} else {
				return "Done"
			}
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%3d", index)},
			{Text: fmt.Sprintf("%s", item.Task)},
			{Text: fmt.Sprintf("%s", done(item.Done))},
			{Text: fmt.Sprintf("%+v", item.Date)},
			{Text: fmt.Sprintf("%+v", item.Completed)},
		},
		)
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
