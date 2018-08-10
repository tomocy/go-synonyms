package thesarus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type bigHuge struct {
	clientKey string
}

func newBigHuge(clientKey string) *bigHuge {
	return &bigHuge{
		clientKey: clientKey,
	}
}

type bigHugeSynonyms struct {
	Nouns bigHugeWords `json:"noun"`
	Verbs bigHugeWords `json:"verb"`
}

type bigHugeWords struct {
	Syns []string `json:"syn"`
}

func (h bigHuge) FetchSynonyms(term string) ([]string, error) {
	syns := make([]string, 0)
	endpoint := "http://words.bighugelabs.com/api/2/" + h.clientKey + "/" + term + "/json"
	resp, err := http.Get(endpoint)
	if err != nil {
		return syns, fmt.Errorf("could not get the reponse from big huge: %s", err)
	}
	defer resp.Body.Close()

	var decodedResp bigHugeSynonyms
	if err := json.NewDecoder(resp.Body).Decode(&decodedResp); err != nil {
		if err != io.EOF {
			return syns, fmt.Errorf("could not decode response body: %s", err)
		}
	}

	syns = append(syns, decodedResp.Nouns.Syns...)
	syns = append(syns, decodedResp.Verbs.Syns...)

	return syns, nil
}
