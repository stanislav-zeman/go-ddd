package request

import "github.com/stanislav-zeman/go-ddd/internal/application/query"

type GetTodoRequest struct {
	TodoID int `form:"todoId"`
}

func (gtr *GetTodoRequest) ToGetTodoQuery() *query.GetTodoQuery {
	return &query.GetTodoQuery{
		ID: gtr.TodoID,
	}
}
