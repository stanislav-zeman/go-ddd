package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stanislav-zeman/go-ddd/internal/api/rest/controller"
	"github.com/stanislav-zeman/go-ddd/internal/application/services"
	"github.com/stanislav-zeman/go-ddd/internal/infrastructure/persistence/postgres"
)

func main() {
	tr := postgres.NewTodoRepository()
	ts := services.NewTodoService(tr)
	tc := controller.NewTodoController(ts)

	e := echo.New()
	v1 := e.Group("/api/v1")
	tc.MountControllers(v1)

	err := e.Start(":8080")
	if err != nil {
		slog.Error("failed running server",
			"error", err,
		)
		os.Exit(1)
	}
}
