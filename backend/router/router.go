package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	storage "github.com/pastepi/url_shortener/backend/data"
	models "github.com/pastepi/url_shortener/backend/models"
	"github.com/pastepi/url_shortener/backend/shortener"
	"github.com/rs/cors"
)

func ServerInit() {
	r := mux.NewRouter()

	r.HandleFunc("/{ShortURL}", handleRedirect).Methods("GET")
	r.HandleFunc("/URL", handleURL).Methods("POST").
		Headers("Content-Type", "application/json")

	handler := cors.Default().Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handleURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var reqURL models.PostLink

	p, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err := json.Unmarshal(p, &reqURL); err != nil {
		panic(err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.CheckLink(&reqURL.Url) // Check for Protocol/Scheme
	storageURLs := storage.ReadURLs()

	var (
		newLink  = storage.FindByOrigURL(reqURL.Url, storageURLs)
		resp     []byte
		marshErr error
	)

	// if newLink is not empty, return the link from the storage
	// else, create a new entry and return it
	if newLink != (models.Link{}) {
		resp, marshErr = json.Marshal(newLink)

		if marshErr != nil {
			panic(err)
		}
	} else {
		newLink = models.Link{
			OriginURL: reqURL.Url,
			ShortURL:  shortener.ShortenURL(reqURL.Url),
		}
		resp, marshErr = json.Marshal(newLink)
		if marshErr != nil {
			panic(err)
		}
		storage.AppendLink(&storageURLs, &newLink)
		jsonURLs := storage.MarshalURLs(storageURLs)
		storage.SaveURLs(jsonURLs)
	}

	w.Write(resp)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	newURL := storage.FindByShortURL(vars["ShortURL"], storage.ReadURLs())
	if newURL == (models.Link{}) {
		newURL.OriginURL = "http://localhost:3000"
	}
	http.Redirect(w, r, newURL.OriginURL, http.StatusSeeOther)
}
