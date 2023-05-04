package controllers

import (
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.PrepareForLogin(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	result, erro := repository.RetrieveUserForLogin(user.Email)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = security.CompareHash(result.Password, user.Password); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	token, erro := autentication.CreateToken(result.ID)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	response := make(map[string]string)
	response["token"] = token
	responses.JSON(w, http.StatusOK, response)
}
