package models

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodoItem(0, "Hello there")

	if todo.Description != "Hello there" {
		t.Log("Failed to create a todo correctly")
		t.Fail()
	}

}

func TestTodoSerializer(t *testing.T) {
	todoItem := Todo{
		1, "Hello there", false, time.Now(), time.Now(),
	}

	serialized := todoItem.SerializeTodo()

	if len(serialized) < 5 {
		t.Fatalf(
			"Error: expexted serialized to have five elements instead got %d with %v values\n",
			len(serialized),
			serialized,
		)
	}

}

func TestTodoDeserializer(t *testing.T) {

	now := time.Now()
	todoItemSerialized := []string{
		"1",
		"Hello there",
		"false",
		now.Format(time.RFC3339) + "Z",
		now.Format(time.RFC3339) + "Z",
	}

	todo, err := DeserializeTodo(todoItemSerialized)
	if err != nil {
		t.Fatalf("Failed to Deserialize todo item with meesage %s\n",
			err.Error(),
		)
	}

	if todo.ID != 1 {
		t.Log("Failed to parse ID")
		t.Fail()
	}

	if todo.Description != "Hello there" {
		t.Log("Failed to parse todo description")
		t.Fail()
	}

	if todo.Completed {
		t.Log("Failed to parse todo status")
		t.Fail()

	}
}

func TestTodoDisplay(t *testing.T) {
	todos := []Todo{
		{
			1, "Hello there", false, time.Now(), time.Now(),
		},
		{
			2, "Hello there again", false, time.Now(), time.Now(),
		},
		{
			3, "Hello there again", false, time.Now(), time.Now(),
		},
	}

	DisplayTodos(&todos)
}
