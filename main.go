package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	check
	Status int `json:"status" db:"status"`
}

type check struct {
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}

type checks []check

func main() {
	seedFlag := flag.Bool("seed", false, "Run application seeds to Redis")
	dropFlag := flag.Bool("drop", false, "Destroy current Redis database")

	flag.Parse()

	initializeDB()

	if *seedFlag == true {
		seedDB()
	}

	if *dropFlag == true {
		dropDB()
	}

	run()
}

func run() {
	router := mux.NewRouter()
	router.HandleFunc("/", checkIndex).Methods("GET")
	router.HandleFunc("/checks", checkCreate).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func checkCreate(w http.ResponseWriter, r *http.Request) {

}

func checkIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	checks, err := redisClient.Keys("*").Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(checks)

	// json.NewEncoder(w).Encode("status: up")
}
