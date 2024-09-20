package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

func init() {
	var err error
	DBPool, err = pgxpool.Connect(context.Background(), "postgresql://postgres:jSOqSEraSTOgPLWCaZTYcmpTHfqpvTKI@postgres.railway.internal:5432/railway")

	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v\n", err)
	}
}
