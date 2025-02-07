package request

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/command"
	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
	valueobject "github.com/stanislav-zeman/go-ddd/internal/domain/value_object"
)

type UpsertTodoRequest struct {
	TodoID  int               `form:"todoId"`
	Content string            `json:"content"`
	State   valueobject.State `json:"state"`
}

func (gtr *UpsertTodoRequest) ToUpsertTodoCommand() *command.UpsertTodoCommand {
	return &command.UpsertTodoCommand{
		Todo: entity.Todo{
			ID:      gtr.TodoID,
			Content: gtr.Content,
			State:   gtr.State,
		},
	}
}
