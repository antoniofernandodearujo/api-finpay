package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"api-finpay/src/db"
	"api-finpay/src/models"

	"github.com/gorilla/mux"
)

func UpdatePagamentos(w http.ResponseWriter, r *http.Request) {
	alunoID := mux.Vars(r)["id"]
	ano := 2024 // Aqui você pode receber o ano como um parâmetro, se necessário
	mes := 9    // Aqui você pode receber o mês como um parâmetro, se necessário

	_, err := db.DBPool.Exec(context.Background(), "UPDATE pagamentos SET pago = TRUE WHERE aluno_fk = $1 AND ano = $2 AND mes = $3", alunoID, ano, mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func FetchPagamentosPendentes(w http.ResponseWriter, r *http.Request) {
	ano := 2024 // Aqui você pode pegar o ano do dispositivo ou como parâmetro
	mes := 9    // Aqui você pode pegar o mês do dispositivo ou como parâmetro

	rows, err := db.DBPool.Query(context.Background(), `
		SELECT alunos.nome, pagamentos.ano, pagamentos.mes, pagamentos.pago
		FROM pagamentos
		JOIN alunos ON pagamentos.aluno_fk = alunos.id
		WHERE pagamentos.ano = $1 AND pagamentos.mes = $2 AND pagamentos.pago = FALSE`, ano, mes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pagamentosPendentes []models.PagamentoPendente
	for rows.Next() {
		var pagamento models.PagamentoPendente
		err = rows.Scan(&pagamento.Nome, &pagamento.Ano, &pagamento.Mes, &pagamento.Pago)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pagamentosPendentes = append(pagamentosPendentes, pagamento)
	}

	json.NewEncoder(w).Encode(pagamentosPendentes)
}
