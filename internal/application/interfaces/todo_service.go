package interfaces

import (
	"context"

	"github.com/stanislav-zeman/go-ddd/internal/application/command"
	"github.com/stanislav-zeman/go-ddd/internal/application/query"
)

type TodoService interface {
	GetTodo(ctx context.Context, query *query.GetTodoQuery) (result *query.GetTodoQueryResult, err error)
	UpsertTodo(
		ctx context.Context,
		command *command.UpsertTodoCommand,
	) (result *command.UpsertTodoCommandResult, err error)
}
