package storage

import (
	"encoding/json"
	"net/url"
	"os"

	models "github.com/pastepi/url_shortener/backend/models"
)

// Save marshalled URL array to the file
func SaveURLs(u []byte) {
	file, err := os.OpenFile("./data/links.json", os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(u); err != nil {
		file.Close()
		panic(err)
	}
}

// Reading and returning unmarshalled URL array
func ReadURLs() []models.Link {

	// Create storage file with "[]" inside of it if it doesn't exist
	if _, err := os.Stat("./data/links.json"); os.IsNotExist(err) {
		os.WriteFile("./data/links.json", []byte("[]"), os.ModePerm)
	} else if err != nil {
		panic(err)
	}

	file, err := os.ReadFile("./data/links.json")
	if err != nil {
		panic(err)
	}
	// Unmarshal the array of URLs inside of links.json file
	// if the file is empty - write "[]" inside of it
	var URLs []models.Link
	if err := json.Unmarshal([]byte(file), &URLs); err != nil {
		os.WriteFile("./data/links.json", []byte("[]"), os.ModePerm)
		json.Unmarshal([]byte(file), &URLs)
	}

	return URLs
}

// Append models.Link struct to a slice
func AppendLink(u *[]models.Link, l *models.Link) {
	*u = append(*u, *l)
}

// Marshal URLs slice to JSON
func MarshalURLs(u []models.Link) []byte {
	jsonURLs, _ := json.MarshalIndent(u, "", "\t")
	return jsonURLs
}

func FindByOrigURL(l string, u []models.Link) models.Link {
	for _, v := range u {
		if v.OriginURL == l {
			return v
		}
	}
	return models.Link{}
}

func FindByShortURL(l string, u []models.Link) models.Link {
	for _, v := range u {
		if v.ShortURL == l {
			return v
		}
	}
	return models.Link{}
}

// Checks for the scheme/protocol of the link - if none exists, adds "HTTPS"
func CheckLink(l *string) {
	p, err := url.Parse(*l)
	if err != nil {
		panic(err)
	}

	if p.Scheme == "" {
		*l = "https://" + *l
	}
}
