package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) validate(create bool) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório")
	}

	if user.Nick == "" {
		return errors.New("o nick é obrigatório")
	}

	if user.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("o email inserido está em um formato inválido")
	}

	if user.Password == "" && create {
		return errors.New("A senha é obrigatória.")
	}

	return nil
}

func (user *User) validateLogin() error {

	if user.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if user.Password == "" {
		return errors.New("a senha é obrigatória")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

// Hash_passowrd vai chamar metodos de hash de biblioteca terceira
func (user *User) Hash_passowrd() error {
	var erro error = nil
	user.Password, erro = security.Hash(user.Password)
	if erro != nil {
		return errors.New("lamento, mas no momento não será possível criar o usuário")
	}
	return erro
}

// Prepare vai chamar metodos de validar e formatar usuario
func (user *User) Prepare(create ...bool) error {
	var param bool = false
	if create != nil {
		param = create[0]
	}
	if erro := user.validate(param); erro != nil {
		return erro
	}

	user.format()
	return nil
}

// PrepareForLogin vai chamar metodos de validar e formatar o login
func (user *User) PrepareForLogin() error {
	if erro := user.validateLogin(); erro != nil {
		return erro
	}

	user.format()
	return nil
}
