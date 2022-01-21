package app

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

type messageStore struct {
	pool *pgxpool.Pool
}

type message struct {
	id       uuid.UUID
	msgType  string
	data     string
	metadata string
}

func (a *App) InitDBs() {
	pool, err := pgxpool.Connect(context.Background(), a.env.dbConnStr)
	if err != nil {
		panic("cannot init DB" + err.Error())
	}

	a.db = db{pool: pool}

	pool, err = pgxpool.Connect(context.Background(), a.env.msgStoreConnStr)
	if err != nil {
		panic("cannot init MESSAGE_STORE_DB" + err.Error())
	}

	a.msgStore = messageStore{pool: pool}
}

func newMessage(msgType string, data string, metadata string) (message, error) {
	if msgType == "" {
		return message{}, errors.New("message cannot be blank")
	}

	msg := message{
		id:       uuid.New(),
		msgType:  msgType,
		data:     data,
		metadata: metadata,
	}

	return msg, nil
}
