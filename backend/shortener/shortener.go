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

// randomizeURL takes in a link, concatenates it with a pseudo-random
// number, so that in the event of the same link being passed in, the ShortenURL function
// is less likely to return the same result.

// Using base58 - primary reason is it excludes easily mistakable characters such as "l" and "I"
// as well as "O" and "0"

// Shorten URL returns last 7 characters of the encoded string - with no calculations included,
// 7 characters is enough for there to be billions of combinations of possible outcomes.
// The collision chance is low enough for this kind of app (MVP).
// In the case that this app is expanded and worked on as a real product - consider including exact calculations.
