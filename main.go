package main

import (
	"fmt"
	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", AllPeople).Methods("GET")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person", PostPerson).Methods("POST")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/person/{id}", UpdatePerson).Methods("PUT")
	http.Handle(" / ", router)
	fmt.Println(" Connected to port 8081 ")
	log.Fatal(http.ListenAndServe(":8081", router))

}
