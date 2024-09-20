package models

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

type PagamentoPendente struct {
	Nome string `json:"nome"`
	Ano  int    `json:"ano"`
	Mes  int    `json:"mes"`
	Pago bool   `json:"pago"`
}
