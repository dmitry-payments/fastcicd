package database

import (
	"context"
	"fastcicd/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GreetingRepo struct {
	pool *pgxpool.Pool
}

func NewGreetingRepo(pool *pgxpool.Pool) *GreetingRepo {
	return &GreetingRepo{pool: pool}
}

func (r *GreetingRepo) GetGreetings(ctx context.Context) ([]models.Greeting, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, message, created_at FROM greetings ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var greetings []models.Greeting
	for rows.Next() {
		var g models.Greeting
		if err := rows.Scan(&g.ID, &g.Message, &g.CreatedAt); err != nil {
			return nil, err
		}
		greetings = append(greetings, g)
	}
	return greetings, rows.Err()
}

func (r *GreetingRepo) AddGreeting(ctx context.Context, message string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO greetings (message) VALUES ($1)", message)
	return err
}
