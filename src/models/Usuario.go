package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario - entidade de usuario
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `email:"nome,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

// Preparar prepara o usuario
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome deve ser informado")
	}

	if usuario.Nick == "" {
		return errors.New("O nick deve ser informado")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail deve ser informado")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("O e-mail não é válido")
	}

	if etapa == "inclusao" && usuario.Senha == "" {
		return errors.New("A senha deve ser informado")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "inclusao" {
		senhaComHash, err := security.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
