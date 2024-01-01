package thesaurus

import "sync"

type SynonymsStore interface {
	AddSynonyms(word, synonym string)
	GetSynonyms(word string) []string
}

type InMemorySynonymsStore struct {
	sync.RWMutex
	data map[string][]string
}

func NewInMemorySynonymsStore() *InMemorySynonymsStore {
	return &InMemorySynonymsStore{
		data: make(map[string][]string),
	}
}

// think about algo complexity
func (s *InMemorySynonymsStore) AddSynonyms(word, synonym string) {
	s.Lock()
	defer s.Unlock()
	allSynonyms := make(map[string]struct{}, 0)
	allSynonyms[word] = struct{}{}
	allSynonyms[synonym] = struct{}{}
	for _, syn := range s.data[word] {
		allSynonyms[syn] = struct{}{}
	}
	for _, syn := range s.data[synonym] {
		allSynonyms[syn] = struct{}{}
	}

	s.data[word] = append(s.data[word], synonym)
	s.data[synonym] = append(s.data[synonym], word)
	allSynonymsArr := make([]string, 0, len(allSynonyms))
	for syn := range allSynonyms {
		allSynonymsArr = append(allSynonymsArr, syn)
	}

	for _, syn := range allSynonymsArr {
		s.data[syn] = allSynonymsArr
	}
}

func (s *InMemorySynonymsStore) GetSynonyms(word string) []string {
	s.RLock()
	defer s.RUnlock()
	return s.data[word]
}
