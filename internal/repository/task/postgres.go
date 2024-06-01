package task

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllTasks(ctx context.Context) ([]Task, error) {
	sql, _, err := goqu.From("task").ToSQL()
	if err != nil {
		fmt.Print("TO SQL ERROR")
	}

	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var result []Task
	for rows.Next() {
		task, err := mapRowToTask(rows)
		if err != nil {
			// TODO: Handle error
			fmt.Println(err)
		}
		result = append(result, task)
	}

	return result, nil
}

// TODO: should return at least id of the created task
func (r *PostgresRepository) CreateTask(ctx context.Context, task TaskForCreate) (err error) {
	insertSQL, args, _ := goqu.Insert("task").Rows(TaskForCreate{
		Title:       task.Title,
		Description: task.Description,
		Color:       task.Color,
		UserID:      task.UserID,
	}).Returning("id").ToSQL()

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	// TODO: handle error
	if err != nil {
		fmt.Println(err)
		return err
	}

	slog.Info("CREATE RESULT", result)
	return nil
}

// TODO: use `uuid` type for taskID instead of `string`
func (r *PostgresRepository) DeleteTask(ctx context.Context, taskID string) (err error) {
	insertSQL, args, _ := goqu.Delete("task").Where(goqu.Ex{"id": taskID}).Returning("id").ToSQL()

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	// TODO: handle error
	if err != nil {
		fmt.Println(err)
		return err
	}

	slog.Info("DELETED TASKS ID", result)
	return nil
}
