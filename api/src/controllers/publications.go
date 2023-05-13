package controllers

import (
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publication
	if erro = json.Unmarshal(body, &publication); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publication.Prepare(); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	publication.AuthorID = tokenUserId
	repository := repositories.PublicationRepository(db)
	publication.ID, erro = repository.Create(publication)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

func ListPublications(w http.ResponseWriter, r *http.Request) {

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

	repository := repositories.PublicationRepository(db)
	result, erro := repository.ListPublicationsThatUserFollows(tokenUserId)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, result)

}

func ListMyPublications(w http.ResponseWriter, r *http.Request) {

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

	repository := repositories.PublicationRepository(db)
	result, erro := repository.ListMyPublications(tokenUserId)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, result)

}

func RetrievePublication(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	result, erro := repository.Retrieve(publicationId)
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

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserId, erro := autentication.ExtractUserIdFromToken(r)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var data models.Publication
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

	repository := repositories.PublicationRepository(db)
	authorID, erro := repository.GetAuthorIdFromPublication(publicationId)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if authorID != tokenUserId {
		var msgForbidden error = errors.New("essa ação não é permitida")
		responses.Error(w, http.StatusForbidden, msgForbidden)
		return
	}

	erro = repository.Update(publicationId, data)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, data)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	authorID, erro := repository.GetAuthorIdFromPublication(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}
	if authorID == 0 {
		var msgPublicationNotFound error = errors.New("publicação não encontrada")
		responses.Error(w, http.StatusNotFound, msgPublicationNotFound)
		return
	}

	if authorID != tokenUserId {
		var msgForbidden error = errors.New("essa ação não é permitida")
		responses.Error(w, http.StatusForbidden, msgForbidden)
		return
	}

	rows_deleted, erro := repository.Delete(publicationId)
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

func Like(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.PublicationRepository(db)
	erro = repository.LikePublication(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	msg := make(map[string]string)
	msg["msg"] = "Publicação curtida com sucesso!"
	responses.JSON(w, http.StatusOK, msg)

}

func RemoveLike(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	publicationId, erro := strconv.ParseUint(params["publicationId"], 10, 64)
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

	repository := repositories.PublicationRepository(db)

	haveSomeLike, erro := repository.PublicationHaveSomeLike(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	if !haveSomeLike {
		var msgError error = errors.New("não é possível remover like de uma publicação sem likes")
		responses.Error(w, http.StatusBadRequest, msgError)
		return
	}

	erro = repository.RemoveLikePublication(publicationId)
	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	msg := make(map[string]string)
	msg["msg"] = "Publicação descurtida com sucesso!"
	responses.JSON(w, http.StatusOK, msg)

}
