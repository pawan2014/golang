package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("Listing at ")
	router := mux.NewRouter()
	router.HandleFunc("/emp", handleEmp).Methods("GET")
	router.HandleFunc("/emp/{ohr}", handleEmpFind).Methods("GET", "DELETE")
	router.HandleFunc("/emp", handleEmpPOST).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

/*
This is a Emp
*/
type Employe struct {
	EmpName string `json:"name"`
	Ohr     string `json:"ohr"`
}

var emp = map[string]*Employe{
	"703062333": &Employe{"Pawan", "703062333"},
	"803062333": &Employe{"Amit", "803062333"},
}

func handleEmp(res http.ResponseWriter, resp *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	outJSON, err := json.Marshal(emp)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outJSON))
	//fmt.Print(movies)
}

func handleEmpFind(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	reqVars := mux.Vars(req)
	myKey := reqVars["ohr"]

	log.Print("Trying with key ", myKey)

	if myemp, ok := emp[myKey]; ok {

		switch req.Method {
		// GET Method
		case "GET":
			log.Println("in Get...")
			outJSON, err := json.Marshal(myemp)
			if err != nil {
				http.Error(res, "Some Fatal Erroe", http.StatusInternalServerError)
				return
			}
			fmt.Fprint(res, string(outJSON))

		case "DELETE":
			// Delete me
			log.Println("in Delete...")
			delete(emp, myKey)
			res.WriteHeader(http.StatusNoContent)

		}

	} else {
		res.WriteHeader(http.StatusNotFound)
		mes, _ := json.Marshal("Key not found, try other key")
		fmt.Fprint(res, string(mes))
	}
}

func handleEmpPOST(res http.ResponseWriter, req *http.Request) {
	log.Println("in POST...")
	newemp := new(Employe)
	decorder := json.NewDecoder(req.Body)
	error := decorder.Decode(&newemp)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	emp[newemp.Ohr] = newemp
	newoutgoingJSON, err := json.Marshal(newemp)
	if err != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return

	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(newoutgoingJSON))
}
