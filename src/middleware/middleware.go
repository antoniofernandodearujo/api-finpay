package middleware

import (
	"log"
	"net/http"

	"api-finpay/src/db"
)

func DatabaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se a conexão com o banco de dados está ativa
		if err := db.DBPool.Ping(r.Context()); err != nil {
			http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
			log.Println("Erro na conexão com o banco de dados:", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
