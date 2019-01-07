package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/season/{season}/", getRandomQuote)
	headersOk := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func main() {
	handleRequests()
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	season := vars["season"]

	plan, _ := ioutil.ReadFile("../office-quotes/season" + season + ".json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err.Error())
	}

	writeToLog(r)
	json.NewEncoder(w).Encode(data)
}

func writeToLog(r *http.Request) {
	filename := "../logs/hits.log"
	text := time.Now().String() + "\n"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
