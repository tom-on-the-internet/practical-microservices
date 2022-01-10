package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

func (a *App) InitDB() {
	pool, err := pgxpool.Connect(context.Background(), a.env.dbConnStr)
	if err != nil {
		panic("cannot init DB" + err.Error())
	}

	a.db = db{pool: pool}
}
