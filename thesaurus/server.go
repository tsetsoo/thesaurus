package thesaurus

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	api Api
}

func NewServer(api Api) Server {
	return Server{api: api}
}

func (s Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/synonyms", s.api.HandleAddSynonyms).Methods("POST")
	r.HandleFunc("/synonyms/{word}", s.api.HandleSearchSynonyms).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
