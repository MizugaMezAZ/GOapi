package controller

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

var (
	dao = models.Account{}
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

	if err := dao.CreateAC(ac); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusCreated, ac)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func ChangePwd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}
