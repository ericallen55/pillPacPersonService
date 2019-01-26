package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Person struct {
	Id        int
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Age       string `json:"age"`
}

type Persons []Person

func AllPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getAllPeopleDb())
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	person := getPersonDb(stringToInt(key, w))
	json.NewEncoder(w).Encode(person)
}

func PostPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	json.NewEncoder(w).Encode(addPersonDb(p))
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	deletePersonDb(stringToInt(key, w))
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var p Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	json.NewEncoder(w).Encode(updatePersonDb(stringToInt(key, w), p))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func stringToInt(s string, w http.ResponseWriter) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Id must be int")
		return 0
	}
	return id
}
