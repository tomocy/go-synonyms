package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	clientKey := os.Getenv("BHT_CLIENT_KEY")
	bigHuge := bigHuge{
		clientKey: clientKey,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		syns, err := bigHuge.fetchSynonyms(scanner.Text())
		if err != nil {
			log.Fatalf("could not fetch synonyms: %s\n", err)
		}
		if len(syns) == 0 {
			fmt.Println("there is no synonyms")
			continue
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
