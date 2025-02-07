package interfaces

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/command"
	"github.com/stanislav-zeman/go-ddd/internal/application/query"
)

type TodoService interface {
	GetTodo(query *query.GetTodoQuery) (result *query.GetTodoQueryResult, err error)
	UpsertTodo(command *command.UpsertTodoCommand) (result *command.UpsertTodoCommandResult, err error)
}
