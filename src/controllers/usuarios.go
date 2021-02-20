package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario cria um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("inclusao"); err != nil {
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
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuario busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Ler as variáveis da URL
	parametros := mux.Vars(r)

	// Pegar ID do usuário das variaveis da URL
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if usuario.Nome != "" {
		response.JSON(w, http.StatusOK, usuario)
	} else {
		response.JSON(w, http.StatusNotFound, usuario)
	}
}

// BuscarUsuarios busca todos usuário
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Conectar()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)

	usuarios, err := repositorio.Buscar(nomeOuNick)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, usuarios)
}

// AtualizarUsuario atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDNoToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDNoToken {
		response.Error(w, http.StatusForbidden, errors.New("You can't update another user"))
		return
	}

	// Ler as variáveis do cropo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Atribuir as variáveis do cirpo da requisição ao struct de usuario
	var usuario models.Usuario
	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("alteracao"); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	err = repositorio.Alterar(usuarioID, usuario)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

// DeletarUsuario deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDNoToken, err := auth.ExtrairUsuarioID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIDNoToken {
		response.Error(w, http.StatusForbidden, errors.New("You can't delete another user"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	err = repositorio.Deletar(usuarioID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
