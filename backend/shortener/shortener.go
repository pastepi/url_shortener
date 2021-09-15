package shortener

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

func randomizeURL(u string) string {
	rand.Seed(time.Now().UnixNano())
	link := u + strconv.Itoa(rand.Intn(999))
	return link
}

func ShortenURL(u string) string {
	shortURL := base58.Encode([]byte(randomizeURL(u)))
	return shortURL[len(shortURL)-7:]
}
