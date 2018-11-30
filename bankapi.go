package main

import (
	"github.com/gorilla/mux"
	"net/http"
	. "bank-service/dao"
	. "bank-service/config"
	. "bank-service/models"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var userDao = UserDAO{}

func AllUser(w http.ResponseWriter, r *http.Request) {
	users, err := userDao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := userDao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func init() {
	config.Read()

	userDao.Server = config.Server
	userDao.Database = config.Database
	userDao.Connect()
}

func startServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	err := http.ListenAndServe(":3000", r)
	return err
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

