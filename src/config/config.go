package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco strign de conexão com o Postgres
	StringConexaoBanco = ""
	// Porta APIS
	Porta = 0

	//SecretKey chave usada para assinar o token
	SecretKey []byte
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

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
