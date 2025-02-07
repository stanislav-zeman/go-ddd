package query

import "github.com/stanislav-zeman/go-ddd/internal/domain/entity"

type GetTodoQuery struct {
	ID int
}

type GetTodoQueryResult struct {
	Todo entity.Todo
}
