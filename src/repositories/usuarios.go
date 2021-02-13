package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Repo representa um repositório de usuários
type Repo struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria novo repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Repo {
	return &Repo{db}
}

// Criar insere um usuario no DB
func (repo Repo) Criar(usuario models.Usuario) (uint64, error) {

	statement, err := repo.db.Prepare(
		`INSERT INTO usuarios 
		(nome, nick, email, senha) VALUES 
		($1, $2, $3, $4)`)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	stmResult, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	fmt.Println(stmResult)

	ultimoIDInserido := 0

	return uint64(ultimoIDInserido), nil
}

// Buscar pesquisa usuários
func (repo Repo) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, err := repo.db.Query(
		`SELECT id, nome, nick, email
		 FROM usuarios
		 WHERE nome LIKE $1 OR nick LIKE $2 `,
		nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID busca um usuário por ID
func (repo Repo) BuscarPorID(ID uint64) (models.Usuario, error) {

	linhas, err := repo.db.Query(
		`SELECT id, nome, nick, email
		 FROM usuarios
		 WHERE id = $1`,
		ID,
	)

	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	for linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil

}

// Alterar insere um usuario no DB
func (repo Repo) Alterar(ID uint64, usuario models.Usuario) error {

	statement, err := repo.db.Prepare(
		`UPDATE usuarios 
		SET nome=$1, nick=$2, email=$3, senha=$4 
		WHERE id=$5`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Senha,
		ID); err != nil {
		return err
	}

	return nil
}

// Deletar apaga um usuario no DB
func (repo Repo) Deletar(ID uint64) error {

	statement, err := repo.db.Prepare(`DELETE FROM usuarios WHERE id=$1`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
