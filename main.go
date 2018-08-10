package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tomocy/synonyms/thesarus"
)

func main() {
	clientKey := os.Getenv("BHT_CLIENT_KEY")
	bigHuge := thesarus.NewBigHuge(clientKey)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		syns, err := bigHuge.FetchSynonyms(scanner.Text())
		if err != nil {
			log.Fatalf("could not fetch synonyms: %s\n", err)
		}
		if len(syns) == 0 {
			log.Fatalln("there is no synonyms")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
