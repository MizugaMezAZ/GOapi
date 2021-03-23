package controller

import (
	"api/models"
	"encoding/json"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var ac models.Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ac.CreateAC(&ac); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusCreated, ac.Username+" Success create.")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var ac models.Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ac.DeleteAC(&ac); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusCreated, ac.Username+" Success delete.")
}

func ChangePwd(w http.ResponseWriter, r *http.Request) {
	var ac models.Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ac.UpdateAC(&ac); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusCreated, ac.Username+" Success change.")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var ac models.Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err, tf := ac.LoginAC(&ac)

	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, "not log")
		return
	}

	if tf {
		responseWithJson(w, http.StatusCreated, ac.Username+" Success login.")
	} else {
		responseWithJson(w, http.StatusCreated, ac.Username+" Fail login.")
	}

}
