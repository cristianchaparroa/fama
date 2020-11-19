package handlers

const (
	StatusOK = "ok"
)

type NumberToWordsResponse struct {
	Status string `json:"status"`
	Words  string `json:"number_in_words"`
	Lang   string `json:"lang"`
}

func newNumberToWordsResponse(status, lang, words string) *NumberToWordsResponse {
	return &NumberToWordsResponse{
		Status: status,
		Lang:   lang,
		Words:  words,
	}
}
