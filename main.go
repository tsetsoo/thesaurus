package main

import (
	"spread.ai/thesaurus/thesaurus"
)

func main() {
	synonymsStore := thesaurus.NewInMemorySynonymsStore()
	api := thesaurus.NewApi(synonymsStore)
	server := thesaurus.NewServer(api)
	server.Run()
}
