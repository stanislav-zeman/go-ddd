package services

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/command"
	"github.com/stanislav-zeman/go-ddd/internal/application/interfaces"
	"github.com/stanislav-zeman/go-ddd/internal/application/query"
	"github.com/stanislav-zeman/go-ddd/internal/domain/repository"
)

var _ interfaces.TodoService = &TodoService{}

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

// GetTodo implements interfaces.TodoService.
func (t *TodoService) GetTodo(q *query.GetTodoQuery) (result *query.GetTodoQueryResult, err error) {
	todo, err := t.repository.GetTodo(q.ID)
	if err != nil {
		return
	}

	result = &query.GetTodoQueryResult{
		Todo: todo,
	}

	return
}

// UpsertTodo implements interfaces.TodoService.
func (t *TodoService) UpsertTodo(
	c *command.UpsertTodoCommand,
) (result *command.UpsertTodoCommandResult, err error) {
	err = t.repository.UpsertTodo(c.Todo)
	if err != nil {
		return
	}

	result = &command.UpsertTodoCommandResult{}

	return
}
