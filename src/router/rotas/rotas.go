package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota Estrutura de rota
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar  coloca as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	// Cria o objeto "rotas" e adiciona a rota de usuarios
	rotas := rotasUsuarios
	// Adiciona a rota login ao objeto "rotas"
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
