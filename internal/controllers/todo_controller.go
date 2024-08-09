package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/iammuuo/r2todo/configs"
	"github.com/iammuuo/r2todo/internal/models"
	"os"
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
	defer file.Close()
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

		if !showComplete || !todo.Completed {
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
	defer file.Close()
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

func (t *TodoController) ToggleTodoStatus(id int) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	updated := false
	t.reader = csv.NewReader(file)

	records, err := t.reader.ReadAll()
	if err != nil {
		return err
	}

	for i, value := range records {
		todo, err := models.DeserializeTodo(value)
		if err != nil {
			return err
		}
		if todo.ID == id {
			updated = true
			todo.Completed = !todo.Completed
			records[i] = todo.SerializeTodo()
		}
	}

	if !updated {
		return fmt.Errorf("Todo with id %d not found!\n", id)
	}

	file.Seek(0, 0)
	t.writer = csv.NewWriter(file)
	t.writer.WriteAll(records)
	t.writer.Flush()

	return nil
}

func (t *TodoController) Delete(id int, deleteAll bool) error {
	if deleteAll {
		file, err := os.OpenFile(name, os.O_TRUNC|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil {
			return err
		}
		return nil
	}

	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	t.reader = csv.NewReader(file)

	records, err := t.reader.ReadAll()
	if err != nil {
		return err
	}

	founIndex := 0
	found := false
	for i, value := range records {
		todo, err := models.DeserializeTodo(value)
		if err != nil {
			return err
		}
		if todo.ID == id {
			found = true
			founIndex = i
		}
	}

	if !found {
		return fmt.Errorf("Todo with id %d not found!\n", id)
	}

	records = append(records[:founIndex], records[founIndex+1:]...)

	file.Truncate(0)
	file.Seek(0, 0)
	t.writer = csv.NewWriter(file)
	t.writer.WriteAll(records)
	t.writer.Flush()

	return nil
}
