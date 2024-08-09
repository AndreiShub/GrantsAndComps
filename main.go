package main

import (
	"log"
	"net/http"
	sw "github.com/gorilla/mux"
	"./go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	router.HandleFunc("/auth/login", AuthLoginPost)
	router.HandleFunc("/auth/check",AuthCheckPost)
	router.HandleFunc("/grants",GrantsGet)
	router.HandleFunc("/grants/{grant_id}",GrantsGrantIdGet)
	router.HandleFunc("/grants/{grant_id}/filters",GrantsGrantIdFiltersPut)

	log.Fatal(http.ListenAndServe(":8080", router))
}
