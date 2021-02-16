package tools

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// CheckError retorna erros
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// CreateSecretKey criar uma chave segredo para criptografia
func CreateSecretKey() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)

	fmt.Println(stringBase64)
}
