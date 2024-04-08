package external

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func ConnectToDatabase() (*pgxpool.Pool, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	dbpool, err := pgxpool.New(ctx, "postgres://postgres:testpass@127.0.0.1:5432/jobs_db")
	if err != nil {
		log.Println(err)
		cancel()
		return nil, nil, err
	}
	//defer dbpool.Close()
	return dbpool, cancel, nil
}
