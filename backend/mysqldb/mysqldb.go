package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	links "github.com/pastepi/url_shortener/backend/models"
)

var db *sql.DB

func OpenConn() *sql.DB {
	var err error
	cfg := mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(10000)

	return db
}

func GetLinkByShortURL(shortLink string) (links.Link, error) {
	var url links.Link

	row := db.QueryRow("SELECT * FROM urls WHERE ShortURL = ?", shortLink)

	if err := row.Scan(&url.ID, &url.OriginURL, &url.ShortURL); err != nil {
		if err == sql.ErrNoRows {
			return url, fmt.Errorf("GetLinkByShortURL %s: no such shortURL", shortLink)
		}
		return url, fmt.Errorf("GetLinkByShortURL %s: %v", shortLink, err)
	}
	return url, nil
}

func GetLinkByOriginURL(originLink string) (links.Link, error) {
	var url links.Link

	row := db.QueryRow("SELECT * FROM urls WHERE OriginURL = ?", originLink)

	if err := row.Scan(&url.ID, &url.OriginURL, &url.ShortURL); err != nil {
		if err == sql.ErrNoRows {
			return url, fmt.Errorf("GetLinkByOriginURL %s: no such URL found", originLink)
		}
		return url, fmt.Errorf("GetLinkByOriginURL %s: %v", originLink, err)
	}
	return url, nil
}

func AddLink(link links.Link) (int64, error) {
	result, err := db.Exec("INSERT INTO urls (OriginURL, ShortURL) VALUES (?, ?)", link.OriginURL, link.ShortURL)
	if err != nil {
		return 0, fmt.Errorf("AddLink: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddLink: %v", err)
	}
	return id, nil
}