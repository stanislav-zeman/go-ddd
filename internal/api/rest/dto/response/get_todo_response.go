package response

import (
	"github.com/stanislav-zeman/go-ddd/internal/application/query"
	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
)

type GetTodoResponse struct {
	Todo entity.Todo
}

func NewGetTodoResponseFromQueryResult(result *query.GetTodoQueryResult) *GetTodoResponse {
	return &GetTodoResponse{
		Todo: result.Todo,
	}
}
