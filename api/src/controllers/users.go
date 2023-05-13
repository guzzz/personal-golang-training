package controllers

import (
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

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

	var prepareForCreation bool = true
	if erro = user.Prepare(prepareForCreation); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.HashPassowrd(); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	result, erro := repository.List(nameOrNick)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, result)

}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	result, erro := repository.Retrieve(userId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if result.ID == 0 {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if userId != tokenUserId {
		var msgForbidden error = errors.New("essa ação não é permitida")
		responses.Error(w, http.StatusForbidden, msgForbidden)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var data models.User
	if erro = json.Unmarshal(body, &data); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = data.Prepare(); erro != nil {
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
	erro = repository.Update(userId, data)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, data)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if userId != tokenUserId {
		var msgForbidden error = errors.New("essa ação não é permitida")
		responses.Error(w, http.StatusForbidden, msgForbidden)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	rows_deleted, erro := repository.Delete(userId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if rows_deleted > 0 {
		responses.JSON(w, http.StatusNoContent, nil)
	} else {
		responses.JSON(w, http.StatusNotFound, nil)
	}

}

func FollowUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if userId == tokenUserId {
		var msgForbidden error = errors.New("Essa ação não é permitida")
		responses.Error(w, http.StatusForbidden, msgForbidden)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	erro = repository.FollowUser(userId, tokenUserId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	msg := make(map[string]string)
	msg["msg"] = "Usuário seguido com sucesso!"
	responses.JSON(w, http.StatusOK, msg)

}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	unfollowed, erro := repository.UnfollowUser(userId, tokenUserId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	msg := make(map[string]string)
	if unfollowed {
		msg["msg"] = "Você deixou de seguir o usuário!"
	} else {
		msg["msg"] = "Não foi possível deixar de seguir este usuário."
	}
	responses.JSON(w, http.StatusOK, msg)

}

func GetFollowers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	result, erro := repository.ListFollowers(userId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
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
	result, erro := repository.ListFollowing(userId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var passwordForm models.Password
	if erro = json.Unmarshal(body, &passwordForm); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	user, erro := repository.RetrieveUserForUpdatePassword(tokenUserId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if user.ID == 0 {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}

	if erro = security.CompareHash(user.Password, passwordForm.Actual); erro != nil {
		erro = errors.New("a senha atual está incorreta")
		responses.Error(w, http.StatusUnauthorized, erro)
		return
	}

	user.Password = passwordForm.NewOne
	if erro = user.HashPassowrd(); erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	erro = repository.UpdatePassword(tokenUserId, user.Password)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	msg := make(map[string]string)
	msg["msg"] = "Senha alterada com sucesso!"
	responses.JSON(w, http.StatusOK, msg)

}
