package search

type SearchEngine interface {
	Search(query string) ([]string, error)
	Index(data interface{}) error
}

type BleveSearchEngine struct{}

func (s BleveSearchEngine) Search(query string) ([]string, error) {
	return nil, nil
}

func (s BleveSearchEngine) Index(data interface{}) error {
	return nil
}
