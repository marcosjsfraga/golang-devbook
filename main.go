package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Carregar variáveis de ambiente
	config.Carregar()

	// Carregar rotas
	r := router.Gerar()

	fmt.Println("Rodando Server API Devboook...")

	// Executar HTTP Server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
