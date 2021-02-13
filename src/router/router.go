package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar um router
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
