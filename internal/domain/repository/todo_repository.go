package repository

import (
	"errors"

	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
)

var ErrTodoNotFound = errors.New("no such todo found")

type TodoRepository interface {
	GetTodo(id int) (entity entity.Todo, err error)
	UpsertTodo(todo entity.Todo) (err error)
}
