package response

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/command"
)

type UpsertTodoResponse struct{}

func NewUpsertTodoResponseFromCommandResult(result *command.UpsertTodoCommandResult) *UpsertTodoResponse {
	return &UpsertTodoResponse{}
}
