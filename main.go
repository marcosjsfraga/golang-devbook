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

	// Criar chave usada para assinar o token
	// tools.CreateSecretKey()

	// Carregar rotas
	r := router.Gerar()

	fmt.Println("------------------------------")
	fmt.Println("Rodando Server API Devboook...")
	fmt.Println("------------------------------")

	// Executar HTTP Server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
