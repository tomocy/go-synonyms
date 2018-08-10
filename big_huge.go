package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type bigHuge struct {
	clientKey string
}

func (h bigHuge) fetchSynonyms(term string) ([]string, error) {
	syn := make([]string, 0)
	endpoint := "http://words.bighugelabs.com/api/2/" + h.clientKey + "/" + term + "/json"
	resp, err := http.Get(endpoint)
	if err != nil {
		return syn, fmt.Errorf("could not get the reponse from big huge: %s", err)
	}
	defer resp.Body.Close()

	var decodedResp synonyms
	if err := json.NewDecoder(resp.Body).Decode(&decodedResp); err != nil {
		if err != io.EOF {
			return syn, fmt.Errorf("could not decode response body: %s", err)
		}
	}

	syn = append(syn, decodedResp.Noun.Syn...)
	syn = append(syn, decodedResp.Verb.Syn...)

	return syn, nil
}
