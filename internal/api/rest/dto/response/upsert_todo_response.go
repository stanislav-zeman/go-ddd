package response

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/command"
)

type UpsertTodoResponse struct{}

func NewUpsertTodoResponseFromCommandResult(_ *command.UpsertTodoCommandResult) *UpsertTodoResponse {
	return &UpsertTodoResponse{}
}
