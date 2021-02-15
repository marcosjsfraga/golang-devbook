package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login efetua a validacao das credenciais do usuario
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	// Abrir conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	// Fechar a coenxão no final da chamada desta função
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioSalvoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioSalvoBanco.Senha, usuario.Senha); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("Usuário autorizado"))
}
