package postgres

import (
	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
	"github.com/stanislav-zeman/go-ddd/internal/domain/repository"
)

var _ repository.TodoRepository = &TodoRepository{}

type TodoRepository struct{}

// GetTodo implements repository.TodoRepository.
func (t *TodoRepository) GetTodo(id int) (todo entity.Todo, err error) {
	panic("unimplemented")
}

// UpsertTodo implements repository.TodoRepository.
func (t *TodoRepository) UpsertTodo(todo entity.Todo) (err error) {
	panic("unimplemented")
}
