package models_test

import (
	"api/src/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var user models.User

func mockUser(user *models.User, name, nick, email, pwd string) models.User {
	*user = models.User{
		ID:        1,
		Name:      name,
		Nick:      nick,
		Email:     email,
		Password:  pwd,
		CreatedAt: time.Time{},
	}
	return *user
}

func TestPrepare(t *testing.T) {

	t.Run("Success case - should return nil", func(t *testing.T) {
		user = mockUser(&user, "name", "nick", "email@gmail.com", "pwd")
		erro := user.Prepare(true)
		assert.Equal(t, erro, nil)
	})

	t.Run("Error cases - should return some error", func(t *testing.T) {
		tests := []struct {
			user models.User
			erro string
		}{
			{
				mockUser(&user, "", " nick ", " email ", "pwd"),
				"o nome é obrigatório",
			},
			{
				mockUser(&user, "name", "", " email ", "pwd"),
				"o nick é obrigatório",
			},
			{
				mockUser(&user, "name", "nick", "", "pwd"),
				"o email é obrigatório",
			},
			{
				mockUser(&user, "name", "nick", "emailXXXXXXXXXX", "pwd"),
				"o email inserido está em um formato inválido",
			},
			{
				mockUser(&user, "name", "nick", "email@gmail.com", ""),
				"a senha é obrigatória",
			},
		}

		for _, tt := range tests {
			user = tt.user
			erro := user.Prepare(true)
			assert.Equal(t, tt.erro, erro.Error())
		}
	})

}

func TestPrepareForLogin(t *testing.T) {

	t.Run("Success case - should return nil", func(t *testing.T) {
		user = mockUser(&user, "name", "nick", "email@gmail.com", "pwd")
		erro := user.PrepareForLogin()
		assert.Equal(t, erro, nil)
	})

	t.Run("Errors case - should return some error", func(t *testing.T) {
		tests := []struct {
			user models.User
			erro string
		}{
			{
				mockUser(&user, "", " nick ", "", "pwd"),
				"o email é obrigatório",
			},
			{
				mockUser(&user, "name", "", "email@email.com", ""),
				"a senha é obrigatória",
			},
		}

		for _, tt := range tests {
			user = tt.user
			erro := user.PrepareForLogin()
			assert.Equal(t, tt.erro, erro.Error())
		}
	})

}

func TestHashPassword(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		user = mockUser(&user, "name", "nick", "email@gmail.com", "pwd")
		erro := user.HashPassowrd()
		assert.Equal(t, erro, nil)
	})

}
