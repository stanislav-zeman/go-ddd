package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stanislav-zeman/go-ddd/internal/api/rest/dto/request"
	"github.com/stanislav-zeman/go-ddd/internal/api/rest/dto/response"
	"github.com/stanislav-zeman/go-ddd/internal/application/interfaces"
	"github.com/stanislav-zeman/go-ddd/internal/domain/repository"
)

type TodoController struct {
	service interfaces.TodoService
}

func NewTodoController(service interfaces.TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

func (td *TodoController) MountControllers(g *echo.Group) {
	g.GET("/todos/:todoId", td.GetTodoController)
	g.POST("/todos", td.UpsertTodoController)
}

func (td *TodoController) GetTodoController(c echo.Context) error {
	var req request.GetTodoRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorResponse{
				Error: fmt.Errorf("failed parsing request: %w", err),
			},
		)
		return nil
	}

	q := req.ToGetTodoQuery()
	result, err := td.service.GetTodo(q)
	if errors.Is(err, repository.ErrTodoNotFound) {
		c.JSON(http.StatusNotFound,
			response.ErrorResponse{
				Error: repository.ErrTodoNotFound,
			},
		)
		return nil
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Error: errors.New("failed fetching todo data"),
			},
		)
		return nil
	}

	res := response.NewGetTodoResponseFromQueryResult(result)

	c.JSON(http.StatusOK, res)

	return nil
}

func (td *TodoController) UpsertTodoController(c echo.Context) error {
	var req request.UpsertTodoRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorResponse{
				Error: fmt.Errorf("failed parsing request: %w", err),
			},
		)
		return nil
	}

	cmnd := req.ToUpsertTodoCommand()
	result, err := td.service.UpsertTodo(cmnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.ErrorResponse{
				Error: errors.New("failed fetching todo data"),
			},
		)
		return nil
	}

	res := response.NewUpsertTodoResponseFromCommandResult(result)

	c.JSON(http.StatusOK, res)

	return nil
}
