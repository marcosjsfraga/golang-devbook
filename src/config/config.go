package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco strgin de coenxão com o Postgres
	StringConexaoBanco = ""
	// Porta APIS
	Porta = 0
)

// Carregar inicializar variaveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_POST"))
	if err != nil {
		Porta = 9000
	}

}
