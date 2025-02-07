package command

import (
	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
)

type UpsertTodoCommand struct {
	Todo entity.Todo
}

type UpsertTodoCommandResult struct{}
