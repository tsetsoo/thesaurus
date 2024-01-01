package thesaurus

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	synonymsStore SynonymsStore
}

func NewApi(synonymsStore SynonymsStore) Api {
	return Api{
		synonymsStore: synonymsStore,
	}
}

func (a Api) HandleAddSynonyms(w http.ResponseWriter, r *http.Request) {
	var payload createSynonymsRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	a.synonymsStore.AddSynonyms(payload.Word, payload.Synonym)
	w.WriteHeader(http.StatusCreated)
}

func (a Api) HandleSearchSynonyms(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	synonyms := a.synonymsStore.GetSynonyms(word)
	response := synonymsResponse{
		Synonyms: synonyms,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
