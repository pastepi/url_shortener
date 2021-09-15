package links

type Link struct {
	OriginURL string `json:"OriginURL"`
	ShortURL  string `json:"ShortURL"`
}

type PostLink struct {
	Url string
}
