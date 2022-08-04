package main

import (
	"fmt"
	"github.com/go-blog-app/database"
	handler "github.com/go-blog-app/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	initRouting()
}

func initRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler).Methods("GET")

	fmt.Println("http://127.0.0.1:8086")
	http.ListenAndServe("127.0.0.1:8086", r)
}
