package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/skip2/go-qrcode"
)

// Random Controllers
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the index v2")
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This could be a post")
}

func ReadInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Another bit of info")
}

func CreateQR(w http.ResponseWriter, r *http.Request) {
	qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	fmt.Fprintf(w, "Create a QR Image")
}

func GetDBPosts(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	results := []string{}

	for rows.Next() {
		var id string
		var title string
		var body string
		err := rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err)
		}
		results = append(results, fmt.Sprintf("%s, %s, %s", id, title, body))
		//fmt.Printf("ID: %s, Title: %s, Body: %s \n", id, title, body)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	fmt.Fprintf(w, results[0])
}
