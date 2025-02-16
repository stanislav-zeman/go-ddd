package repository

import (
	"context"
	"errors"

	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
)

var ErrTodoNotFound = errors.New("no such todo found")

type TodoRepository interface {
	GetTodo(ctx context.Context, id int) (entity entity.Todo, err error)
	UpsertTodo(ctx context.Context, todo entity.Todo) (err error)
}
