package controllers

import (
	"encoding/csv"
	"os"
	// "time"

	"github.com/iammuuo/r2todo/configs"
	"github.com/iammuuo/r2todo/internal/models"
)

type TodoController struct {
	Cfg    *configs.Config
	Todos  []models.Todo
	reader *csv.Reader
	writer *csv.Writer
}

const name string = "hello.csv"

func (t *TodoController) ListTodos(showComplete bool, showOverDue bool) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}

	t.reader = csv.NewReader(file)

	records, err := t.reader.ReadAll()
	if err != nil {
		return err
	}

	for _, value := range records {
		todo, err := models.DeserializeTodo(value)
		if err != nil {
			return err
		}

		if !showComplete && !todo.Completed {
			continue
		}
		//
		// if !showOverDue && !todo.DueDate.After(time.Now()) {
		// 	continue
		// }
		t.Todos = append(t.Todos, *todo)
	}
	models.DisplayTodos(&t.Todos)

	return nil
}

func (t *TodoController) CreateTodo(todoDescription string) error {

	file, err := os.OpenFile(name, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	t.reader = csv.NewReader(file)

	records, err := t.reader.ReadAll()
	if err != nil {
		return err
	}

	for _, value := range records {
		todo, err := models.DeserializeTodo(value)
		if err != nil {
			return err
		}
		t.Todos = append(t.Todos, *todo)
	}

	lastID := 0
	if len(t.Todos) > 0 {
		lastID = t.Todos[len(t.Todos)-1].ID
	}

	todo := models.NewTodoItem(lastID+1, todoDescription)

	data := todo.SerializeTodo()
	t.writer = csv.NewWriter(file)
	err = t.writer.Write(data)
	if err != nil {
		return err
	}
	t.writer.Flush()
	return nil
}
