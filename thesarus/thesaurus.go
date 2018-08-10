package thesarus

type Thesarus interface {
	Synonyms(string) ([]string, error)
}
