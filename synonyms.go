package main

type synonyms struct {
	Noun words `json:"noun"`
	Verb words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}
