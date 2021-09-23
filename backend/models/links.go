package links

type Link struct {
	ID        int64  `json:"ID"`
	OriginURL string `json:"OriginURL"`
	ShortURL  string `json:"ShortURL"`
}

type PostLink struct {
	Url string
}
