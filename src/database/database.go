package database

import (
	"api/src/tools"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Conectar coencta no DB
func Conectar() (*sql.DB, error) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		5432,
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	db, err := sql.Open("postgres", psqlconn)
	tools.CheckError(err)

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
