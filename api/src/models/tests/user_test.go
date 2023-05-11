package models

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

	// t.Run("sad path:  name is empty", func(t *testing.T) {
	// 	mockUser(&user, "", " nick ", " email ", "pwd")
	// 	erro := user.Prepare(true)
	// 	assert.Equal(t, erro.Error(), "o nome é obrigatório")
	// })

	// t.Run("sad path:  nick is empty", func(t *testing.T) {
	// 	mockUser(&user, "name", "", " email ", "pwd")
	// 	erro := user.Prepare(true)
	// 	assert.Equal(t, erro.Error(), "o nick é obrigatório")
	// })

	// t.Run("sad path:  email is empty", func(t *testing.T) {
	// 	mockUser(&user, "name", "nick", "", "pwd")
	// 	erro := user.Prepare(true)
	// 	assert.Equal(t, erro.Error(), "o email é obrigatório")
	// })

	// t.Run("sad path:  email is invalid", func(t *testing.T) {
	// 	mockUser(&user, "name", "nick", "emailXXXXXXXXXX", "pwd")
	// 	erro := user.Prepare(true)
	// 	assert.Equal(t, erro.Error(), "o email inserido está em um formato inválido")
	// })

	// t.Run("sad path:  pwd is empty", func(t *testing.T) {
	// 	mockUser(&user, "name", "nick", "email@gmail.com", "")
	// 	erro := user.Prepare(true)
	// 	assert.Equal(t, erro.Error(), "a senha é obrigatória")
	// })
}
