package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(pool *pgxpool.Pool) error {
	ctx := context.Background()

	query := `
		CREATE TABLE IF NOT EXISTS greetings (
			id SERIAL PRIMARY KEY,
			message TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		INSERT INTO greetings (message) 
		VALUES ('Hello World from PostgreSQL!')
		ON CONFLICT DO NOTHING;
	`

	_, err := pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	log.Println("Database migration completed")
	return nil
}
