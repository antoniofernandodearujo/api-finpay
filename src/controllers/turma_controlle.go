package controllers

import (
	"encoding/json"
	"net/http"

	"api-finpay/src/db"
	"api-finpay/src/models"

	"context"
)

func FetchTurmas(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DBPool.Query(context.Background(), "SELECT * FROM turmas")
	var turmas []models.Turma

	for rows.Next() {
		var t models.Turma
		rows.Scan(&t.ID, &t.Nome)
		turmas = append(turmas, t)
	}

	json.NewEncoder(w).Encode(turmas)
}
