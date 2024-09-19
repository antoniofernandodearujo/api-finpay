package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

var dbpool *pgxpool.Pool

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	TurmaID   int    `json:"turma_id"`
	Pagamento bool   `json:"pagamento"`
}

type Turma struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func initDB() {
	var err error

	// conexão com o banco de dados
	dbpool, err = pgxpool.Connect(context.Background(), "postgres://admin:admin@localhost:5432/finpay")

	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: %v\n", err)
	}

	fmt.Println("Conexão com o banco de dados realizada com sucesso!")
}

func main() {
	initDB()
	defer dbpool.Close()

	router := mux.NewRouter()

	http.ListenAndServe(":8080", router)
}

func getTurmas(w http.ResponseWriter, r *http.Request) {
	rows, _ := dbpool.Query(context.Background(), "SELECT id, nome FROM turmas")
	var turmas []Turma

	for rows.Next() {
		var t Turma
		rows.Scan(&t.ID, &t.Nome)
		turmas = append(turmas, t)
	}

	json.NewEncoder(w).Encode(turmas)
}

func createAluno(w http.ResponseWriter, r *http.Request) {
	var aluno Aluno
	json.NewDecoder(r.Body).Decode(&aluno)
	_, err := dbpool.Exec(context.Background(), "INSERT INTO alunos(nome, turma_id, pagamento) VALUES($1, $2, $3)", aluno.Nome, aluno.TurmaID, aluno.Pagamento)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
