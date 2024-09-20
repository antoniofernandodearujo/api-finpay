package routes

import (
	"api-finpay/src/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {

	// Aluno routes
	// cria um novo aluno
	r.HandleFunc("/alunos", controllers.CreateNewAluno).Methods("POST")
	// busca todos os alunos
	r.HandleFunc("/alunos", controllers.GetAllAlunos).Methods("GET")
	// busca um aluno por id
	r.HandleFunc("/alunos/nome/{name}", controllers.GetAlunoByNome).Methods("GET")
	// busca alunos por turma
	r.HandleFunc("/alunos/turma/{id:[0-9]+}", controllers.GetAlunosByTurmaID).Methods("GET")
	// atualiza um aluno
	r.HandleFunc("/alunos", controllers.DeleteAlunos).Methods("DELETE")

	// Pagamento routes
	// atualiza pagamentos de um aluno
	r.HandleFunc("/alunos/{id:[0-9]+}/pagamentos/{ano}/{mes}", controllers.UpdatePagamentos).Methods("PUT")
	// busca pagamentos pendentes
	r.HandleFunc("/pagamentos", controllers.FetchPagamentosPendentes).Methods("GET")

	// Turma routes
	// busca todas as turmas
	r.HandleFunc("/turmas", controllers.FetchTurmas).Methods("GET")
}
