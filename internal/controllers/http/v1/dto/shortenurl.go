package dto

type ShortenURLRequest struct {
	URL string `json:"url"`
}

func (sh *ShortenURLRequest) Validate() bool {
	return len(sh.URL) != 0
}

type ShortenURLResponse struct {
	URL string `json:"result"`
}
