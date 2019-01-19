package mock

type Retriever struct{
	Connects string
}

func (r Retriever) Get(url string) string{
	return r.Contents
}