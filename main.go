package main

import (
	"GoAppTest/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/posts", controllers.GetPost)
	http.HandleFunc("/info", controllers.ReadInfo)
	http.HandleFunc("/qr", controllers.CreateQR)
	http.HandleFunc("/db", controllers.GetDBPosts)

	http.ListenAndServe(":8000", nil)
}
