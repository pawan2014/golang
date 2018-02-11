package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go get github.com/gorilla/mux

func main1() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ex1/{message}", index).Methods("GET")
	http.ListenAndServe(":8081", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling function")
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", vars)

	log.Println(vars)
}
