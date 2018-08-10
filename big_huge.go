package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type bigHuge struct {
	clientKey string
}

func (h bigHube) fetchSynonyms(term string) ([]string, error) {
	syn := make([]string, 0)
	resp, err := http.Get("http://word.bighugelabs.com/api/2/" + h.clientKey + "/" + term + "/json")
	if err != nil {
		return syn, fmt.Errorf("could not get the reponse from big huge: %s", err)
	}
	defer resp.Body.Close()

	var decodedResp synonyms
	if err := json.NewDecoder(resp.Body).Decode(&decodedResp); err != nil {
		return syn, fmt.Errorf("could not decode response body: %s", err)
	}

	syn = append(syn, decodedResp.Norn.Syn...)
	syn = append(syn, decodedResp.Verb.Syn...)

	return syn, nil
}
