package controllers

import (
	"github.com/iammuuo/r2todo/configs"
	"github.com/iammuuo/r2todo/internal/models"
)

type TodoController struct {
	Cfg   *configs.Config
	Todos []models.Todo
}

func (t *TodoController) ListTodos(showComplete bool, showOverDue bool) error {
	return nil
}
