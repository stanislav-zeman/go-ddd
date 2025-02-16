package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stanislav-zeman/go-ddd/internal/api/rest/controller"
	"github.com/stanislav-zeman/go-ddd/internal/application/services"
	"github.com/stanislav-zeman/go-ddd/internal/infrastructure/persistence/postgres"
)

func main() {
	err := runApp()
	if err != nil {
		slog.Error("failed running app",
			"error", err,
		)
		os.Exit(1)
	}
}

func runApp() error {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		err = fmt.Errorf("failed creating database connection pool: %w", err)
		return err
	}

	defer db.Close()

	tr := postgres.NewTodoRepository(db)
	ts := services.NewTodoService(tr)
	tc := controller.NewTodoController(ts)

	e := echo.New()
	v1 := e.Group("/api/v1")
	tc.MountControllers(v1)

	err = e.Start(":8080")
	if err != nil {
		err = fmt.Errorf("failed running server: %w", err)
		return err
	}

	return nil
}
