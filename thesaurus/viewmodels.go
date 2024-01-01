package thesaurus

type createSynonymsRequest struct {
	Word    string `json:"word"`
	Synonym string `json:"synonym"`
}

type synonymsResponse struct {
	Synonyms []string `json:"synonyms"`
}
