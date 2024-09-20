package controllers

import (
	"encoding/json"
	"net/http"

	"strconv"
	"strings"

	"api-finpay/src/db"
	"api-finpay/src/models"

	"github.com/gorilla/mux"
)

func CreateNewAluno(w http.ResponseWriter, r *http.Request) {
	var aluno models.Aluno
	err := json.NewDecoder(r.Body).Decode(&aluno)
	if err != nil || aluno.Nome == "" || aluno.TurmaID == 0 {
		http.Error(w, "Nome é obrigatório", http.StatusBadRequest)
		return
	}

	_, err = db.DBPool.Exec(r.Context(), "INSERT INTO alunos(nome, turma_id, pagamento) VALUES($1, $2, $3)", aluno.Nome, aluno.TurmaID, aluno.Pagamento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllAlunos(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DBPool.Query(r.Context(), "SELECT * FROM alunos")
	var alunos []models.Aluno

	for rows.Next() {
		var aluno models.Aluno
		rows.Scan(&aluno.ID, &aluno.Nome, &aluno.TurmaID)
		alunos = append(alunos, aluno)
	}

	json.NewEncoder(w).Encode(alunos)
}

func GetAlunoByNome(w http.ResponseWriter, r *http.Request) {
	nome := mux.Vars(r)["nome"]
	query := "SELECT id, name, turma_id FROM alunos WHERE nome ILIKE '%' || $1 || '%'"
	rows, err := db.DBPool.Query(r.Context(), query, nome)
	if err != nil {
		http.Error(w, "Erro ao buscar alunos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var alunos []models.Aluno
	for rows.Next() {
		var aluno models.Aluno
		if err := rows.Scan(&aluno.ID, &aluno.Nome, &aluno.TurmaID); err != nil {
			http.Error(w, "Erro ao ler dados do aluno", http.StatusInternalServerError)
			return
		}
		alunos = append(alunos, aluno)
	}

	if len(alunos) == 0 {
		http.Error(w, "Nenhum aluno encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(alunos)
}

func GetAlunosByTurmaID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	rows, _ := db.DBPool.Query(r.Context(), "SELECT * FROM alunos WHERE turma_id=$1", id)
	var alunos []models.Aluno

	for rows.Next() {
		var aluno models.Aluno
		rows.Scan(&aluno.ID, &aluno.Nome, &aluno.TurmaID, &aluno.Pagamento)
		alunos = append(alunos, aluno)
	}

	json.NewEncoder(w).Encode(alunos)
}

func DeleteAlunos(w http.ResponseWriter, r *http.Request) {
	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Erro ao decodificar os IDs: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Gera uma string de placeholders para a query
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "$" + strconv.Itoa(i+1) // Cria os placeholders como $1, $2, ...
	}

	query := "DELETE FROM alunos WHERE id IN (" + strings.Join(placeholders, ", ") + ")"
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	_, err = db.DBPool.Exec(r.Context(), query, args...)
	if err != nil {
		http.Error(w, "Erro ao deletar alunos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
