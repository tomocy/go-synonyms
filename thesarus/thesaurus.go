package thesarus

type Thesarus interface {
	FetchSynonyms(string) ([]string, error)
}

func NewBigHuge(clientKey string) Thesarus {
	return newBigHuge(clientKey)
}
