package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stanislav-zeman/go-ddd/internal/domain/entity"
	"github.com/stanislav-zeman/go-ddd/internal/domain/repository"
)

var _ repository.TodoRepository = &TodoRepository{}

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

// GetTodo implements repository.TodoRepository.
func (tr *TodoRepository) GetTodo(ctx context.Context, id int) (todo entity.Todo, err error) {
	rows, err := tr.db.Query(ctx, `
		SELECT * FROM todos WHERE id=$1
	`, id,
	)
	if err != nil {
		err = fmt.Errorf("failed querying database: %w", err)
		return
	}

	todo, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Todo])
	if errors.Is(err, pgx.ErrNoRows) {
		err = repository.ErrTodoNotFound
		return
	}

	if err != nil {
		err = fmt.Errorf("failed collecting row to struct: %w", err)
		return
	}

	return
}

// UpsertTodo implements repository.TodoRepository.
func (tr *TodoRepository) UpsertTodo(ctx context.Context, todo entity.Todo) (err error) {
	t, err := tr.db.Exec(ctx, `
		INSERT INTO todos (id, content, state)
		VALUES ($1, $2, $3)
		ON CONFLICT (id)
		DO UPDATE SET
			content = $2,
			state = $3
		`,
		todo.ID,
		todo.Content,
		todo.State,
	)
	if err != nil {
		err = fmt.Errorf("failed querying database: %w", err)
		return
	}

	if t.RowsAffected() != 1 {
		err = errors.New("no rows were affected")
		return
	}

	return
}
