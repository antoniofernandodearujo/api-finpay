package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

func init() {
	var err error
	DBPool, err = pgxpool.Connect(context.Background(), "postgres://admin:admin@localhost:5432/finpay")

	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: %v\n", err)
	}
}
